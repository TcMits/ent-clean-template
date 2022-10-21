package v1

import (
	"fmt"

	"github.com/TcMits/ent-clean-template/internal/controller/http/v1/middleware"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"

	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

const (
	_publicMeSubPath   = "/me"
	_publicMeRouteName = "publicMe"
)

var _wrapPublicMeReadBodyError = func(err error) error {
	return model.NewTranslatableError(
		fmt.Errorf("internal.controller.http.v1.user.RegisterPublicMeController: %w", err),
		_oneOrMoreFieldsFailedToBeValidatedMessage,
		_useCaseInputValidationError,
		nil,
	)
}

// @Summary Me endpoints
// @Tags    me
// @Accept  mpfd,x-www-form-urlencoded,json
// @Produce json
// @Success 200 {object} usecase.publicMeUseCaseUpdateSerializedInfo
// @Failure 400 {object} errorResponse
// @Failure 403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router  /me [get]
// @Param   payload body     useCaseModel.PublicMeUseCaseUpdateInput true "Payload"
// @Param   payload formData useCaseModel.PublicMeUseCaseUpdateInput true "Payload"
// @Router  /me [put]
// @Router  /me [patch]
func RegisterPublicMeController(
	handler iris.Party,
	getUseCase usecase.GetModelUseCase[*model.User, *struct{}],
	getAndUpdateUseCase usecase.GetAndUpdateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput],
	serializeUseCase usecase.SerializeModelUseCase[*model.User, map[string]any],
	l logger.Interface,
) {
	getHandler := getDetailHandler(
		getUseCase,
		serializeUseCase,
		l,
		func(err error) error { return err },
		func(err error) error { return err },
	)
	updateHandler := getUpdateHandler(
		getAndUpdateUseCase,
		serializeUseCase,
		l,
		func(err error) error { return err },
		func(err error) error { return err },
		_wrapPublicMeReadBodyError,
	)
	isAuthenticatedMiddleware := middleware.Permission(
		handleError,
		l,
		usecase.NewIsAuthenticatedPermissionChecker(),
	)

	handler.Get(_publicMeSubPath, isAuthenticatedMiddleware, getHandler).Name = _publicMeRouteName
	handler.Put(_publicMeSubPath, isAuthenticatedMiddleware, updateHandler).Name = _publicMeRouteName
	handler.Patch(_publicMeSubPath, isAuthenticatedMiddleware, updateHandler).Name = _publicMeRouteName
}
