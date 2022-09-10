package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func Test_LoginHandler(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	l := &testutils.NullLogger{}
	u := usecase.NewMockLoginUseCase[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](ctrl)

	u.EXPECT().Login(
		gomock.Eq(ctx), gomock.Eq(&useCaseModel.LoginInput{
			Username: "tsolution",
			Password: "12345678",
		}),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	u.EXPECT().Login(
		gomock.Eq(ctx), gomock.Eq(&useCaseModel.LoginInput{
			Username: "tsolution",
			Password: "1234567",
		}),
	).Return(
		nil, useCaseModel.NewUseCaseError(
			errors.New(""), "test", "test", usecase.AuthenticationError,
		),
	).AnyTimes()

	handler := NewHandler()
	RegisterLoginController[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](handler, u, l)
	handler.Build()

	e := httptest.New(t, handler)
	e.POST(loginSubPath).WithForm(
		useCaseModel.LoginInput{
			Username: "tsolution",
			Password: "12345678",
		},
	).Expect().Status(iris.StatusOK)
	e.POST(loginSubPath).WithForm(
		useCaseModel.LoginInput{
			Username: "tsolution",
			Password: "1234567",
		},
	).Expect().Status(iris.StatusUnauthorized)
}

func Test_RefreshTokenHandler(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	l := &testutils.NullLogger{}
	u := usecase.NewMockLoginUseCase[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](ctrl)

	u.EXPECT().RefreshToken(
		gomock.Eq(ctx), gomock.Eq(&useCaseModel.RefreshTokenInput{
			RefreshToken: "test",
			RefreshKey:   "test",
		}),
	).Return(
		"token", nil,
	).AnyTimes()

	u.EXPECT().RefreshToken(
		gomock.Eq(ctx), gomock.Eq(&useCaseModel.RefreshTokenInput{
			RefreshToken: "tes",
			RefreshKey:   "test",
		}),
	).Return(
		"", useCaseModel.NewUseCaseError(
			errors.New(""), "test", "test", usecase.AuthenticationError,
		),
	).AnyTimes()

	handler := NewHandler()
	RegisterLoginController[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](handler, u, l)
	handler.Build()

	e := httptest.New(t, handler)

	e.POST(refreshTokenSubPath).WithJSON(
		useCaseModel.RefreshTokenInput{
			RefreshToken: "test",
			RefreshKey:   "test",
		},
	).Expect().Status(iris.StatusOK)
	e.POST(refreshTokenSubPath).WithJSON(
		useCaseModel.RefreshTokenInput{
			RefreshToken: "tes",
			RefreshKey:   "test",
		},
	).Expect().Status(iris.StatusUnauthorized)
}

func Test_VerifyTokenHandler(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	l := &testutils.NullLogger{}
	u := usecase.NewMockLoginUseCase[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](ctrl)

	u.EXPECT().VerifyToken(
		gomock.Eq(ctx), gomock.Eq("test"),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	u.EXPECT().VerifyToken(
		gomock.Eq(ctx), gomock.Eq("tes"),
	).Return(
		nil, useCaseModel.NewUseCaseError(
			errors.New(""), "test", "test", usecase.AuthenticationError,
		),
	).AnyTimes()

	handler := NewHandler()
	RegisterLoginController[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](handler, u, l)

	handler.Build()
	e := httptest.New(t, handler)
	e.POST(verifyTokenSubPath).WithForm(
		verifyTokenRequest{
			Token: "test",
		},
	).Expect().Status(iris.StatusOK)
	e.POST(verifyTokenSubPath).WithForm(
		verifyTokenRequest{
			Token: "tes",
		},
	).Expect().Status(iris.StatusUnauthorized)
}
