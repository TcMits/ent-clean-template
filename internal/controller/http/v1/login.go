package v1

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

const (
	loginSubPath        = "/login"
	refreshTokenSubPath = "/refresh-token"
	verifyTokenSubPath  = "/verify-token"

	loginRouteName        = "login"
	refreshTokenRouteName = "refreshToken"
	verifyTokenRouteName  = "verifyToken"
)

var (
	_wrapInvalidLoginInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("v1 - getLoginHandler - ctx.ReadBody: %w", err),
			defaultInvalidErrorTranslateKey,
			translationFunc,
			defaultInvalidErrorMessage,
			UscaseInputValidationError,
		)
	}
	_wrapInvalidRefreshInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("v1 - getRefreshTokenHandler - ctx.ReadBody: %w", err),
			defaultInvalidErrorTranslateKey,
			translationFunc,
			defaultInvalidErrorMessage,
			UscaseInputValidationError,
		)
	}
	_wrapInvalidVerifyTokenInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("v1 - getVerifyTokenHandler - ctx.ReadBody: %w", err),
			defaultInvalidErrorTranslateKey,
			translationFunc,
			defaultInvalidErrorMessage,
			UscaseInputValidationError,
		)
	}
)

type refreshTokenResponse struct {
	Token string `json:"token"`
}

type verifyTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

func RegisterLoginController[
	PLoginInputType interface{ *LoginInputType },
	JWTAuthenticatedPayloadType any,
	PRefreshTokenInputType interface{ *RefreshTokenInputType },
	UserType,
	LoginInputType,
	RefreshTokenInputType any,
](
	handler iris.Party,
	useCase usecase.LoginUseCase[PLoginInputType, JWTAuthenticatedPayloadType, PRefreshTokenInputType, UserType],
	l logger.Interface,
) {
	if handler == nil {
		panic("handler is required")
	}
	if useCase == nil {
		panic("useCase is required")
	}
	if l == nil {
		panic("l is required")
	}
	handler.Post(loginSubPath, getLoginHandler(useCase, l)).Name = loginRouteName
	handler.Post(refreshTokenSubPath, getRefreshTokenHandler(useCase, l)).Name = refreshTokenRouteName
	handler.Post(verifyTokenSubPath, getVerifyTokenHandler(useCase, l)).Name = verifyTokenRouteName
}

func getLoginHandler[
	PLoginInputType interface{ *LoginInputType },
	JWTAuthenticatedPayloadType any,
	PRefreshTokenInputType interface{ *RefreshTokenInputType },
	UserType,
	LoginInputType,
	RefreshTokenInputType any,
](
	useCase usecase.LoginUseCase[PLoginInputType, JWTAuthenticatedPayloadType, PRefreshTokenInputType, UserType],
	l logger.Interface,
) iris.Handler {
	return func(ctx iris.Context) {
		loginInput := PLoginInputType(new(LoginInputType))
		if err := ctx.ReadBody(loginInput); err != nil {
			handleBindingError(ctx, err, l, loginInput, _wrapInvalidLoginInput)
			return
		}
		authenticatedPayload, err := useCase.Login(ctx.Request().Context(), loginInput)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		ctx.JSON(authenticatedPayload)
	}
}

func getRefreshTokenHandler[
	PLoginInputType interface{ *LoginInputType },
	JWTAuthenticatedPayloadType any,
	PRefreshTokenInputType interface{ *RefreshTokenInputType },
	UserType,
	LoginInputType,
	RefreshTokenInputType any,
](
	useCase usecase.LoginUseCase[PLoginInputType, JWTAuthenticatedPayloadType, PRefreshTokenInputType, UserType],
	l logger.Interface,

) iris.Handler {
	return func(ctx iris.Context) {
		refreshTokenInput := PRefreshTokenInputType(new(RefreshTokenInputType))
		if err := ctx.ReadJSON(refreshTokenInput); err != nil {
			handleBindingError(ctx, err, l, refreshTokenInput, _wrapInvalidRefreshInput)
			return
		}
		token, err := useCase.RefreshToken(ctx.Request().Context(), refreshTokenInput)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		ctx.JSON(refreshTokenResponse{Token: token})
	}
}

func getVerifyTokenHandler[
	PLoginInputType interface{ *LoginInputType },
	JWTAuthenticatedPayloadType any,
	PRefreshTokenInputType interface{ *RefreshTokenInputType },
	UserType,
	LoginInputType,
	RefreshTokenInputType any,
](
	useCase usecase.LoginUseCase[PLoginInputType, JWTAuthenticatedPayloadType, PRefreshTokenInputType, UserType],
	l logger.Interface,

) iris.Handler {
	return func(ctx iris.Context) {
		verifyTokenInput := new(verifyTokenRequest)
		if err := ctx.ReadBody(verifyTokenInput); err != nil {
			handleBindingError(ctx, err, l, verifyTokenInput, _wrapInvalidVerifyTokenInput)
			return
		}
		_, err := useCase.VerifyToken(ctx.Request().Context(), verifyTokenInput.Token)
		if err != nil {
			handleError(ctx, err, l)
			return
		}
		ctx.JSON(iris.Map{})
	}
}
