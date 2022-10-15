package usecase

import (
	"context"
	"errors"
	"io"
	"math/rand"
	"reflect"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"go.beyondstorage.io/v5/pkg/randbytes"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

func Test_updateModelUseCase_Update(t *testing.T) {
	type fields struct {
		getUseCase          GetModelUseCase[*struct{}, *struct{}]
		repository          repository.UpdateModelRepository[*struct{}, *struct{}]
		validateFunc        UpdateValidateFunc[*struct{}, *struct{}, *struct{}]
		wrapUpdateErrorFunc func(error) error
	}
	type args struct {
		ctx         context.Context
		whereInput  *struct{}
		updateInput *struct{}
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getUseCase := NewMockGetModelUseCase[*struct{}, *struct{}](ctrl)
	updateRepo := repository.NewMockUpdateModelRepository[*struct{}, *struct{}](ctrl)

	getUseCase.EXPECT().Get(
		gomock.Eq(ctx), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	getUseCase.EXPECT().Get(
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
				getUseCase: getUseCase,
				repository: updateRepo,
				validateFunc: func(c context.Context, ins, uui *struct{}) (*struct{}, error) {
					return uui, nil
				},
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
			name: "GetError",
			fields: fields{
				getUseCase: getUseCase,
				repository: updateRepo,
				validateFunc: func(c context.Context, ins, uui *struct{}) (*struct{}, error) {
					return uui, nil
				},
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
				getUseCase: getUseCase,
				repository: updateRepo,
				validateFunc: func(c context.Context, ins, uui *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
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
				getUseCase: getUseCase,
				repository: updateRepo,
				validateFunc: func(c context.Context, ins, uui *struct{}) (*struct{}, error) {
					return uui, nil
				},
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
			l := &updateModelUseCase[*struct{}, *struct{}, *struct{}, *struct{}]{
				getUseCase:          tt.fields.getUseCase,
				repository:          tt.fields.repository,
				validateFunc:        tt.fields.validateFunc,
				wrapUpdateErrorFunc: tt.fields.wrapUpdateErrorFunc,
			}
			got, err := l.GetAndUpdate(tt.args.ctx, tt.args.whereInput, tt.args.updateInput)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"getAndUpdateModelUseCase.GetAndUpdate() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAndUpdateModelUseCase.GetAndUpdate() = %v, want %v", got, *tt.want)
			}
		})
	}
}

func Test_updateModelInTransactionUseCase_Update(t *testing.T) {
	type fields struct {
		repository            repository.UpdateWithClientModelRepository[*struct{}, *struct{}]
		getRepository         repository.GetWithClientModelRepository[*struct{}, *struct{}]
		transactionRepository repository.TransactionRepository
		toRepoWhereInputFunc  ConvertFunc[*struct{}, *struct{}]
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
	// GetAndUpdate an SQLite memory database and generate the schema.
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getRepo := repository.NewMockGetWithClientModelRepository[*struct{}, *struct{}](ctrl)
	updateRepo := repository.NewMockUpdateWithClientModelRepository[*struct{}, *struct{}](ctrl)
	transactionRepository := repository.NewMockTransactionRepository(ctrl)

	transactionRepository.EXPECT().Start(
		gomock.Eq(ctx),
	).Return(
		nil, func() error { return nil }, func() error { return nil }, nil,
	).AnyTimes()

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
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(c context.Context, ins, uui *struct{}, client *ent.Client) (*struct{}, error) {
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
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				validateFunc: func(c context.Context, ins, uui *struct{}, client *ent.Client) (*struct{}, error) {
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
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(c context.Context, ins, uui *struct{}, client *ent.Client) (*struct{}, error) {
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
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(c context.Context, ins, uui *struct{}, client *ent.Client) (*struct{}, error) {
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
				toRepoWhereInputFunc: func(c context.Context, uwi *struct{}) (*struct{}, error) {
					return uwi, nil
				},
				validateFunc: func(c context.Context, ins, uui *struct{}, client *ent.Client) (*struct{}, error) {
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
				t.Errorf(
					"getAndUpdateModelUseCase.GetAndUpdate() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAndUpdateModelUseCase.GetAndUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAndUpdateModelWithFileUseCase_GetAndUpdate(t *testing.T) {
	type fields struct {
		getAndUpdateUseCase GetAndUpdateModelUseCase[*struct{}, *struct{}, *struct{}]
		existFileRepository repository.ExistFileRepository
		writeFileRepository repository.WriteFileRepository
		validateFunc        UpdateWithFileValidateFunc[*struct{}, *struct{}]
		writeFileTimeout    time.Duration
		l                   logger.Interface
	}
	type args struct {
		ctx         context.Context
		updateInput *struct{}
		whereInput  *struct{}
	}
	// GetAndUpdate an SQLite memory database and generate the schema.
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	getAndUpdateModelUseCase := NewMockGetAndUpdateModelUseCase[*struct{}, *struct{}, *struct{}](
		ctrl,
	)
	existFileRepository := repository.NewMockExistFileRepository(ctrl)
	writeFileRepository := repository.NewMockWriteFileRepository(ctrl)
	size := rand.Int63n(4 * 1024 * 1024)
	r := io.LimitReader(randbytes.NewRand(), size)

	writeFileRepository.EXPECT().Write(
		gomock.Any(), gomock.Eq("test.txt"), gomock.Eq(r), gomock.Eq(size),
	).Return(
		size, nil,
	).AnyTimes()

	writeFileRepository.EXPECT().Write(
		gomock.Any(), gomock.Eq("hello-world/test.txt"), gomock.Eq(r), gomock.Eq(size),
	).Return(
		size, nil,
	).AnyTimes()

	writeFileRepository.EXPECT().Write(
		gomock.Any(), gomock.Eq("test2.txt"), gomock.Eq(r), gomock.Eq(size),
	).Return(
		int64(0), errors.New(""),
	).AnyTimes()

	getAndUpdateModelUseCase.EXPECT().GetAndUpdate(
		gomock.Eq(ctx), gomock.Eq(new(struct{})), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	getAndUpdateModelUseCase.EXPECT().GetAndUpdate(
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
				getAndUpdateUseCase: getAndUpdateModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*UpdateFile, error) {
					return new(struct{}), append(
						[]*UpdateFile{},
						&UpdateFile{Filename: "test.txt", Size: size, Reader: r},
					), nil
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:         ctx,
				updateInput: new(struct{}),
				whereInput:  new(struct{}),
			},
			want: new(struct{}),
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				getAndUpdateUseCase: getAndUpdateModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*UpdateFile, error) {
					return nil, nil, errors.New("")
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:         ctx,
				updateInput: new(struct{}),
				whereInput:  new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "GetAndUpdateError",
			fields: fields{
				getAndUpdateUseCase: getAndUpdateModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*UpdateFile, error) {
					return nil, append(
						[]*UpdateFile{},
						&UpdateFile{Filename: "test.txt", Size: size, Reader: r},
					), nil
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:         ctx,
				updateInput: new(struct{}),
				whereInput:  new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "WriteError",
			fields: fields{
				getAndUpdateUseCase: getAndUpdateModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*UpdateFile, error) {
					return new(struct{}), append(
						[]*UpdateFile{},
						&UpdateFile{Filename: "test2.txt", Size: size, Reader: r},
					), nil
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:         ctx,
				updateInput: new(struct{}),
				whereInput:  new(struct{}),
			},
			want: new(struct{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &getAndUpdateModelWithFileUseCase[*struct{}, *struct{}, *struct{}, *struct{}]{
				getAndUpdateUseCase: tt.fields.getAndUpdateUseCase,
				existFileRepository: tt.fields.existFileRepository,
				writeFileRepository: tt.fields.writeFileRepository,
				validateFunc:        tt.fields.validateFunc,
				writeFileTimeout:    tt.fields.writeFileTimeout,
				l:                   tt.fields.l,
			}
			got, err := l.GetAndUpdate(tt.args.ctx, tt.args.whereInput, tt.args.updateInput)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"getAndUpdateModelHavingFileUseCase.GetAndUpdate() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"getAndUpdateModelHavingFileUseCase.GetAndUpdate() = %v, want%v",
					got,
					tt.want,
				)
			}
		})
	}
}
