package v1

import (
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	modelUseCase "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

const (
	UnknownCode = "UNKNOWN"
)

func getCodeFromError(err error) string {
	haveCodeErr, ok := err.(interface{ Code() string })
	if !ok {
		return UnknownCode
	}
	return haveCodeErr.Code()
}

func getStatusCodeFromCode(code string) int {
	switch code {
	case usecase.PermissionDeniedError:
		return iris.StatusForbidden
	case usecase.AuthenticationError:
		return iris.StatusUnauthorized
	case usecase.InternalServerError, UnknownCode:
		return iris.StatusInternalServerError
	case usecase.ValidationError:
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
		usecase.NotFoundError:
		l.Info(err.Error())
	case usecase.DBError:
		l.Warn(err.Error())
	case usecase.InternalServerError, UnknownCode:
		l.Error(err)
	default:
		l.Info(err.Error())
	}
}

func HandleError(ctx iris.Context, err error, l logger.Interface) {
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
