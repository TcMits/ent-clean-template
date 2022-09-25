package v1

import (
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func TestRegisterHealthCheckController(t *testing.T) {
	type args struct {
		handler *iris.Application
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				handler: iris.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterHealthCheckController(tt.args.handler)
			e := httptest.New(t, tt.args.handler)

			e.GET(_healthCheckSubPath).Expect().Status(iris.StatusOK).Body().Equal("pong")
		})
	}
}
