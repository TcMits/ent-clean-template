package v1_test

import (
	"context"
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestDocs(t *testing.T) {
	client := getEntClient(t)
	ctx := context.Background()
	isSuperuser := true
	u := createUser(t, ctx, client, map[string]any{"IsSuperuser": &isSuperuser})
	u2 := createUser(t, ctx, client, nil)

	token := ""
	token2 := ""
	Test(t, getAccessTokenISteps(u.Username, "12345678", &token)...)
	Test(t, getAccessTokenISteps(u2.Username, "12345678", &token2)...)

	tests := []struct {
		name string
		args goHitArgs
	}{
		{
			name: "Success",
			args: goHitArgs{
				args: []IStep{
					Description("Redirect Success"),
					Get(_docsPath),
					Send().Headers("Content-Type").Add("application/json"),
					Expect().Status().Equal(http.StatusForbidden),
				},
			},
		},
		{
			name: "IndexSuccess",
			args: goHitArgs{
				args: []IStep{
					Description("Index Success"),
					Get(_docsIndexPath),
					Send().Headers("Content-Type").Add("application/json"),
					getSendAuthenticationHeaderIStep(token),
					Expect().Status().Equal(http.StatusOK),
				},
			},
		},
		{
			name: "PermissionDeniedWithNonSuperuser",
			args: goHitArgs{
				args: []IStep{
					Description("Index Success"),
					Get(_docsIndexPath),
					Send().Headers("Content-Type").Add("application/json"),
					getSendAuthenticationHeaderIStep(token2),
					Expect().Status().Equal(http.StatusForbidden),
				},
			},
		},
		{
			name: "PermissionDeniedWithoutUser",
			args: goHitArgs{
				args: []IStep{
					Description("Index Success"),
					Get(_docsIndexPath),
					Send().Headers("Content-Type").Add("application/json"),
					Expect().Status().Equal(http.StatusForbidden),
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
