package v1

import (
	"github.com/TcMits/ent-clean-template/internal/controller/http/v1/middleware"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

const (
	_swaggerDocPath      = "/api/v1/swagger/doc.json"
	_swaggerRedirectPath = "/api/v1/swagger/index.html"
	_swaggerSubPath      = "/swagger"
	_swaggerSubPathAsset = "/swagger/{any:path}"
)

func RegisterDocsController(handler iris.Party, l logger.Interface) {
	swaggerHandler := swagger.WrapHandler(
		swaggerFiles.Handler,
		swagger.URL(_swaggerDocPath),
	)
	isSuperuser := middleware.Permission(
		handleError,
		l,
		usecase.NewIsSuperuserPermissionChecker(),
	)

	handler.Get(_swaggerSubPath, func(ctx iris.Context) {
		ctx.Redirect(_swaggerRedirectPath, iris.StatusPermanentRedirect)
	})
	handler.Get(_swaggerSubPathAsset, isSuperuser, swaggerHandler)
}
