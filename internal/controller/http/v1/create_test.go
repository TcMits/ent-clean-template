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

type MockCreateInput struct {
	WantErr *bool `form:"want_error" json:"want_error"`
}

func Test_getCreateHandler(t *testing.T) {
	type args struct {
		createUseCase     usecase.CreateModelUseCase[*struct{}, *MockCreateInput]
		serializeUseCase  usecase.SerializeModelUseCase[*struct{}, *struct{}]
		l                 logger.Interface
		wrapReadBodyError func(error) error
		createInput       *MockCreateInput
	}

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	wantErr := true
	notWantErr := false
	createUseCase := usecase.NewMockCreateModelUseCase[*struct{}, *MockCreateInput](ctrl)
	serializeUseCase := usecase.NewMockSerializeModelUseCase[*struct{}, *struct{}](ctrl)

	createUseCase.EXPECT().Create(
		gomock.Eq(ctx),
		gomock.Eq(new(MockCreateInput)),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	createUseCase.EXPECT().Create(
		gomock.Eq(ctx),
		gomock.Eq(&MockCreateInput{WantErr: &notWantErr}),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	createUseCase.EXPECT().Create(
		gomock.Eq(ctx),
		gomock.Eq(&MockCreateInput{WantErr: &wantErr}),
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
			name: "Create",
			args: args{
				createUseCase:     createUseCase,
				serializeUseCase:  serializeUseCase,
				l:                 testutils.NullLogger{},
				wrapReadBodyError: func(err error) error { return errors.New("") },
				createInput:       new(MockCreateInput),
			},
		},
		{
			name: "CreateError",
			args: args{
				createUseCase:     createUseCase,
				serializeUseCase:  serializeUseCase,
				l:                 testutils.NullLogger{},
				wrapReadBodyError: func(err error) error { return errors.New("") },
				createInput:       &MockCreateInput{WantErr: &wantErr},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCreateHandler(
				tt.args.createUseCase,
				tt.args.serializeUseCase,
				tt.args.l,
				tt.args.wrapReadBodyError,
			)

			handler := NewHandler()
			handler.Post("/test", got)
			e := httptest.New(t, handler)
			e.POST("/test").WithForm(tt.args.createInput).Expect()
		})
	}
}
