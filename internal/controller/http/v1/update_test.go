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

type MockUpdateInput struct {
	WantError *bool `json:"want_error" form:"want_error"`
}

func Test_getUpdateHandler(t *testing.T) {
	type args struct {
		getAndUpdateUseCase usecase.GetAndUpdateModelUseCase[*struct{}, *struct{}, *MockUpdateInput]
		serializeUseCase    usecase.SerializeModelUseCase[*struct{}, *struct{}]
		l                   logger.Interface
		wrapReadParamsError func(error) error
		wrapReadQueryError  func(error) error
		wrapReadBodyError   func(error) error
		updateInput         *MockUpdateInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	wantErr := true
	notWantErr := false
	serializeUseCase := usecase.NewMockSerializeModelUseCase[*struct{}, *struct{}](ctrl)
	getAndUpdateUseCase := usecase.NewMockGetAndUpdateModelUseCase[*struct{}, *struct{}, *MockUpdateInput](
		ctrl,
	)

	getAndUpdateUseCase.EXPECT().GetAndUpdate(
		gomock.Eq(ctx),
		gomock.Eq(new(struct{})),
		gomock.Eq(new(MockUpdateInput)),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	getAndUpdateUseCase.EXPECT().GetAndUpdate(
		gomock.Eq(ctx),
		gomock.Eq(new(struct{})),
		gomock.Eq(&MockUpdateInput{WantError: &notWantErr}),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	getAndUpdateUseCase.EXPECT().GetAndUpdate(
		gomock.Eq(ctx),
		gomock.Eq(new(struct{})),
		gomock.Eq(&MockUpdateInput{WantError: &wantErr}),
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
			name: "GetAndUpdate",
			args: args{
				getAndUpdateUseCase: getAndUpdateUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				wrapReadBodyError:   func(err error) error { return errors.New("") },
				updateInput:         new(MockUpdateInput),
			},
		},
		{
			name: "GetAndUpdateError",
			args: args{
				getAndUpdateUseCase: getAndUpdateUseCase,
				serializeUseCase:    serializeUseCase,
				l:                   testutils.NullLogger{},
				wrapReadParamsError: func(err error) error { return errors.New("") },
				wrapReadQueryError:  func(err error) error { return errors.New("") },
				wrapReadBodyError:   func(err error) error { return errors.New("") },
				updateInput:         &MockUpdateInput{WantError: &wantErr},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getUpdateHandler(
				tt.args.getAndUpdateUseCase,
				tt.args.serializeUseCase,
				tt.args.l,
				tt.args.wrapReadParamsError,
				tt.args.wrapReadQueryError,
				tt.args.wrapReadBodyError,
			)

			handler := NewHandler()
			handler.Put("/test-put", got)
			handler.Patch("/test-patch", got)
			e := httptest.New(t, handler)

			req := e.PUT("/test-put")
			req.WithForm(tt.args.updateInput).Expect()

			req = e.PATCH("/test-patch")
			req.WithForm(tt.args.updateInput).Expect()
		})
	}
}
