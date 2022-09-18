package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/lazy"
	"github.com/kataras/iris/v12"
)

const (
	JWTPrefix     = "JWT"
	UserKey       = "User"
	AuthHeaderKey = "Authorization"
)

// FromHeader is a token extractor.
// It reads the token from the Authorization request header of form:
// Authorization: "{JWTPrefix} {token}".
func fromHeader(ctx iris.Context) string {
	authHeader := ctx.GetHeader(AuthHeaderKey)
	if authHeader == "" {
		return ""
	}

	// pure check: authorization header format must be Bearer {token}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != JWTPrefix {
		return ""
	}

	return authHeaderParts[1]
}

// FromQuery is a token extractor.
// It reads the token from the "{JWTPrefix}" url query parameter.
func fromQuery(ctx iris.Context) string {
	return ctx.URLParam(JWTPrefix)
}

func Auth[
	LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType any,
](loginUseCase usecase.LoginUseCase[
	LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType,
]) iris.Handler {
	if loginUseCase == nil {
		panic("loginUseCase is required")
	}
	return func(ctx iris.Context) {
		request := ctx.Request()
		requestCtx := request.Context()
		token := fromHeader(ctx)
		if token == "" && ctx.Method() == http.MethodGet {
			token = fromQuery(ctx)
		}
		getUser := func() UserType {
			user, _ := loginUseCase.VerifyToken(requestCtx, token)
			return user
		}

		// set both iris context and request context
		lazyUser := lazy.NewLazyObject(getUser)
		ctx.ResetRequest(request.WithContext(context.WithValue(requestCtx, UserKey, lazyUser)))
		ctx.Values().Set(UserKey, lazyUser)
		ctx.Next()
	}
}
