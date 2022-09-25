package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

func TestNewLoginRepository(t *testing.T) {
	type args struct {
		client *ent.Client
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()

	tests := []struct {
		name string
		args args
		want LoginRepository[*model.User, *model.UserWhereInput, *useCaseModel.LoginInput]
	}{
		{
			name: "Success",
			args: args{client: client},
			want: &loginRepository{client: client},
		},
	}
	for _, tt := range tests {
		if got := NewLoginRepository(tt.args.client); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewLoginRepository() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginRepository_Get(t *testing.T) {
	type fields struct {
		client *ent.Client
	}

	type args struct {
		ctx            context.Context
		userWhereInput *model.UserWhereInput
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	userA, err := factory.GetUserFactory().Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	_, err = factory.GetUserFactory().Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	testFirstName := "test"
	testIsActive := true

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name:   "Success",
			fields: fields{client: client},
			args: args{ctx: ctx, userWhereInput: &model.UserWhereInput{
				ID: &userA.ID,
			}},
			want: userA,
		},
		{
			name:   "NotFound",
			fields: fields{client: client},
			args: args{ctx: ctx, userWhereInput: &model.UserWhereInput{
				FirstName: &testFirstName,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "ManyRecords",
			fields: fields{client: client},
			args: args{ctx: ctx, userWhereInput: &model.UserWhereInput{
				IsActive: &testIsActive,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		repo := &loginRepository{
			client: tt.fields.client,
		}
		got, err := repo.Get(tt.args.ctx, tt.args.userWhereInput)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginRepository.Get() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != nil && tt.want != nil && !reflect.DeepEqual(got.ID, tt.want.ID) {
			t.Errorf("%q. loginRepository.Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loginRepository_Login(t *testing.T) {
	type fields struct {
		client *ent.Client
	}
	type args struct {
		ctx        context.Context
		loginInput *useCaseModel.LoginInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx))
	userA, err := factory.GetUserFactory().Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name:   "Success",
			fields: fields{client: client},
			args: args{ctx: ctx, loginInput: &useCaseModel.LoginInput{
				Username: userA.Username,
				Password: "12345678",
			}},
			want: userA,
		},
		{
			name:   "WrongUsername",
			fields: fields{client: client},
			args: args{ctx: ctx, loginInput: &useCaseModel.LoginInput{
				Username: userA.Username + "wrong",
				Password: "12345678",
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "WrongPassword",
			fields: fields{client: client},
			args: args{ctx: ctx, loginInput: &useCaseModel.LoginInput{
				Username: userA.Username,
				Password: "123456789",
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		repo := &loginRepository{
			client: tt.fields.client,
		}
		got, err := repo.Login(tt.args.ctx, tt.args.loginInput)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginRepository.Login() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != nil && tt.want != nil && !reflect.DeepEqual(got.ID, tt.want.ID) {
			t.Errorf("%q. loginRepository.Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
