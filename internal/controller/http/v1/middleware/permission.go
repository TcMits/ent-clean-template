package middleware

import (
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool/generic"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
	"github.com/kataras/iris/v12"
)

func Permission[UserType any](l logger.Interface, checkers ...usecase.UserPermissionCheckerUseCase[UserType]) iris.Handler {
	checker := usecase.NewAllowAnyPermissionChecker[UserType]()
	for _, c := range checkers {
		checker = checker.And(c)
	}
	return func(ctx iris.Context) {
		user, ok := ctx.Values().Get(UserKey).(lazy.LazyObject[UserType])
		context := ctx.Request().Context()
		if !ok {
			user = lazy.NewLazyObject(func() UserType { return generic.Zero[UserType]() })
		}
		userValue := user.Value()
		if err := checker.Check(context, userValue); err != nil {
			v1.HandleError(ctx, err, l)
			return
		}
		ctx.Next()
	}
}
