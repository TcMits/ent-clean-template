package v1

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

func Test_LoginHandler(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	l := &testutils.NullLogger{}
	u := usecase.NewMockLoginUseCase[
		*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User,
	](ctrl)

	u.EXPECT().Login(
		gomock.Eq(ctx), gomock.Eq(&useCaseModel.LoginInput{
			Username: "tsolution",
			Password: "12345678",
		}),
	).Return(
		new(useCaseModel.JWTAuthenticatedPayload), nil,
	).AnyTimes()

	u.EXPECT().Login(
		gomock.Eq(ctx), gomock.Eq(&useCaseModel.LoginInput{
			Username: "tsolution",
			Password: "1234567",
		}),
	).Return(
		nil, model.NewTranslatableError(
			errors.New(""), &i18n.Message{
				ID:    "test",
				Other: "test",
			}, usecase.AuthenticationError, nil,
		),
	).AnyTimes()

	handler := NewHandler()
	RegisterLoginController(handler, u, l)
	handler.Build()

	e := httptest.New(t, handler)
	e.POST(_loginSubPath).WithForm(
		useCaseModel.LoginInput{
			Username: "tsolution",
			Password: "12345678",
		},
	).Expect().Status(iris.StatusOK)
	e.POST(_loginSubPath).WithForm(
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
		*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User,
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
		"", model.NewTranslatableError(
			errors.New(""), &i18n.Message{
				ID:    "test",
				Other: "test",
			}, usecase.AuthenticationError, nil,
		),
	).AnyTimes()

	handler := NewHandler()
	RegisterLoginController(handler, u, l)
	handler.Build()

	e := httptest.New(t, handler)

	e.POST(_refreshTokenSubPath).WithJSON(
		useCaseModel.RefreshTokenInput{
			RefreshToken: "test",
			RefreshKey:   "test",
		},
	).Expect().Status(iris.StatusOK)
	e.POST(_refreshTokenSubPath).WithJSON(
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
		*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User,
	](ctrl)

	u.EXPECT().VerifyToken(
		gomock.Eq(ctx), gomock.Eq("test"),
	).Return(
		new(model.User), nil,
	).AnyTimes()

	u.EXPECT().VerifyToken(
		gomock.Eq(ctx), gomock.Eq("tes"),
	).Return(
		nil, model.NewTranslatableError(
			errors.New(""), &i18n.Message{
				ID:    "test",
				Other: "test",
			}, usecase.AuthenticationError, nil,
		),
	).AnyTimes()

	handler := NewHandler()
	RegisterLoginController(handler, u, l)

	handler.Build()
	e := httptest.New(t, handler)
	e.POST(_verifyTokenSubPath).WithForm(
		verifyTokenRequest{
			Token: "test",
		},
	).Expect().Status(iris.StatusOK)
	e.POST(_verifyTokenSubPath).WithForm(
		verifyTokenRequest{
			Token: "tes",
		},
	).Expect().Status(iris.StatusUnauthorized)
}
