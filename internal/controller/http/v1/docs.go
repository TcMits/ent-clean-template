package v1

import (
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

func RegisterDocsController(handler iris.Party) {
	swaggerHandler := swagger.WrapHandler(
		swaggerFiles.Handler,
		swagger.URL(_swaggerDocPath),
	)
	handler.Get(_swaggerSubPath, func(ctx iris.Context) {
		ctx.Redirect(_swaggerRedirectPath, iris.StatusPermanentRedirect)
	})
	handler.Get(_swaggerSubPathAsset, swaggerHandler)
}
