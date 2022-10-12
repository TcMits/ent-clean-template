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

func Test_getDeleteHandler(t *testing.T) {
	type args struct {
		getAndDeleteUseCase usecase.GetAndDeleteModelUseCase[*struct{}, *MockWhereInput]
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
	notWantErr := false
	serializeUseCase := usecase.NewMockSerializeModelUseCase[*struct{}, *struct{}](ctrl)
	getAndDeleteUseCase := usecase.NewMockGetAndDeleteModelUseCase[*struct{}, *MockWhereInput](ctrl)

	getAndDeleteUseCase.EXPECT().GetAndDelete(
		gomock.Eq(ctx),
		gomock.Eq(new(MockWhereInput)),
	).Return(
		nil,
	).AnyTimes()

	getAndDeleteUseCase.EXPECT().GetAndDelete(
		gomock.Eq(ctx),
		gomock.Eq(&MockWhereInput{WantError: &notWantErr}),
	).Return(
		nil,
	).AnyTimes()

	getAndDeleteUseCase.EXPECT().GetAndDelete(
		gomock.Eq(ctx),
		gomock.Eq(&MockWhereInput{WantError: &wantErr}),
	).Return(
		errors.New(""),
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
			name: "Delete",
			args: args{
				getAndDeleteUseCase: getAndDeleteUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				whereInput:          new(MockWhereInput),
			},
		},
		{
			name: "DeleteError",
			args: args{
				getAndDeleteUseCase: getAndDeleteUseCase,
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
			got := getDeleteHandler(
				tt.args.getAndDeleteUseCase,
				tt.args.serializeUseCase,
				tt.args.l,
				tt.args.wrapReadParamsError,
				tt.args.wrapReadQueryError,
			)

			handler := NewHandler()
			handler.Delete("/test", got)
			e := httptest.New(t, handler)

			req := e.DELETE("/test")
			if tt.args.whereInput.WantError != nil {
				req = req.WithQuery("want_error", *tt.args.whereInput.WantError)
			}
			req.Expect()
		})
	}
}
