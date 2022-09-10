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
	logErr := err
	if unwrapedErr := errors.Unwrap(err); unwrapedErr != nil {
		logErr = unwrapedErr
	}
	switch code {
	case usecase.PermissionDeniedError,
		usecase.AuthenticationError,
		usecase.ValidationError,
		usecase.NotFoundError,
		UscaseInputValidationError:
		l.Info(logErr.Error())
	case usecase.DBError:
		l.Warn(logErr.Error())
	case usecase.InternalServerError, UnknownError:
		l.Error(err)
	default:
		l.Info(logErr.Error())
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

// alias handleError
func HandleError(ctx iris.Context, err error, l logger.Interface) {
	handleError(ctx, err, l)
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

func handleBindingError(
	ctx iris.Context,
	err error,
	l logger.Interface,
	input any,
	wrapTranslateError func(model.TranslateFunc, error) error,
) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		err = translatableErrorFromValidationErrors(input, errs, ctx.Tr)
	} else if wrapTranslateError != nil {
		err = wrapTranslateError(ctx.Tr, err)
	}
	handleError(ctx, err, l)
}
