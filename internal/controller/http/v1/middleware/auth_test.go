package middleware

import (
	"context"
	"errors"
	"reflect"
	"testing"

	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
	"github.com/golang/mock/gomock"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
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
		nil, useCaseModel.NewUseCaseError(
			errors.New(""), "test", "test", usecase.AuthenticationError,
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
		nil, useCaseModel.NewUseCaseError(
			errors.New(""), "test", "test", usecase.AuthenticationError,
		),
	).AnyTimes()

	handler := v1.NewHandler()
	handler.Use(Auth[
		*useCaseModel.LoginInput, *struct{}, *useCaseModel.RefreshTokenInput, *struct{},
	](u))
	handler.Get("/test", func(ctx iris.Context) {
		userIrisContext, ok := ctx.Values().Get(UserKey).(lazy.LazyObject[*struct{}])
		if !ok {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{})
			return
		}
		userRequestContext, ok := ctx.Request().Context().Value(UserKey).(lazy.LazyObject[*struct{}])
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
	e.GET("/test").WithHeader(AuthHeaderKey, JWTPrefix+" "+"tes").Expect().Status(iris.StatusUnauthorized)
	e.GET("/test").WithQuery(JWTPrefix, "tes").Expect().Status(iris.StatusUnauthorized)
	e.GET("/test").WithHeader(AuthHeaderKey, JWTPrefix+" "+"test").Expect().Status(iris.StatusOK)
	e.GET("/test").WithQuery(JWTPrefix, "test").Expect().Status(iris.StatusOK)

}