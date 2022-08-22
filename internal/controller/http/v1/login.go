package v1

import (
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type loginController[LoginInputType, RefreshTokenInputType any] struct {
	repository usecase.LoginUseCase[LoginInputType, any, RefreshTokenInputType, any]
	logger     logger.Interface
}

func RegisterLoginController[LoginInputType, RefreshTokenInputType any](
	handler iris.Party,
	repository usecase.LoginUseCase[LoginInputType, any, RefreshTokenInputType, any],
	logger logger.Interface,
) {
	controller := &loginController[LoginInputType, RefreshTokenInputType]{repository: repository, logger: logger}
	handler.Post("/login", controller.Login)
	handler.Post("/refresh-token", controller.RefreshToken)
	handler.Post("/verify-token", controller.VerifyToken)
}

type refreshTokenResponse struct {
	Token string `json:"token"`
}

type verifyTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

func (c *loginController[LoginInputType, _]) Login(ctx iris.Context) {
	loginInput := new(LoginInputType)
	if err := ctx.ReadBody(*loginInput); err != nil {
		HandleError(ctx, err, c.logger)
		return
	}
	authenticatedPayload, err := c.repository.Login(ctx.Request().Context(), *loginInput)
	if err != nil {
		HandleError(ctx, err, c.logger)
		return
	}
	ctx.JSON(authenticatedPayload)
}

func (c *loginController[_, RefreshTokenInputType]) RefreshToken(ctx iris.Context) {
	refreshTokenInput := new(RefreshTokenInputType)
	if err := ctx.ReadJSON(*refreshTokenInput); err != nil {
		HandleError(ctx, err, c.logger)
		return
	}
	token, err := c.repository.RefreshToken(ctx.Request().Context(), *refreshTokenInput)
	if err != nil {
		HandleError(ctx, err, c.logger)
		return
	}
	ctx.JSON(refreshTokenResponse{
		Token: token,
	})
}

func (c *loginController[_, _]) VerifyToken(ctx iris.Context) {
	verifyTokenInput := new(verifyTokenRequest)
	if err := ctx.ReadBody(*verifyTokenInput); err != nil {
		HandleError(ctx, err, c.logger)
		return
	}
	_, err := c.repository.VerifyToken(ctx.Request().Context(), verifyTokenInput.Token)
	if err != nil {
		HandleError(ctx, err, c.logger)
		return
	}
	ctx.JSON(iris.Map{})
}
