package middleware

import (
	"context"
	"testing"

	"github.com/TcMits/ent-clean-template/ent"
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/require"
)

func newLoginUseCase(client *ent.Client) usecase.LoginUseCase[
	*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User] {
	return usecase.NewLoginUseCase(
		repository.NewLoginRepository(client),
		"Dummy",
	)
}

func TestAuth(t *testing.T) {
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	loginUseCase := newLoginUseCase(client)
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	jwtPayload, err := loginUseCase.Login(ctx, &useCaseModel.LoginInput{
		Username: u.Username,
		Password: "12345678",
	})
	require.NoError(t, err)

	handler := v1.NewHandler()
	handler.Use(Auth(loginUseCase))
	handler.Get("/test", func(ctx iris.Context) {
		userIrisContext, ok := ctx.Values().Get(UserKey).(lazy.LazyObject[*model.User])
		if !ok {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{})
			return
		}
		userRequestContext, ok := ctx.Request().Context().Value(UserKey).(lazy.LazyObject[*model.User])
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
		if user1.ID != user2.ID {
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
	e.GET("/test").WithHeader(AuthHeaderKey, JWTPrefix+" "+jwtPayload.AccessToken).Expect().Status(iris.StatusOK)
	e.GET("/test").WithQuery(JWTPrefix, jwtPayload.AccessToken).Expect().Status(iris.StatusOK)

}
