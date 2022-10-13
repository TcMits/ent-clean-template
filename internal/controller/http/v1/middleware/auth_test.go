package middleware

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
)

func Test_Auth(t *testing.T) {
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

	handler := iris.New()
	handler.Use(Auth[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](u))
	handler.Get("/test", func(ctx iris.Context) {
		userIrisContext, ok := ctx.Values().Get(_userKey).(lazy.LazyObject[*struct{}])
		if !ok {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{})
			return
		}
		userRequestContext, ok := ctx.Request().Context().Value(_userKey).(lazy.LazyObject[*struct{}])
		if !ok {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{})
			return
		}
		user1 := userIrisContext.Value()
		user2 := userRequestContext.Value()
		if user1 == nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{})
			return
		}
		if user2 == nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{})
			return
		}
		if !reflect.DeepEqual(user1, user2) {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{})
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{})
	})
	handler.Build()

	e := httptest.New(t, handler)

	e.GET("/test").Expect().Status(iris.StatusUnauthorized)
	e.GET("/test").
		WithHeader(_authHeaderKey, _JWTPrefix+" "+"tes").
		Expect().
		Status(iris.StatusUnauthorized)
	e.GET("/test").WithQuery(_JWTPrefix, "tes").Expect().Status(iris.StatusUnauthorized)
	e.GET("/test").WithHeader(_authHeaderKey, _JWTPrefix+" "+"test").Expect().Status(iris.StatusOK)
	e.GET("/test").WithQuery(_JWTPrefix, "test").Expect().Status(iris.StatusOK)
}
