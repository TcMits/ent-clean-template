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
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_updateModelUseCase_Update(t *testing.T) {
	type fields struct {
		repository           repository.UpdateModelRepository[*model.User, *model.UserUpdateInput]
		getRepository        repository.GetModelRepository[*model.User, *model.UserWhereInput]
		toRepoWhereInputFunc ConverFunc[*model.UserWhereInput, *model.UserWhereInput]
		validateFunc         UpdateValidateFunc[*model.User, *model.UserUpdateInput, *model.UserUpdateInput]
		wrapGetErrorFunc     func(error) error
		wrapUpdateErrorFunc  func(error) error
	}
	type args struct {
		ctx         context.Context
		whereInput  *model.UserWhereInput
		updateInput *model.UserUpdateInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u1, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	u2, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	newUUID := uuid.New()

	testEmail := "test@gmail.com"
	repository := repository.NewUserRepository(client)
	updateInput := &model.UserUpdateInput{Email: &testEmail}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserUpdateInput
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository:    repository,
				getRepository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u1.ID},
				updateInput: updateInput,
			},
			want: updateInput,
		},
		{
			name: "toRepoWhereInputFuncError",
			fields: fields{
				repository:    repository,
				getRepository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return nil, errors.New("test")
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u2.ID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository:    repository,
				getRepository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &newUUID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository:    repository,
				getRepository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput) (*model.UserUpdateInput, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u2.ID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
		{
			name: "UpdateError",
			fields: fields{
				repository:    repository,
				getRepository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u2.ID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &updateModelUseCase[*model.User, *model.UserWhereInput, *model.UserUpdateInput, *model.UserWhereInput, *model.UserUpdateInput]{
				repository:           tt.fields.repository,
				getRepository:        tt.fields.getRepository,
				toRepoWhereInputFunc: tt.fields.toRepoWhereInputFunc,
				validateFunc:         tt.fields.validateFunc,
				wrapGetErrorFunc:     tt.fields.wrapGetErrorFunc,
				wrapUpdateErrorFunc:  tt.fields.wrapUpdateErrorFunc,
			}
			got, err := l.GetAndUpdate(tt.args.ctx, tt.args.whereInput, tt.args.updateInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("createModelUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				if got.Email != *tt.want.Email {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.Email, *tt.want.Email)
				}
			}
		})
	}
}

func Test_updateModelInTransactionUseCase_Update(t *testing.T) {
	type fields struct {
		repository            repository.UpdateWithClientModelRepository[*model.User, *model.UserUpdateInput]
		getRepository         repository.GetWithClientModelRepository[*model.User, *model.UserWhereInput]
		transactionRepository repository.TransactionRepository
		toRepoWhereInputFunc  ConverFunc[*model.UserWhereInput, *model.UserWhereInput]
		validateFunc          UpdateInTransactionValidateFunc[*model.User, *model.UserUpdateInput, *model.UserUpdateInput]
		selectForUpdate       bool
		wrapGetErrorFunc      func(error) error
		wrapUpdateErrorFunc   func(error) error
	}
	type args struct {
		ctx         context.Context
		whereInput  *model.UserWhereInput
		updateInput *model.UserUpdateInput
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u1, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	u2, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	newUUID := uuid.New()

	testEmail := "test@gmail.com"
	repo := repository.NewUserRepository(client)
	transactionRepository := repository.NewTransactionRepository(client)
	updateInput := &model.UserUpdateInput{Email: &testEmail}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserUpdateInput
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository:            repo,
				getRepository:         repo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput, client *ent.Client) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u1.ID},
				updateInput: updateInput,
			},
			want: updateInput,
		},
		{
			name: "toRepoWhereInputFuncError",
			fields: fields{
				repository:            repo,
				getRepository:         repo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return nil, errors.New("test")
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput, client *ent.Client) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u2.ID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository:            repo,
				getRepository:         repo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput, client *ent.Client) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &newUUID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository:            repo,
				getRepository:         repo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput, client *ent.Client) (*model.UserUpdateInput, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u2.ID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
		{
			name: "UpdateError",
			fields: fields{
				repository:            repo,
				getRepository:         repo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return uwi, nil
				},
				validateFunc: func(ins *model.User, uui *model.UserUpdateInput, client *ent.Client) (*model.UserUpdateInput, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  &model.UserWhereInput{ID: &u2.ID},
				updateInput: updateInput,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &updateModelInTransactionUseCase[*model.User, *model.UserWhereInput, *model.UserUpdateInput, *model.UserWhereInput, *model.UserUpdateInput]{
				repository:            tt.fields.repository,
				getRepository:         tt.fields.getRepository,
				transactionRepository: tt.fields.transactionRepository,
				toRepoWhereInputFunc:  tt.fields.toRepoWhereInputFunc,
				validateFunc:          tt.fields.validateFunc,
				selectForUpdate:       tt.fields.selectForUpdate,
				wrapGetErrorFunc:      tt.fields.wrapGetErrorFunc,
				wrapUpdateErrorFunc:   tt.fields.wrapUpdateErrorFunc,
			}
			got, err := l.GetAndUpdate(tt.args.ctx, tt.args.whereInput, tt.args.updateInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("createModelUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				if got.Email != *tt.want.Email {
					t.Errorf("createModelUseCase.Create() = %v, want %v", got.Email, *tt.want.Email)
				}
			}
		})
	}
}
