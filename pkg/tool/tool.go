package tool

import (
	"net/url"

	"github.com/kataras/iris/v12/core/router"
)

func GetIrisReverseFunc(
	routeName string,
	provider router.RoutesProvider,
) func(*url.URL, url.Values, ...any) string {
	return func(u *url.URL, v url.Values, a ...any) string {
		result := ""
		options := make([]router.RoutePathReverserOption, 0, 2)
		if u != nil && u.IsAbs() {
			options = append(
				options, router.WithHost(u.Host), router.WithScheme(u.Scheme),
			)
		}
		reverser := router.NewRoutePathReverser(provider, options...)
		if len(options) > 0 {
			result = reverser.URL(routeName, a...)
		} else {
			result = reverser.Path(routeName, a...)
		}

		return result + "?" + v.Encode()
	}
}
