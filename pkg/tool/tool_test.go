package tool

import (
	"net/url"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/stretchr/testify/require"
)

func TestGetIrisReverseFunc(t *testing.T) {
	type args struct {
		routeName string
		provider  router.RoutesProvider
		a         []any
		v         url.Values
		u         *url.URL
	}

	handler := iris.New()
	handler.Get("/test/{id:string}", func(ctx iris.Context) {}).Name = "testDetail"
	u, err := url.Parse("https://sampleserver.com/hello/world?test2=2")
	require.NoError(t, err)

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "WithoutCurrentURL",
			args: args{
				routeName: "testDetail",
				provider:  handler,
				a:         append(make([]any, 0, 1), 1),
				v:         url.Values{"test": []string{"2"}},
			},
			want: "/test/1?test=2",
		},
		{
			name: "WithoutURL",
			args: args{
				routeName: "testDetail",
				provider:  handler,
				a:         append(make([]any, 0, 1), 1),
				v:         url.Values{"test": []string{"2"}},
				u:         u,
			},
			want: u.Scheme + "://" + u.Host + "/test/1?test=2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIrisReverseFunc(tt.args.routeName, tt.args.provider)
			if gotResult := got(tt.args.u, tt.args.v, tt.args.a...); gotResult != tt.want {
				t.Errorf("GetIrisReverseFunc() = %v, want %v", gotResult, tt.want)
			}
		})
	}
}
