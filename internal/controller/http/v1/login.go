package v1

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

var (
	_wrapInvalidLoginInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("loginController - Login - ctx.ReadBody: %w", err),
			defaultInvalidErrorTranslateKey,
			translationFunc,
			defaultInvalidErrorMessage,
			UscaseInputValidationError,
		)
	}
	_wrapInvalidRefreshInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("loginController - RefreshToken - ctx.ReadBody: %w", err),
			defaultInvalidErrorTranslateKey,
			translationFunc,
			defaultInvalidErrorMessage,
			UscaseInputValidationError,
		)
	}
	_wrapInvalidVerifyTokenInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("loginController - VerifyToken - ctx.ReadBody: %w", err),
			defaultInvalidErrorTranslateKey,
			translationFunc,
			defaultInvalidErrorMessage,
			UscaseInputValidationError,
		)
	}
)

type loginController[
	LoginInputType, AuthenticatedPayloadType, RefreshTokenInputType, UserType any] struct {
	useCase usecase.LoginUseCase[LoginInputType, AuthenticatedPayloadType, RefreshTokenInputType, UserType]
	logger  logger.Interface
}

func RegisterLoginController[
	LoginInputType, AuthenticatedPayloadType, RefreshTokenInputType, UserType any](
	handler iris.Party,
	useCase usecase.LoginUseCase[LoginInputType, AuthenticatedPayloadType, RefreshTokenInputType, UserType],
	logger logger.Interface,
) {
	controller := &loginController[
		LoginInputType, AuthenticatedPayloadType, RefreshTokenInputType, UserType,
	]{useCase: useCase, logger: logger}
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

func (c *loginController[LoginInputType, _, _, _]) Login(ctx iris.Context) {
	loginInput := new(LoginInputType)
	if err := ctx.ReadBody(loginInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			err = translatableErrorFromValidationErrors(
				loginInput, errs, ctx.Tr,
			)
		} else {
			err = _wrapInvalidLoginInput(ctx.Tr, err)
		}
		handleError(ctx, err, c.logger)
		return
	}
	authenticatedPayload, err := c.useCase.Login(ctx.Request().Context(), loginInput)
	if err != nil {
		handleError(ctx, err, c.logger)
		return
	}
	ctx.JSON(authenticatedPayload)
}

func (c *loginController[_, _, RefreshTokenInputType, _]) RefreshToken(ctx iris.Context) {
	refreshTokenInput := new(RefreshTokenInputType)
	if err := ctx.ReadJSON(refreshTokenInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			err = translatableErrorFromValidationErrors(
				refreshTokenInput, errs, ctx.Tr,
			)
		} else {
			err = _wrapInvalidRefreshInput(ctx.Tr, err)
		}
		handleError(ctx, err, c.logger)
		return
	}
	token, err := c.useCase.RefreshToken(ctx.Request().Context(), refreshTokenInput)
	if err != nil {
		handleError(ctx, err, c.logger)
		return
	}
	ctx.JSON(refreshTokenResponse{Token: token})
}

func (c *loginController[_, _, _, _]) VerifyToken(ctx iris.Context) {
	verifyTokenInput := new(verifyTokenRequest)
	if err := ctx.ReadBody(verifyTokenInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			err = translatableErrorFromValidationErrors(
				verifyTokenInput, errs, ctx.Tr,
			)
		} else {
			err = _wrapInvalidVerifyTokenInput(ctx.Tr, err)
		}
		handleError(ctx, err, c.logger)
		return
	}
	_, err := c.useCase.VerifyToken(ctx.Request().Context(), verifyTokenInput.Token)
	if err != nil {
		handleError(ctx, err, c.logger)
		return
	}
	ctx.JSON(iris.Map{})
}
