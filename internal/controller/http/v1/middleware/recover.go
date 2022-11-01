package middleware

import (
	"fmt"
	"net/http/httputil"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func getRequestLogs(ctx *context.Context) string {
	rawReq, _ := httputil.DumpRequest(ctx.Request(), false)
	return string(rawReq)
}

// New returns a new recover middleware,
// it recovers from panics and logs
// the panic message to the application's logger "Warn" level.
func Recover(
	handleErrorFunc func(iris.Context, error, logger.Interface),
	l logger.Interface,
) context.Handler {
	if l == nil {
		panic("l is required")
	}
	if handleErrorFunc == nil {
		panic("handleErrorFunc is required")
	}

	return func(ctx *context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() { // handled by other middleware.
					return
				}

				var callers []string
				for i := 1; ; i++ {
					_, file, line, got := runtime.Caller(i)
					if !got {
						break
					}

					callers = append(callers, fmt.Sprintf("%s:%d", file, line))
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprint(getRequestLogs(ctx))
				logMessage += fmt.Sprintf("%s\n", err)
				logMessage += fmt.Sprintf("%s\n", strings.Join(callers, "\n"))
				l.Warn(logMessage)

				// get the list of registered handlers and the
				// handler which panic derived from.
				handlers := ctx.Handlers()
				handlersFileLines := make([]string, 0, len(handlers))
				currentHandlerIndex := ctx.HandlerIndex(-1)
				currentHandlerFileLine := "???"
				for i, h := range ctx.Handlers() {
					file, line := context.HandlerFileLine(h)
					fileline := fmt.Sprintf("%s:%d", file, line)
					handlersFileLines = append(handlersFileLines, fileline)
					if i == currentHandlerIndex {
						currentHandlerFileLine = fileline
					}
				}

				panicErr := &context.ErrPanicRecovery{
					Cause:              err,
					Callers:            callers,
					Stack:              debug.Stack(),
					RegisteredHandlers: handlersFileLines,
					CurrentHandler:     currentHandlerFileLine,
				}

				if handleErrorFunc != nil {
					handleErrorFunc(ctx, panicErr, l)
				} else {
					// see accesslog.wasRecovered too.
					ctx.StopWithPlainError(500, panicErr)
				}
			}
		}()

		ctx.Next()
	}
}
