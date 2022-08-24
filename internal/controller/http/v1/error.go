package v1

import (
	"errors"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	modelUseCase "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

const (
	UnknownError               = "UNKNOWN"
	UscaseInputValidationError = "USECASE_INPUT_VALIDATION_ERROR"
)

const (
	defaultInvalidErrorTranslateKey = "internal.controller.http.v1.error.InvalidError"
	defaultInvalidErrorMessage      = "One or more fields failed to be validated"
)

func getCodeFromError(err error) string {
	haveCodeErr, ok := err.(interface{ Code() string })
	if !ok {
		return UnknownError
	}
	return haveCodeErr.Code()
}

func getStatusCodeFromCode(code string) int {
	switch code {
	case usecase.PermissionDeniedError:
		return iris.StatusForbidden
	case usecase.AuthenticationError:
		return iris.StatusUnauthorized
	case usecase.InternalServerError, UnknownError:
		return iris.StatusInternalServerError
	case usecase.ValidationError, UscaseInputValidationError:
		return iris.StatusBadRequest
	case usecase.NotFoundError:
		return iris.StatusNotFound
	case usecase.DBError:
		return iris.StatusNotAcceptable
	default:
		return iris.StatusInternalServerError
	}
}

func logError(err error, code string, l logger.Interface) {
	switch code {
	case usecase.PermissionDeniedError,
		usecase.AuthenticationError,
		usecase.ValidationError,
		usecase.NotFoundError,
		UscaseInputValidationError:
		l.Info(errors.Unwrap(err).Error())
	case usecase.DBError:
		l.Warn(errors.Unwrap(err).Error())
	case usecase.InternalServerError, UnknownError:
		l.Error(err)
	default:
		l.Info(errors.Unwrap(err).Error())
	}
}

func handleError(ctx iris.Context, err error, l logger.Interface) {
	code := getCodeFromError(err)
	statusCode := getStatusCodeFromCode(code)
	logError(err, code, l)

	switch foundedError := err.(type) {
	case *modelUseCase.UseCaseError:
		translatableErr := model.TranslatableErrorFromUseCaseError(foundedError, ctx.Tr)
		ctx.StopWithJSON(statusCode, iris.Map{
			"code":    code,
			"message": translatableErr.Error(),
		})
	default:
		ctx.StopWithJSON(statusCode, iris.Map{
			"code":    code,
			"message": foundedError.Error(),
		})
	}
}

func translatableErrorFromValidationErrors(
	inputStructure any, errs validator.ValidationErrors, tr model.TranslateFunc,
) *model.TranslatableError {
	verboser, ok := inputStructure.(interface {
		GetErrorMessageFromStructField(string) (string, string)
	})
	var err error = errs
	translateKey := defaultInvalidErrorTranslateKey
	defaultErrorMessage := defaultInvalidErrorMessage
	if ok {
		for _, validationErr := range errs {
			err = validationErr
			translateKey, defaultErrorMessage = verboser.GetErrorMessageFromStructField(
				validationErr.StructField(),
			)
			break
		}
	}
	return model.NewTranslatableError(
		err, translateKey, tr, defaultErrorMessage, UscaseInputValidationError,
	)
}
