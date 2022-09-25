package middleware

import (
	"context"

	"github.com/kataras/iris/v12"
)

const _currentURLKey = "URL"

func Common() iris.Handler {
	return func(ctx iris.Context) {
		request := ctx.Request()
		requestCtx := request.Context()
		u := request.URL

		ctx.ResetRequest(request.WithContext(context.WithValue(requestCtx, _currentURLKey, u)))
		ctx.Next()
	}
}
