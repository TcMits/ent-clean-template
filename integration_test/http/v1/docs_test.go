package v1_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestDocs(t *testing.T) {
	tests := []struct {
		name string
		args goHitArgs
	}{
		{
			name: "Success",
			args: goHitArgs{
				args: []IStep{
					Description("Success"),
					Get(_docsPath),
					Send().Headers("Content-Type").Add("application/json"),
					Expect().Status().Equal(http.StatusOK),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Test(t, tt.args.args...)
		})
	}
}
