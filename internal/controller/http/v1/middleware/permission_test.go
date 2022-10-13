package middleware

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
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
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
		nil, model.NewTranslatableError(
			errors.New(""), &i18n.Message{
				ID:    "test",
				Other: "test",
			}, usecase.AuthenticationError, nil,
		),
	).AnyTimes()

	u.EXPECT().VerifyToken(
		gomock.Eq(ctx), gomock.Eq("test"),
	).Return(
		new(struct{}), nil,
	).AnyTimes()

	handler := iris.New()
	handler.Use(Auth[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](u))
	handler.Use(Permission(
		func(ctx iris.Context, err error, i logger.Interface) {
			ctx.StopWithJSON(iris.StatusForbidden, iris.Map{})
		},
		testutils.NullLogger{}, usecase.NewDisallowZeroPermissionChecker[*struct{}](
			model.NewTranslatableError(
				errors.New(""), &i18n.Message{
					ID:    "test",
					Other: "test",
				}, usecase.AuthenticationError, nil,
			),
		)))
	handler.Get("/test", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{})
	})
	handler.Build()

	e := httptest.New(t, handler)

	e.GET("/test").Expect().Status(iris.StatusForbidden)
	e.GET("/test").WithHeader(_authHeaderKey, _JWTPrefix+" "+"test").Expect().Status(iris.StatusOK)
}
