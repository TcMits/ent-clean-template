package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	gomock "github.com/golang/mock/gomock"
)

func Test_updateModelUseCase_Update(t *testing.T) {
	type fields struct {
		repository           repository.UpdateModelRepository[*struct{}, *struct{}]
		getRepository        repository.GetModelRepository[*struct{}, *struct{}]
		toRepoWhereInputFunc ConverFunc[*struct{}, *struct{}]
		validateFunc         UpdateValidateFunc[*struct{}, *struct{}, *struct{}]
		wrapGetErrorFunc     func(error) error
		wrapUpdateErrorFunc  func(error) error
	}
	type args struct {
		ctx         context.Context
		whereInput  *struct{}
		updateInput *struct{}
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetModelRepository[*struct{}, *struct{}](ctrl)
	updateRepo := repository.NewMockUpdateModelRepository[*struct{}, *struct{}](ctrl)

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

	updateRepo.EXPECT().Update(
		gomock.Eq(ctx), gomock.Eq(new(struct{})), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	updateRepo.EXPECT().Update(
		gomock.Eq(ctx), gomock.Eq(new(struct{})), gomock.Nil(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *struct{}
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository:    updateRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: new(struct{}),
			},
			want: new(struct{}),
		},
		{
			name: "toRepoWhereInputFuncError",
			fields: fields{
				repository:    updateRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				validateFunc: func(ins *struct{}, uui *struct{}) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository:    updateRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  nil,
				updateInput: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository:    updateRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "UpdateError",
			fields: fields{
				repository:    updateRepo,
				getRepository: getRepo,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &updateModelUseCase[*struct{}, *struct{}, *struct{}, *struct{}, *struct{}]{
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

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createModelUseCase.Create() = %v, want %v", got, *tt.want)
			}
		})
	}
}

func Test_updateModelInTransactionUseCase_Update(t *testing.T) {
	type fields struct {
		repository            repository.UpdateWithClientModelRepository[*struct{}, *struct{}]
		getRepository         repository.GetWithClientModelRepository[*struct{}, *struct{}]
		transactionRepository repository.TransactionRepository
		toRepoWhereInputFunc  ConverFunc[*struct{}, *struct{}]
		validateFunc          UpdateInTransactionValidateFunc[*struct{}, *struct{}, *struct{}]
		selectForUpdate       bool
		wrapGetErrorFunc      func(error) error
		wrapUpdateErrorFunc   func(error) error
	}
	type args struct {
		ctx         context.Context
		whereInput  *struct{}
		updateInput *struct{}
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	getRepo := repository.NewMockGetWithClientModelRepository[*struct{}, *struct{}](ctrl)
	updateRepo := repository.NewMockUpdateWithClientModelRepository[*struct{}, *struct{}](ctrl)

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

	updateRepo.EXPECT().UpdateWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Eq(new(struct{})), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	updateRepo.EXPECT().UpdateWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Eq(new(struct{})), gomock.Nil(),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	transactionRepository := repository.NewTransactionRepository(client)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *struct{}
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repository:            updateRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}, client *ent.Client) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: new(struct{}),
			},
			want: new(struct{}),
		},
		{
			name: "toRepoWhereInputFuncError",
			fields: fields{
				repository:            updateRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				validateFunc: func(ins *struct{}, uui *struct{}, client *ent.Client) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "GetError",
			fields: fields{
				repository:            updateRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}, client *ent.Client) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  nil,
				updateInput: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository:            updateRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}, client *ent.Client) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "UpdateError",
			fields: fields{
				repository:            updateRepo,
				getRepository:         getRepo,
				transactionRepository: transactionRepository,
				toRepoWhereInputFunc: func(uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(ins *struct{}, uui *struct{}, client *ent.Client) (*struct{}, error) {
					return uui, nil
				},
				wrapGetErrorFunc:    func(err error) error { return err },
				wrapUpdateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:         ctx,
				whereInput:  new(struct{}),
				updateInput: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &updateModelInTransactionUseCase[*struct{}, *struct{}, *struct{}, *struct{}, *struct{}]{
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

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createModelUseCase.Create() = %v, want %v", got, *tt.want)
			}
		})
	}
}
