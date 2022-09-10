package middleware

import (
	"context"
	"errors"
	"testing"

	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func Test_Permission(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	u := usecase.NewMockLoginUseCase[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](ctrl)

	u.EXPECT().VerifyToken(
		gomock.Eq(ctx), gomock.Eq(""),
	).Return(
		nil, useCaseModel.NewUseCaseError(
			errors.New(""), "test", "test", usecase.AuthenticationError,
		),
	).AnyTimes()

	u.EXPECT().VerifyToken(
		gomock.Eq(ctx), gomock.Eq("test"),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	handler := v1.NewHandler()
	handler.Use(Auth[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](u))
	handler.Use(Permission(testutils.NullLogger{}, usecase.NewDisallowZeroPermissionChecker[*struct{}](
		useCaseModel.NewUseCaseError(
			errors.New(""), "test", "test", usecase.PermissionDeniedError,
		),
	)))
	handler.Get("/test", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{})
	})
	handler.Build()

	e := httptest.New(t, handler)

	e.GET("/test").Expect().Status(iris.StatusForbidden)
	e.GET("/test").WithHeader(AuthHeaderKey, JWTPrefix+" "+"test").Expect().Status(iris.StatusOK)

}
