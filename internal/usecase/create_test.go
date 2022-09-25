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

func Test_createModelUseCase_Create(t *testing.T) {
	type fields struct {
		repository          repository.CreateModelRepository[*struct{}, *struct{}]
		validateFunc        CreateValidateFunc[*struct{}, *struct{}]
		wrapCreateErrorFunc func(error) error
	}
	type args struct {
		ctx   context.Context
		input *struct{}
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockCreateModelRepository[*struct{}, *struct{}](ctrl)

	repo.EXPECT().Create(
		gomock.Eq(ctx), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	repo.EXPECT().Create(
		gomock.Eq(ctx), gomock.Nil(),
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
				repository: repo,
				validateFunc: func(c context.Context, uci *struct{}) (*struct{}, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			want: new(struct{}),
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository: repo,
				validateFunc: func(c context.Context, uci *struct{}) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "CreateError",
			fields: fields{
				repository: repo,
				validateFunc: func(c context.Context, uci *struct{}) (*struct{}, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
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
			l := &createModelUseCase[*struct{}, *struct{}, *struct{}]{
				repository:          tt.fields.repository,
				validateFunc:        tt.fields.validateFunc,
				wrapCreateErrorFunc: tt.fields.wrapCreateErrorFunc,
			}
			got, err := l.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("createModelUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createModelUseCase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createModelInTransactionUseCase_Create(t *testing.T) {
	type fields struct {
		repository            repository.CreateWithClientModelRepository[*struct{}, *struct{}]
		transactionRepository repository.TransactionRepository
		validateFunc          CreateInTransactionValidateFunc[*struct{}, *struct{}]
		wrapCreateErrorFunc   func(error) error
	}
	type args struct {
		ctx   context.Context
		input *struct{}
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockCreateWithClientModelRepository[*struct{}, *struct{}](ctrl)
	transactionRepository := repository.NewMockTransactionRepository(ctrl)

	transactionRepository.EXPECT().Start(
		gomock.Eq(ctx),
	).Return(
		nil, func() error { return nil }, func() error { return nil }, nil,
	).AnyTimes()

	repo.EXPECT().CreateWithClient(
		gomock.Eq(ctx), gomock.Nil(), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	repo.EXPECT().CreateWithClient(
		gomock.Eq(ctx), gomock.Nil(), gomock.Nil(),
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
				repository:            repo,
				transactionRepository: transactionRepository,
				validateFunc: func(c context.Context, uci *struct{}, client *ent.Client) (*struct{}, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			want: new(struct{}),
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				repository:            repo,
				transactionRepository: transactionRepository,
				validateFunc: func(c context.Context, uci *struct{}, client *ent.Client) (*struct{}, error) {
					return nil, errors.New("test")
				},
				wrapCreateErrorFunc: func(err error) error { return err },
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "CreateError",
			fields: fields{
				repository:            repo,
				transactionRepository: transactionRepository,
				validateFunc: func(c context.Context, uci *struct{}, client *ent.Client) (*struct{}, error) {
					return uci, nil
				},
				wrapCreateErrorFunc: func(err error) error { return err },
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
			l := &createModelInTransactionUseCase[*struct{}, *struct{}, *struct{}]{
				repository:            tt.fields.repository,
				transactionRepository: tt.fields.transactionRepository,
				validateFunc:          tt.fields.validateFunc,
				wrapCreateErrorFunc:   tt.fields.wrapCreateErrorFunc,
			}
			got, err := l.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"createModelInTransactionUseCase.Create() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createModelInTransactionUseCase.Create() = %v, want%v", got, tt.want)
			}
		})
	}
}

func Test_createModelHavingFileUseCase_Create(t *testing.T) {
	type fields struct {
		createUseCase       CreateModelUseCase[*struct{}, *struct{}]
		existFileRepository repository.ExistFileRepository
		writeFileRepository repository.WriteFileRepository
		validateFunc        CreateWithFileValidateFunc[*struct{}, *struct{}]
		writeFileTimeout    time.Duration
		l                   logger.Interface
	}
	type args struct {
		ctx   context.Context
		input *struct{}
	}
	// Create an SQLite memory database and generate the schema.
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	createModelUseCase := NewMockCreateModelUseCase[*struct{}, *struct{}](ctrl)
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

	createModelUseCase.EXPECT().Create(
		gomock.Eq(ctx), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	createModelUseCase.EXPECT().Create(
		gomock.Eq(ctx), gomock.Nil(),
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
				createUseCase:       createModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*CreateFile, error) {
					return new(struct{}), append(
						[]*CreateFile{},
						&CreateFile{Filename: "test.txt", Size: size, Reader: r},
					), nil
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			want: new(struct{}),
		},
		{
			name: "ValidateFuncError",
			fields: fields{
				createUseCase:       createModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*CreateFile, error) {
					return nil, nil, errors.New("")
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "CreateError",
			fields: fields{
				createUseCase:       createModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*CreateFile, error) {
					return nil, append(
						[]*CreateFile{},
						&CreateFile{Filename: "test.txt", Size: size, Reader: r},
					), nil
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			wantErr: true,
		},
		{
			name: "WriteError",
			fields: fields{
				createUseCase:       createModelUseCase,
				existFileRepository: existFileRepository,
				writeFileRepository: writeFileRepository,
				validateFunc: func(ctx context.Context, s *struct{}) (*struct{}, []*CreateFile, error) {
					return new(struct{}), append(
						[]*CreateFile{},
						&CreateFile{Filename: "test2.txt", Size: size, Reader: r},
					), nil
				},
				writeFileTimeout: 15 * time.Minute,
				l:                testutils.NullLogger{},
			},
			args: args{
				ctx:   ctx,
				input: new(struct{}),
			},
			want: new(struct{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &createModelHavingFileUseCase[*struct{}, *struct{}, *struct{}]{
				createUseCase:       tt.fields.createUseCase,
				existFileRepository: tt.fields.existFileRepository,
				writeFileRepository: tt.fields.writeFileRepository,
				validateFunc:        tt.fields.validateFunc,
				writeFileTimeout:    tt.fields.writeFileTimeout,
				l:                   tt.fields.l,
			}
			got, err := l.Create(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"createModelHavingFileUseCase.Create() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createModelHavingFileUseCase.Create() = %v, want%v", got, tt.want)
			}
		})
	}
}
