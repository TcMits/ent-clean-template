package v1

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

const (
	_unknownError                = "UNKNOWN"
	_usecaseInputValidationError = "USECASE_INPUT_VALIDATION_ERROR"
)

func getCodeFromError(err error) string {
	haveCodeErr, ok := err.(interface{ Code() string })
	if !ok {
		return _unknownError
	}
	return haveCodeErr.Code()
}

func getStatusCodeFromCode(code string) int {
	switch code {
	case usecase.PermissionDeniedError:
		return iris.StatusForbidden
	case usecase.AuthenticationError:
		return iris.StatusUnauthorized
	case usecase.InternalServerError, _unknownError:
		return iris.StatusInternalServerError
	case usecase.ValidationError, _usecaseInputValidationError:
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
		_usecaseInputValidationError:
		l.Info(logErr.Error())
	case usecase.DBError:
		l.Warn(logErr.Error())
	case usecase.InternalServerError, _unknownError:
		l.Error(err)
	default:
		l.Info(logErr.Error())
	}
}

func getTranslateFunc(tr func(string, ...any) string) model.TranslateFunc {
	return func(m *i18n.Message, a ...any) string {
		return tr(m.ID, a...)
	}
}

func handleError(ctx iris.Context, err error, l logger.Interface) {
	code := getCodeFromError(err)
	statusCode := getStatusCodeFromCode(code)
	logError(err, code, l)

	switch foundedError := err.(type) {
	case model.TranslatableError:
		translatableErr := foundedError.SetTranslateFunc(getTranslateFunc(ctx.Tr))
		ctx.StopWithJSON(statusCode, errorResponse{
			Code:    code,
			Message: translatableErr.Error(),
		})
	case *model.TranslatableError:
		translatableErr := foundedError.SetTranslateFunc(getTranslateFunc(ctx.Tr))
		ctx.StopWithJSON(statusCode, errorResponse{
			Code:    code,
			Message: translatableErr.Error(),
		})
	default:
		ctx.StopWithJSON(statusCode, errorResponse{
			Code:    code,
			Message: foundedError.Error(),
		})
	}
}

// alias handleError.
func HandleError(ctx iris.Context, err error, l logger.Interface) {
	handleError(ctx, err, l)
}

func translatableErrorFromValidationErrors(
	inputStructure any, errs validator.ValidationErrors,
) *model.TranslatableError {
	verboser, ok := inputStructure.(interface {
		GetErrorMessageFromStructField(string) *i18n.Message
	})
	var err error = errs
	i18nMessage := _oneOrMoreFieldsFailedToBeValidatedMessage
	if ok {
		for _, validationErr := range errs {
			err = validationErr
			i18nMessage = verboser.GetErrorMessageFromStructField(
				validationErr.StructField(),
			)
			break
		}
	}
	return model.NewTranslatableError(err, i18nMessage, _usecaseInputValidationError, nil)
}

func handleBindingError(
	ctx iris.Context,
	err error,
	l logger.Interface,
	input any,
	wrapTranslateError func(error) error,
) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		err = translatableErrorFromValidationErrors(input, errs)
	} else if wrapTranslateError != nil {
		err = wrapTranslateError(err)
	}
	handleError(ctx, err, l)
}
