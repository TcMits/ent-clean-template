package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/stretchr/testify/require"
)

func Test_deleteModelUseCase_Delete(t *testing.T) {
	type fields struct {
		repository           repository.DeleteModelRepository[*model.User]
		getRepository        repository.GetModelRepository[*model.User, *model.UserWhereInput]
		toRepoWhereInputFunc ConverFunc[*model.UserWhereInput, *model.UserWhereInput]
		wrapGetErrorFunc     func(error) error
		wrapDeleteErrorFunc  func(error) error
	}
	type args struct {
		ctx   context.Context
		input *model.UserWhereInput
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	repository := repository.NewUserRepository(client)

	tests := []struct {
		name    string
		fields  fields
		args    args
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
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: &model.UserWhereInput{ID: &u.ID},
			},
		},
		{
			name: "toRepoWhereInputFuncuccessError",
			fields: fields{
				repository:    repository,
				getRepository: repository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: &model.UserWhereInput{ID: &u.ID},
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
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: &model.UserWhereInput{ID: &u.ID},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &deleteModelUseCase[*model.User, *model.UserWhereInput, *model.UserWhereInput]{
				repository:           tt.fields.repository,
				getRepository:        tt.fields.getRepository,
				toRepoWhereInputFunc: tt.fields.toRepoWhereInputFunc,
				wrapGetErrorFunc:     tt.fields.wrapGetErrorFunc,
				wrapDeleteErrorFunc:  tt.fields.wrapDeleteErrorFunc,
			}
			err := l.GetAndDelete(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createModelUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
func Test_deleteModelInTransactionUseCase_Delete(t *testing.T) {
	type fields struct {
		repository            repository.DeleteWithClientModelRepository[*model.User]
		getRepository         repository.GetWithClientModelRepository[*model.User, *model.UserWhereInput]
		transactionRepository repository.TransactionRepository
		toRepoWhereInputFunc  ConverFunc[*model.UserWhereInput, *model.UserWhereInput]
		selectForUpdate       bool
		wrapGetErrorFunc      func(error) error
		wrapDeleteErrorFunc   func(error) error
	}
	type args struct {
		ctx   context.Context
		input *model.UserWhereInput
	}

	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	repo := repository.NewUserRepository(client)
	transactionRepository := repository.NewTransactionRepository(client)

	tests := []struct {
		name    string
		fields  fields
		args    args
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
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: &model.UserWhereInput{ID: &u.ID},
			},
		},
		{
			name: "toRepoWhereInputFuncuccessError",
			fields: fields{
				repository:            repo,
				getRepository:         repo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *model.UserWhereInput) (*model.UserWhereInput, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: &model.UserWhereInput{ID: &u.ID},
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
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: &model.UserWhereInput{ID: &u.ID},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &deleteModelInTransactionUseCase[*model.User, *model.UserWhereInput, *model.UserWhereInput]{
				repository:            tt.fields.repository,
				getRepository:         tt.fields.getRepository,
				transactionRepository: tt.fields.transactionRepository,
				toRepoWhereInputFunc:  tt.fields.toRepoWhereInputFunc,
				selectForUpdate:       tt.fields.selectForUpdate,
				wrapGetErrorFunc:      tt.fields.wrapGetErrorFunc,
				wrapDeleteErrorFunc:   tt.fields.wrapDeleteErrorFunc,
			}
			err := l.GetAndDelete(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createModelUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
