package v1

import (
	"context"
	"testing"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/testutils"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/factory"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
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

func Test_loginController_Login(t *testing.T) {
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	loginUseCase := newLoginUseCase(client)
	l := &testutils.NullLogger{}
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)

	handler := NewHandler()
	RegisterLoginController(handler, loginUseCase, l)
	handler.Build()

	e := httptest.New(t, handler)
	e.POST(loginSubPath).WithForm(
		useCaseModel.LoginInput{
			Username: u.Username,
			Password: "12345678",
		},
	).Expect().Status(iris.StatusOK)
	e.POST(loginSubPath).WithForm(
		useCaseModel.LoginInput{
			Username: u.Username,
			Password: "1234567",
		},
	).Expect().Status(iris.StatusUnauthorized)
	e.POST(loginSubPath).WithForm(
		useCaseModel.LoginInput{
			Username: u.Username + "wrong",
			Password: "12345678",
		},
	).Expect().Status(iris.StatusUnauthorized)
}

func Test_loginController_RefreshToken(t *testing.T) {
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	loginUseCase := newLoginUseCase(client)
	l := &testutils.NullLogger{}
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	jwtPayload, err := loginUseCase.Login(ctx, &useCaseModel.LoginInput{
		Username: u.Username,
		Password: "12345678",
	})
	require.NoError(t, err)

	handler := NewHandler()
	RegisterLoginController(handler, loginUseCase, l)
	e := httptest.New(t, handler)
	e.POST(refreshTokenSubPath).WithJSON(
		useCaseModel.RefreshTokenInput{
			RefreshToken: jwtPayload.RefreshToken,
			RefreshKey:   jwtPayload.RefreshKey,
		},
	).Expect().Status(iris.StatusOK)
	e.POST(refreshTokenSubPath).WithJSON(
		useCaseModel.RefreshTokenInput{
			RefreshToken: jwtPayload.RefreshToken + "wrong",
			RefreshKey:   jwtPayload.RefreshKey,
		},
	).Expect().Status(iris.StatusUnauthorized)
	e.POST(refreshTokenSubPath).WithJSON(
		useCaseModel.RefreshTokenInput{
			RefreshToken: jwtPayload.RefreshToken,
			RefreshKey:   jwtPayload.RefreshKey + "wrong",
		},
	).Expect().Status(iris.StatusUnauthorized)
}

func Test_loginController_VerifyToken(t *testing.T) {
	ctx := context.Background()
	client := testutils.GetSqlite3TestClient(ctx, t)
	defer client.Close()
	loginUseCase := newLoginUseCase(client)
	l := &testutils.NullLogger{}
	u, err := factory.UserFactory.Create(ctx, client.User.Create(), map[string]any{})
	require.NoError(t, err)
	jwtPayload, err := loginUseCase.Login(ctx, &useCaseModel.LoginInput{
		Username: u.Username,
		Password: "12345678",
	})
	require.NoError(t, err)

	handler := NewHandler()
	RegisterLoginController(handler, loginUseCase, l)
	e := httptest.New(t, handler)
	e.POST(verifyTokenSubPath).WithForm(
		verifyTokenRequest{
			Token: jwtPayload.AccessToken,
		},
	).Expect().Status(iris.StatusOK)
	e.POST(verifyTokenSubPath).WithForm(
		verifyTokenRequest{
			Token: jwtPayload.AccessToken + "wrong",
		},
	).Expect().Status(iris.StatusUnauthorized)
}
