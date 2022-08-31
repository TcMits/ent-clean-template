package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
)

func Test_createModelUseCase_Create(t *testing.T) {
	type fields struct {
		repository          repository.CreateModelRepository[*model.User, *model.UserCreateInput]
		validateFunc        CreateValidateFunc[*model.UserCreateInput, *model.UserCreateInput]
		wrapCreateErrorFunc func(error) error
	}
	type args struct {
		ctx   context.Context
		input *model.UserCreateInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()

	repository := repository.NewUserRepository(client)
	createInput := factory.UserFactory.Build(ctx, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserCreateInput
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository: repository,
				validateFunc: func(uci *model.UserCreateInput) (*model.UserCreateInput, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: createInput,
			},
			want: createInput,
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository: repository,
				validateFunc: func(uci *model.UserCreateInput) (*model.UserCreateInput, error) {
					return nil, errors.New("test")
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: createInput,
			},
			wantErr: true,
		},
		{
			name: "CreateError",
			fields: fields{
				repository: repository,
				validateFunc: func(uci *model.UserCreateInput) (*model.UserCreateInput, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: createInput,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &createModelUseCase[*model.User, *model.UserCreateInput, *model.UserCreateInput]{
				repository:          tt.fields.repository,
				validateFunc:        tt.fields.validateFunc,
				wrapCreateErrorFunc: tt.fields.wrapCreateErrorFunc,
			}
			got, err := l.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createModelUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				if got.Password != *tt.want.Password {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.Password, *tt.want.Password)
				}
				if got.Username != tt.want.Username {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.Username, tt.want.Username)
				}
				if got.FirstName != *tt.want.FirstName {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.FirstName, *tt.want.FirstName)
				}
				if got.LastName != *tt.want.LastName {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.FirstName, *tt.want.LastName)
				}
				if got.Email != tt.want.Email {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.Email, tt.want.Email)
				}
				if got.IsStaff != *tt.want.IsStaff {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.IsStaff, *tt.want.IsStaff)
				}
				if got.IsActive != *tt.want.IsActive {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.IsActive, *tt.want.IsActive)
				}
			}
		})
	}
}

func Test_createModelInTransactionUseCase_Create(t *testing.T) {
	type fields struct {
		repository            repository.CreateWithClientModelRepository[*model.User, *model.UserCreateInput]
		transactionRepository repository.TransactionRepository
		validateFunc          CreateInTransactionValidateFunc[*model.UserCreateInput, *model.UserCreateInput]
		wrapCreateErrorFunc   func(error) error
	}
	type args struct {
		ctx   context.Context
		input *model.UserCreateInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()

	repo := repository.NewUserRepository(client)
	transactionRepository := repository.NewTransactionRepository(client)
	createInput := factory.UserFactory.Build(ctx, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserCreateInput
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository:            repo,
				transactionRepository: transactionRepository,
				validateFunc: func(uci *model.UserCreateInput, client *ent.Client) (*model.UserCreateInput, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: createInput,
			},
			want: createInput,
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository:            repo,
				transactionRepository: transactionRepository,
				validateFunc: func(uci *model.UserCreateInput, client *ent.Client) (*model.UserCreateInput, error) {
					return nil, errors.New("test")
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: createInput,
			},
			wantErr: true,
		},
		{
			name: "CreateError",
			fields: fields{
				repository:            repo,
				transactionRepository: transactionRepository,
				validateFunc: func(uci *model.UserCreateInput, client *ent.Client) (*model.UserCreateInput, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: createInput,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &createModelInTransactionUseCase[*model.User, *model.UserCreateInput, *model.UserCreateInput]{
				repository:            tt.fields.repository,
				transactionRepository: tt.fields.transactionRepository,
				validateFunc:          tt.fields.validateFunc,
				wrapCreateErrorFunc:   tt.fields.wrapCreateErrorFunc,
			}
			got, err := l.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createModelInTransactionUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				if got.Password != *tt.want.Password {
					t.Errorf("createModelInTransactionUseCase.Create() = %v, want %v", got.Password, *tt.want.Password)
				}
				if got.Username != tt.want.Username {
					t.Errorf("createModelInTransactionUseCase.Create() = %v, want %v", got.Username, tt.want.Username)
				}
				if got.FirstName != *tt.want.FirstName {
					t.Errorf("createModelInTransactionUseCase.Create() = %v, want %v", got.FirstName, *tt.want.FirstName)
				}
				if got.LastName != *tt.want.LastName {
					t.Errorf("createModelInTransactionUseCase.Create() = %v, want %v", got.FirstName, *tt.want.LastName)
				}
				if got.Email != tt.want.Email {
					t.Errorf("createModelInTransactionUseCase.Create() = %v, want %v", got.Email, tt.want.Email)
				}
				if got.IsStaff != *tt.want.IsStaff {
					t.Errorf("createModelInTransactionUseCase.Create() = %v, want %v", got.IsStaff, *tt.want.IsStaff)
				}
				if got.IsActive != *tt.want.IsActive {
					t.Errorf("createModelInTransactionUseCase.Create() = %v, want %v", got.IsActive, *tt.want.IsActive)
				}
			}
		})
	}
}
