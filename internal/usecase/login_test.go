package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	jwtKit "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

func TestNewLoginUseCase(t *testing.T) {
	type args struct {
		repository repository.LoginRepository[model.User, model.PredicateUser, useCaseModel.LoginInput]
		secret     string
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx))

	loginRepository := repository.NewLoginRepository(client)

	want := loginUseCase{
		repository: loginRepository,
		secret:     "secret",
	}

	tests := []struct {
		name string
		args args
		want LoginUseCase[useCaseModel.LoginInput, useCaseModel.JWTAuthenticatedPayload, useCaseModel.RefreshTokenInput, model.User]
	}{
		{
			name: "Success",
			args: args{
				repository: loginRepository,
				secret:     "secret",
			},
			want: &want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginUseCase(tt.args.repository, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_getUserMapClaims(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository[model.User, model.PredicateUser, useCaseModel.LoginInput]
		secret     string
	}
	type args struct {
		user *model.User
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx))
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   jwtKit.MapClaims
	}{
		{
			name: "Success",
			fields: fields{
				repository: loginRepository,
				secret:     "secret",
			},
			args: args{
				user: u,
			},
			want: jwtKit.MapClaims{
				"id":    u.ID.String(),
				"email": u.Email,
				"key":   u.JwtTokenKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository: tt.fields.repository,
				secret:     tt.fields.secret,
			}
			if got := l.getUserMapClaims(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginUseCase.getUserMapClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_getUserFromMapClaims(t *testing.T) {
	type fields struct {
		repository repository.LoginRepository[model.User, model.PredicateUser, useCaseModel.LoginInput]
		secret     string
	}
	type args struct {
		ctx          context.Context
		jwtMapClaims jwtKit.MapClaims
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx))
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository: loginRepository,
				secret:     "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					"id":    u.ID.String(),
					"email": u.Email,
					"key":   u.JwtTokenKey,
				},
			},
			want: u,
		},
		{
			name: "MissingID",
			fields: fields{
				repository: loginRepository,
				secret:     "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					"email": u.Email,
					"key":   u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongUUID",
			fields: fields{
				repository: loginRepository,
				secret:     "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					"id":    "hello",
					"email": u.Email,
					"key":   u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongUUID",
			fields: fields{
				repository: loginRepository,
				secret:     "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					"id":    uuid.NewString(),
					"email": u.Email,
					"key":   u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongKey",
			fields: fields{
				repository: loginRepository,
				secret:     "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					"id":    u.ID,
					"email": u.Email,
					"key":   "hello",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository: tt.fields.repository,
				secret:     tt.fields.secret,
			}
			got, err := l.getUserFromMapClaims(tt.args.ctx, tt.args.jwtMapClaims)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.getUserFromMapClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.getUserFromMapClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}
