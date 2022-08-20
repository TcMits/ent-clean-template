package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/ent/user"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/stretchr/testify/require"
)

func TestNewLoginRepository(t *testing.T) {
	type args struct {
		client *ent.Client
	}
	// Create an SQLite memory database and generate the schema.
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(context.Background()))

	tests := []struct {
		name string
		args args
		want LoginRepository
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
		predicateUsers []model.PredicateUser
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	defer client.Close()
	require.NoError(t, client.Schema.Create(ctx))

	userA, err := client.
		User.
		Create().
		SetUsername("userA").
		SetFirstName("user").
		SetLastName("A").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.
		User.
		Create().
		SetUsername("userB").
		SetFirstName("user").
		SetLastName("B").
		Save(ctx)
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
			args: args{ctx: ctx, predicateUsers: []model.PredicateUser{
				user.IDEQ(userA.ID),
			}},
			want: userA,
		},
		{
			name:   "NotFound",
			fields: fields{client: client},
			args: args{ctx: ctx, predicateUsers: []model.PredicateUser{
				user.FirstNameEQ("test"),
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "ManyRecords",
			fields: fields{client: client},
			args: args{ctx: ctx, predicateUsers: []model.PredicateUser{
				user.FirstNameEQ("user"),
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		repo := &loginRepository{
			client: tt.fields.client,
		}
		got, err := repo.Get(tt.args.ctx, tt.args.predicateUsers...)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. loginRepository.Get() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
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

	userA, err := client.
		User.
		Create().
		SetUsername("userA").
		SetFirstName("user").
		SetLastName("A").
		Save(ctx)
	require.NoError(t, err)

	_, err = client.
		User.
		Create().
		SetUsername("userB").
		SetFirstName("user").
		SetLastName("B").
		Save(ctx)
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
				Username: "",
				Password: "",
			}},
			want: userA,
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
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loginRepository.Login() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
