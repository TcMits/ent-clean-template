package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	jwtKit "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewLoginUseCase(t *testing.T) {
	type args struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

	want := loginUseCase{
		repository:    loginRepository,
		getRepository: getRepository,
		secret:        "secret",
	}

	tests := []struct {
		name string
		args args
		want LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User]
	}{
		{
			name: "Success",
			args: args{
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "secret",
			},
			want: &want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginUseCase(tt.args.repository, tt.args.getRepository, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_getUserMapClaims(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		user *model.User
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   jwtKit.MapClaims
	}{
		{
			name: "Success",
			fields: fields{
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				user: u,
			},
			want: jwtKit.MapClaims{
				idFieldName:  u.ID.String(),
				"email":      u.Email,
				keyFieldName: u.JwtTokenKey,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			if got := l.getUserMapClaims(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loginUseCase.getUserMapClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_getUserFromMapClaims(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx          context.Context
		jwtMapClaims jwtKit.MapClaims
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					idFieldName:  u.ID.String(),
					"email":      u.Email,
					keyFieldName: u.JwtTokenKey,
				},
			},
			want: u,
		},
		{
			name: "MissingID",
			fields: fields{
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					"email":      u.Email,
					keyFieldName: u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongUUID",
			fields: fields{
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					idFieldName:  "hello",
					"email":      u.Email,
					keyFieldName: u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongUUID",
			fields: fields{
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					idFieldName:  uuid.NewString(),
					"email":      u.Email,
					keyFieldName: u.JwtTokenKey,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "WrongKey",
			fields: fields{
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "secret",
			},
			args: args{
				ctx: ctx,
				jwtMapClaims: jwtKit.MapClaims{
					idFieldName:  u.ID,
					"email":      u.Email,
					keyFieldName: "hello",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
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

func Test_loginUseCase_createAccessToken(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		user *model.User
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				u,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.createAccessToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.createAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if newU, err := l.parseAccessToken(ctx, got); err != nil || !reflect.DeepEqual(newU.ID, tt.want.ID) {
				t.Errorf("loginUseCase.createAccessToken() = %v, want %v", newU, tt.want)
			}
		})
	}
}

func Test_loginUseCase_createRefreshToken(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		user *model.User
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{user: u},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.createRefreshToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.createRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if newU, err := l.parseRefreshToken(ctx, got); err != nil || !reflect.DeepEqual(newU.ID, tt.want.ID) {
				t.Errorf("loginUseCase.createRefreshToken() = %v, want %v", newU, tt.want)
			}
		})
	}
}

func Test_loginUseCase_parseAccessToken(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx   context.Context
		token string
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

	l := &loginUseCase{
		repository:    loginRepository,
		getRepository: getRepository,
		secret:        "Dummy",
	}
	token, err := l.createAccessToken(u)
	require.NoError(t, err)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:   ctx,
				token: token,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.parseAccessToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.parseAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.parseAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_parseRefreshToken(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx               context.Context
		refreshTokenInput *useCaseModel.RefreshTokenInput
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

	l := &loginUseCase{
		repository:    loginRepository,
		getRepository: getRepository,
		secret:        "Dummy",
	}
	refreshTokenInput, err := l.createRefreshToken(u)
	require.NoError(t, err)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:               ctx,
				refreshTokenInput: refreshTokenInput,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.parseRefreshToken(tt.args.ctx, tt.args.refreshTokenInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.parseRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.parseRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loginUseCase_Login(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx        context.Context
		loginInput *useCaseModel.LoginInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx: ctx,
				loginInput: &useCaseModel.LoginInput{
					Username: u.Username,
					Password: "12345678",
				},
			},
			want: u,
		},
		{
			name: "WrongPassword",
			fields: fields{
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx: ctx,
				loginInput: &useCaseModel.LoginInput{
					Username: u.Username,
					Password: "1234567",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.Login(tt.args.ctx, tt.args.loginInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if newU, err := l.parseAccessToken(ctx, got.AccessToken); err != nil || !reflect.DeepEqual(newU.ID, tt.want.ID) {
					t.Errorf("loginUseCase.Login() = %v, want %v", newU, tt.want)
				}
				if newU, err := l.parseRefreshToken(ctx, &useCaseModel.RefreshTokenInput{
					RefreshToken: got.RefreshToken,
					RefreshKey:   got.RefreshKey,
				}); err != nil || !reflect.DeepEqual(newU.ID, tt.want.ID) {
					t.Errorf("loginUseCase.Login() = %v, want %v", newU, tt.want)
				}
			}
		})
	}
}

func Test_loginUseCase_RefreshToken(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx               context.Context
		refreshTokenInput *useCaseModel.RefreshTokenInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

	l := &loginUseCase{
		repository:    loginRepository,
		getRepository: getRepository,
		secret:        "Dummy",
	}
	refreshTokenInput, err := l.createRefreshToken(u)
	require.NoError(t, err)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:               ctx,
				refreshTokenInput: refreshTokenInput,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.RefreshToken(tt.args.ctx, tt.args.refreshTokenInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if newU, err := l.parseAccessToken(ctx, got); err != nil || !reflect.DeepEqual(newU.ID, tt.want.ID) {
				t.Errorf("loginUseCase.RefreshToken() = %v, want %v", newU, tt.want)
			}
		})
	}
}

func Test_loginUseCase_VerifyToken(t *testing.T) {
	type fields struct {
		repository    repository.LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
		getRepository repository.GetModelRepository[*model.User, *model.UserWhereInput]
		secret        string
	}
	type args struct {
		ctx   context.Context
		token string
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	loginRepository := repository.NewLoginRepository(client)
	getRepository := repository.NewUserRepository(client)

	l := &loginUseCase{
		repository:    loginRepository,
		getRepository: getRepository,
		secret:        "Dummy",
	}
	token, err := l.createAccessToken(u)
	require.NoError(t, err)

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
				repository:    loginRepository,
				getRepository: getRepository,
				secret:        "Dummy",
			},
			args: args{
				ctx:   ctx,
				token: token,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &loginUseCase{
				repository:    tt.fields.repository,
				getRepository: tt.fields.getRepository,
				secret:        tt.fields.secret,
			}
			got, err := l.VerifyToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUseCase.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ID, tt.want.ID) {
				t.Errorf("loginUseCase.VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
