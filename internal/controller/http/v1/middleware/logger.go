package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

func Logger(l logger.Interface) iris.Handler {
	return func(ctx iris.Context) {
		method := ctx.Method()
		ip := ctx.RemoteAddr()
		path := ctx.Request().URL.RequestURI()
		startTime := time.Now()
		ctx.Next()
		latency := time.Now().Sub(startTime)
		status := strconv.Itoa(ctx.GetStatusCode())
		l.Info(fmt.Sprintf("%v %4v %s %s %s", status, latency, ip, method, path))
	}
}