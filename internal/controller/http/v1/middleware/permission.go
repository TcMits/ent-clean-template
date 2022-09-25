package middleware

import (
	"github.com/kataras/iris/v12"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
)

func Permission[UserType any](
	handleErrorFunc func(iris.Context, error, logger.Interface),
	l logger.Interface,
	checkers ...usecase.UserPermissionCheckerUseCase[UserType],
) iris.Handler {
	if handleErrorFunc == nil {
		panic("handleErrorFunc is required")
	}
	if l == nil {
		panic("l is required")
	}

	checker := usecase.NewAllowAnyPermissionChecker[UserType]()
	for _, c := range checkers {
		checker = checker.And(c)
	}
	return func(ctx iris.Context) {
		user, ok := ctx.Values().Get(_userKey).(lazy.LazyObject[UserType])
		context := ctx.Request().Context()
		if !ok {
			user = lazy.NewLazyObject(func() UserType { return generic.Zero[UserType]() })
		}
		userValue := user.Value()
		if err := checker.Check(context, userValue); err != nil {
			handleErrorFunc(ctx, err, l)
			return
		}
		ctx.Next()
	}
}
