package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12/httptest"

	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

type MockWhereInput struct {
	ID             *string `param:"id"`
	WantError      *bool   `           url:"want_error"`
	WantCountError *bool   `           url:"want_count_error"`
}

func Test_getDetailHandler(t *testing.T) {
	type args struct {
		getUseCase          usecase.GetModelUseCase[*struct{}, *MockWhereInput]
		serializeUseCase    usecase.SerializeModelUseCase[*struct{}, *struct{}]
		l                   logger.Interface
		wrapReadParamsError func(error) error
		wrapReadQueryError  func(error) error
		whereInput          *MockWhereInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	wantErr := true
	getUseCase := usecase.NewMockGetModelUseCase[*struct{}, *MockWhereInput](ctrl)
	serializeUseCase := usecase.NewMockSerializeModelUseCase[*struct{}, *struct{}](ctrl)

	getUseCase.EXPECT().Get(
		gomock.Eq(ctx),
		gomock.Eq(new(MockWhereInput)),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	getUseCase.EXPECT().Get(
		gomock.Eq(ctx),
		gomock.Eq(&MockWhereInput{WantError: &wantErr}),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	serializeUseCase.EXPECT().Serialize(
		gomock.Eq(ctx),
		gomock.Eq(new(struct{})),
	).Return(new(struct{})).AnyTimes()

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get",
			args: args{
				getUseCase:          getUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				whereInput:          new(MockWhereInput),
			},
		},
		{
			name: "GetError",
			args: args{
				getUseCase:          getUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				whereInput:          &MockWhereInput{WantError: &wantErr},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDetailHandler(
				tt.args.getUseCase,
				tt.args.serializeUseCase,
				tt.args.l,
				tt.args.wrapReadParamsError,
				tt.args.wrapReadQueryError,
			)

			handler := NewHandler()
			handler.Get("/test", got)
			e := httptest.New(t, handler)

			req := e.GET("/test")
			if tt.args.whereInput.WantError != nil {
				req = req.WithQuery("want_error", *tt.args.whereInput.WantError)
			}
			if tt.args.whereInput.ID != nil {
				req = req.WithPath("id", *tt.args.whereInput.ID)
			}
			req.Expect()
		})
	}
}

func Test_getListHandler(t *testing.T) {
	type args struct {
		listUseCase         usecase.ListModelUseCase[*struct{}, *struct{}, *MockWhereInput]
		countUseCase        usecase.CountModelUseCase[*MockWhereInput]
		serializeUseCase    usecase.SerializeModelUseCase[*struct{}, *struct{}]
		l                   logger.Interface
		wrapReadParamsError func(error) error
		wrapReadQueryError  func(error) error
		whereInput          *MockWhereInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	wantErr := true
	listUseCase := usecase.NewMockListModelUseCase[*struct{}, *struct{}, *MockWhereInput](ctrl)
	serializeUseCase := usecase.NewMockSerializeModelUseCase[*struct{}, *struct{}](ctrl)
	countUseCase := usecase.NewMockCountModelUseCase[*MockWhereInput](ctrl)
	eleven := 11
	one := 1

	countUseCase.EXPECT().Count(
		gomock.Eq(ctx),
		gomock.Eq(new(MockWhereInput)),
	).Return(
		20, nil,
	).AnyTimes()

	countUseCase.EXPECT().Count(
		gomock.Eq(ctx),
		gomock.Eq(&MockWhereInput{WantCountError: &wantErr}),
	).Return(
		0, errors.New(""),
	).AnyTimes()

	listUseCase.EXPECT().List(
		gomock.Eq(ctx),
		gomock.Eq(&eleven),
		gomock.Eq(&one),
		gomock.Eq(new(struct{})),
		gomock.Eq(new(MockWhereInput)),
	).Return(
		make([]*struct{}, 11), nil,
	).AnyTimes()

	listUseCase.EXPECT().List(
		gomock.Eq(ctx),
		gomock.Eq(&eleven),
		gomock.Eq(&one),
		gomock.Eq(new(struct{})),
		gomock.Eq(&MockWhereInput{WantCountError: &wantErr}),
	).Return(
		make([]*struct{}, 11), nil,
	).AnyTimes()

	listUseCase.EXPECT().List(
		gomock.Eq(ctx),
		gomock.Eq(&eleven),
		gomock.Eq(&one),
		gomock.Eq(new(struct{})),
		gomock.Eq(&MockWhereInput{WantError: &wantErr}),
	).Return(
		nil, errors.New(""),
	).AnyTimes()

	serializeUseCase.EXPECT().Serialize(
		gomock.Eq(ctx),
		gomock.Nil(),
	).Return(new(struct{})).AnyTimes()

	tests := []struct {
		name string
		args args
	}{
		{
			name: "List",
			args: args{
				listUseCase:         listUseCase,
				countUseCase:        countUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				whereInput:          new(MockWhereInput),
			},
		},
		{
			name: "ListError",
			args: args{
				listUseCase:         listUseCase,
				countUseCase:        countUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				whereInput:          &MockWhereInput{WantError: &wantErr},
			},
		},
		{
			name: "CountError",
			args: args{
				listUseCase:         listUseCase,
				countUseCase:        countUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				whereInput:          &MockWhereInput{WantCountError: &wantErr},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getListHandler(
				tt.args.listUseCase,
				tt.args.countUseCase,
				tt.args.serializeUseCase,
				tt.args.l,
				tt.args.wrapReadParamsError,
				tt.args.wrapReadQueryError,
			)

			handler := NewHandler()
			handler.Get("/test", got)
			e := httptest.New(t, handler)

			req := e.GET("/test")
			if tt.args.whereInput.WantError != nil {
				req = req.WithQuery("want_error", *tt.args.whereInput.WantError)
			}
			if tt.args.whereInput.WantCountError != nil {
				req = req.WithQuery("want_count_error", *tt.args.whereInput.WantCountError)
			}
			if tt.args.whereInput.ID != nil {
				req = req.WithPath("id", *tt.args.whereInput.ID)
			}
			req.WithQuery("limit", 10).WithQuery("offset", 1).WithQuery("with_count", true).Expect()
		})
	}
}
