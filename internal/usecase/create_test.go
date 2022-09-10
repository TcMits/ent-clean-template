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
				validateFunc: func(uci *struct{}) (*struct{}, error) {
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
				validateFunc: func(uci *struct{}) (*struct{}, error) {
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
				validateFunc: func(uci *struct{}) (*struct{}, error) {
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
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := repository.NewMockCreateWithClientModelRepository[*struct{}, *struct{}](ctrl)
	transactionRepository := repository.NewTransactionRepository(client)

	repo.EXPECT().CreateWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Eq(new(struct{})),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	repo.EXPECT().CreateWithClient(
		gomock.Eq(ctx), gomock.Any(), gomock.Nil(),
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
				validateFunc: func(uci *struct{}, client *ent.Client) (*struct{}, error) {
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
				validateFunc: func(uci *struct{}, client *ent.Client) (*struct{}, error) {
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
				validateFunc: func(uci *struct{}, client *ent.Client) (*struct{}, error) {
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
				t.Errorf("createModelInTransactionUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createModelInTransactionUseCase.Create() = %v, want%v", got, tt.want)
			}
		})
	}
}
