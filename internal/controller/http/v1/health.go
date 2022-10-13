package v1

import (
	"github.com/kataras/iris/v12"
)

const (
	_healthCheckSubPath   = "/ping"
	_healthCheckRouteName = "healthCheck"
)

func RegisterHealthCheckController(handler iris.Party) {
	handler.Get(_healthCheckSubPath, func(ctx iris.Context) {
		ctx.WriteString("pong")
	}).Describe("healthcheck").Name = _healthCheckRouteName
}
