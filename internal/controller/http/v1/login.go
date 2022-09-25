package v1

import (
	"fmt"

	"github.com/kataras/iris/v12"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

const (
	_loginSubPath        = "/login"
	_refreshTokenSubPath = "/refresh-token"
	_verifyTokenSubPath  = "/verify-token"

	_loginRouteName        = "login"
	_refreshTokenRouteName = "refreshToken"
	_verifyTokenRouteName  = "verifyToken"
)

var (
	_wrapInvalidLoginInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("v1 - getLoginHandler - ctx.ReadBody: %w", err),
			_defaultInvalidErrorTranslateKey,
			translationFunc,
			_defaultInvalidErrorMessage,
			_uscaseInputValidationError,
		)
	}
	_wrapInvalidRefreshInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("v1 - getRefreshTokenHandler - ctx.ReadBody: %w", err),
			_defaultInvalidErrorTranslateKey,
			translationFunc,
			_defaultInvalidErrorMessage,
			_uscaseInputValidationError,
		)
	}
	_wrapInvalidVerifyTokenInput = func(translationFunc model.TranslateFunc, err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("v1 - getVerifyTokenHandler - ctx.ReadBody: %w", err),
			_defaultInvalidErrorTranslateKey,
			translationFunc,
			_defaultInvalidErrorMessage,
			_uscaseInputValidationError,
		)
	}
)

type refreshTokenResponse struct {
	Token string `json:"token"`
}

type verifyTokenRequest struct {
	Token string `json:"token" validate:"required"`
}

func RegisterLoginController(
	handler iris.Party,
	useCase usecase.LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User],
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
	handler.Post(_loginSubPath, getLoginHandler(useCase, l)).Name = _loginRouteName
	handler.Post(_refreshTokenSubPath, getRefreshTokenHandler(useCase, l)).Name = _refreshTokenRouteName
	handler.Post(_verifyTokenSubPath, getVerifyTokenHandler(useCase, l)).Name = _verifyTokenRouteName
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
