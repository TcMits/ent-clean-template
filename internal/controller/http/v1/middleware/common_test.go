package middleware

import (
	"reflect"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestCommon(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Success",
			want: iris.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Common()
			handler := iris.New()
			handler.Use(got)
			handler.Get("/test", func(ctx iris.Context) {
				request := ctx.Request()
				requestCtx := request.Context()
				if !reflect.DeepEqual(requestCtx.Value(_currentURLKey), request.URL) {
					ctx.StatusCode(iris.StatusBadRequest)
				}
				ctx.JSON(iris.Map{})
			})

			e := httptest.New(t, handler)
			e.GET("/test").Expect().Status(tt.want)
		})
	}
}
