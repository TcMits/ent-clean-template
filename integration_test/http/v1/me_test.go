package v1_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestMe(t *testing.T) {
	client := getEntClient(t)
	ctx := context.Background()
	u := createUser(t, ctx, client, nil)

	token := ""
	Test(t, getAccessTokenISteps(u.Username, "12345678", &token)...)

	tests := []struct {
		name string
		args goHitArgs
	}{
		{
			name: "GetMeSuccess",
			args: goHitArgs{
				args: []IStep{
					Description("Get Me"),
					Get(_mePath),
					Send().Headers("Content-Type").Add("application/json"),
					getSendAuthenticationHeaderIStep(token),
					Expect().Status().Equal(http.StatusOK),
					Expect().Body().JSON().JQ(".id").Equal(u.ID.String()),
					Expect().Body().JSON().JQ(".username").Equal(u.Username),
					Expect().Body().JSON().JQ(".first_name").Equal(u.FirstName),
					Expect().Body().JSON().JQ(".last_name").Equal(u.LastName),
					Expect().Body().JSON().JQ(".email").Equal(u.Email),
					Expect().Body().JSON().JQ(".is_staff").Equal(u.IsStaff),
					Expect().Body().JSON().JQ(".is_superuser").Equal(u.IsSuperuser),
					Expect().Body().JSON().JQ(".is_active").Equal(u.IsActive),
					Expect().Body().JSON().JQ(".self").Equal(_meSubPath + "?"),
				},
			},
		},
		{
			name: "PutMeSuccess",
			args: goHitArgs{
				args: []IStep{
					Description("Put Me"),
					Put(_mePath),
					Send().Headers("Content-Type").Add("application/json"),
					getSendAuthenticationHeaderIStep(token),
					Send().Body().String(fmt.Sprintf(`{
                "username": "%s",
                "first_name": "%s",
                "last_name": "%s",
                "email": "%s"
              }`,
						"newusername",
						"new first name",
						"new last name",
						"newemail@gmail.com",
					)),

					Expect().Status().Equal(http.StatusOK),
					Expect().Body().JSON().JQ(".id").Equal(u.ID.String()),
					Expect().Body().JSON().JQ(".username").Equal("newusername"),
					Expect().Body().JSON().JQ(".first_name").Equal("new first name"),
					Expect().Body().JSON().JQ(".last_name").Equal("new last name"),
					Expect().Body().JSON().JQ(".email").Equal("newemail@gmail.com"),
					Expect().Body().JSON().JQ(".is_staff").Equal(u.IsStaff),
					Expect().Body().JSON().JQ(".is_superuser").Equal(u.IsSuperuser),
					Expect().Body().JSON().JQ(".is_active").Equal(u.IsActive),
					Expect().Body().JSON().JQ(".self").Equal(_meSubPath + "?"),
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
