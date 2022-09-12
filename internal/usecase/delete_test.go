package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	gomock "github.com/golang/mock/gomock"
)

func Test_deleteModelUseCase_Delete(t *testing.T) {
	type fields struct {
		repository           repository.DeleteModelRepository[*struct{}]
		getRepository        repository.GetModelRepository[*struct{}, *struct{}]
		toRepoWhereInputFunc ConverFunc[*struct{}, *struct{}]
		wrapGetErrorFunc     func(error) error
		wrapDeleteErrorFunc  func(error) error
	}
	type args struct {
		ctx   context.Context
		input *struct{}
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetModelRepository[*struct{}, *struct{}](ctrl)
	deleteRepo := repository.NewMockDeleteModelRepository[*struct{}](ctrl)

	getRepo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	getRepo.EXPECT().Get(
		gomock.Eq(ctx), gomock.Nil(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	deleteRepo.EXPECT().Delete(
		gomock.Eq(ctx), gomock.Eq(new(struct{})),
	).Return(
		nil,
	).AnyTimes()

	deleteRepo.EXPECT().Delete(
		gomock.Eq(ctx), gomock.Nil(),
	).Return(
		errors.New(""),
	).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository:    deleteRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
		},
		{
			name: "toRepoWhereInputFuncuccessError",
			fields: fields{
				repository:    deleteRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository:    deleteRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &deleteModelUseCase[*struct{}, *struct{}, *struct{}]{
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
		repository            repository.DeleteWithClientModelRepository[*struct{}]
		getRepository         repository.GetWithClientModelRepository[*struct{}, *struct{}]
		transactionRepository repository.TransactionRepository
		toRepoWhereInputFunc  ConverFunc[*struct{}, *struct{}]
		selectForUpdate       bool
		wrapGetErrorFunc      func(error) error
		wrapDeleteErrorFunc   func(error) error
	}
	type args struct {
		ctx   context.Context
		input *struct{}
	}

	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetWithClientModelRepository[*struct{}, *struct{}](ctrl)
	deleteRepo := repository.NewMockDeleteWithClientModelRepository[*struct{}](ctrl)

	getRepo.EXPECT().GetWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Eq(new(struct{})), false,
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	getRepo.EXPECT().GetWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Nil(), false,
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	deleteRepo.EXPECT().DeleteWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Eq(new(struct{})),
	).Return(
		nil,
	).AnyTimes()

	deleteRepo.EXPECT().DeleteWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Nil(),
	).Return(
		errors.New(""),
	).AnyTimes()

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
				repository:            deleteRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
		},
		{
			name: "toRepoWhereInputFuncuccessError",
			fields: fields{
				repository:            deleteRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository:            deleteRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapDeleteErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &deleteModelInTransactionUseCase[*struct{}, *struct{}, *struct{}]{
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