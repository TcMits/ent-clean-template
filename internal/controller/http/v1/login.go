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
	_wrapInvalidLoginInput = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.controller.http.v1.login.getLoginHandler: %w", err),
			_oneOrMoreFieldsFailedToBeValidatedMessage,
			_useCaseInputValidationError,
			nil,
		)
	}
	_wrapInvalidRefreshInput = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.controller.http.v1.login.getRefreshTokenHandler: %w", err),
			_oneOrMoreFieldsFailedToBeValidatedMessage,
			_useCaseInputValidationError,
			nil,
		)
	}
	_wrapInvalidVerifyTokenInput = func(err error) error {
		return model.NewTranslatableError(
			fmt.Errorf("internal.controller.http.v1.login.getVerifyTokenHandler: %w", err),
			_oneOrMoreFieldsFailedToBeValidatedMessage,
			_useCaseInputValidationError,
			nil,
		)
	}
)

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
	registerLogin(handler, useCase, l)
	registerRefreshToken(handler, useCase, l)
	registerVerifyToken(handler, useCase, l)
}

// @Summary Login endpoint
// @Tags    login
// @Accept  mpfd,x-www-form-urlencoded,json
// @Produce json
// @Param   payload body     useCaseModel.LoginInput true "Payload"
// @Param   payload formData useCaseModel.LoginInput true "Payload"
// @Success 200     {object} useCaseModel.JWTAuthenticatedPayload
// @Failure 400     {object} errorResponse
// @Failure 401     {object} errorResponse
// @Failure 500     {object} errorResponse
// @Router  /login [post]
func registerLogin(
	handler iris.Party,
	useCase usecase.LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User],
	l logger.Interface,
) {
	handler.Post(_loginSubPath, getLoginHandler(useCase, l)).Name = _loginRouteName
}

// @Summary Refresh token endpoint
// @Tags    refresh-token
// @Accept  mpfd,x-www-form-urlencoded,json
// @Produce json
// @Param   payload body     useCaseModel.RefreshTokenInput true "Payload"
// @Param   payload formData useCaseModel.RefreshTokenInput true "Payload"
// @Success 200     {object} refreshTokenResponse
// @Failure 400     {object} errorResponse
// @Failure 401     {object} errorResponse
// @Failure 500     {object} errorResponse
// @Router  /refresh-token [post]
func registerRefreshToken(
	handler iris.Party,
	useCase usecase.LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User],
	l logger.Interface,
) {
	handler.Post(_refreshTokenSubPath, getRefreshTokenHandler(useCase, l)).Name = _refreshTokenRouteName
}

// @Summary Verify token endpoint
// @Tags    verify-token
// @Accept  mpfd,x-www-form-urlencoded,json
// @Produce json
// @Param   payload body     verifyTokenRequest true "Payload"
// @Param   payload formData verifyTokenRequest true "Payload"
// @Success 200     {object} emptyResponse
// @Failure 400     {object} errorResponse
// @Failure 401     {object} errorResponse
// @Failure 500     {object} errorResponse
// @Router  /verify-token [post]
func registerVerifyToken(
	handler iris.Party,
	useCase usecase.LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User],
	l logger.Interface,
) {
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
		if err := ctx.ReadBody(refreshTokenInput); err != nil {
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
