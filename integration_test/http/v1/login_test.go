package v1_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestLogin(t *testing.T) {
	client := getEntClient(t)
	ctx := context.Background()
	u := createUser(t, ctx, client, nil)

	tests := []struct {
		name string
		args goHitArgs
	}{
		{
			name: "Success",
			args: goHitArgs{
				args: []IStep{
					Description("Login Success"),
					Post(_loginPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "username": "%s",
              "password": "%s"
          }`, u.Username, "12345678")),
					Expect().Status().Equal(http.StatusOK),
					Expect().Body().JSON().JQ(".access_token").NotEqual(""),
					Expect().Body().JSON().JQ(".refresh_token").NotEqual(""),
					Expect().Body().JSON().JQ(".refresh_key").NotEqual(""),
				},
			},
		},
		{
			name: "WrongPassword",
			args: goHitArgs{
				args: []IStep{
					Description("Login With Wrong Password"),
					Post(_loginPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "username": "%s",
              "password": "%s"
          }`, u.Username, "1234567")),
					Expect().Status().Equal(http.StatusUnauthorized),
				},
			},
		},
		{
			name: "WrongUsername",
			args: goHitArgs{
				args: []IStep{
					Description("Login With Wrong Username"),
					Post(_loginPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "username": "%s",
              "password": "%s"
          }`, u.Username+"wrong", "12345678")),
					Expect().Status().Equal(http.StatusUnauthorized),
				},
			},
		},
		{
			name: "WithoutUsername",
			args: goHitArgs{
				args: []IStep{
					Description("Login Without Username"),
					Post(_loginPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "password": "%s"
          }`, "12345678")),
					Expect().Status().Equal(http.StatusBadRequest),
				},
			},
		},
		{
			name: "WithoutPassword",
			args: goHitArgs{
				args: []IStep{
					Description("Login Without Password"),
					Post(_loginPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "username": "%s",
          }`, u.Username+"wrong")),
					Expect().Status().Equal(http.StatusBadRequest),
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

func TestRefreshToken(t *testing.T) {
	client := getEntClient(t)
	ctx := context.Background()
	u := createUser(t, ctx, client, nil)

	refreshToken := ""
	refreshKey := ""

	Test(
		t,
		Post(_loginPath),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(fmt.Sprintf(`{
              "username": "%s",
              "password": "%s"
          }`, u.Username, "12345678")),
		Store().Response().Body().JSON().JQ(".refresh_token").In(&refreshToken),
		Store().Response().Body().JSON().JQ(".refresh_key").In(&refreshKey),
	)

	tests := []struct {
		name string
		args goHitArgs
	}{
		{
			name: "Success",
			args: goHitArgs{
				args: []IStep{
					Description("Refresh Token Success"),
					Post(_refreshTokenPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "refresh_token": "%s",
              "refresh_key": "%s"
          }`, refreshToken, refreshKey)),
					Expect().Status().Equal(http.StatusOK),
					Expect().Body().JSON().JQ(".token").NotEqual(""),
				},
			},
		},
		{
			name: "WithoutRefreshToken",
			args: goHitArgs{
				args: []IStep{
					Description("Without Refresh Token"),
					Post(_refreshTokenPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "refresh_key": "%s"
          }`, refreshKey)),
					Expect().Status().Equal(http.StatusBadRequest),
				},
			},
		},
		{
			name: "WithoutRefreshKey",
			args: goHitArgs{
				args: []IStep{
					Description("Without Refresh Key"),
					Post(_refreshTokenPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "refresh_token": "%s"
          }`, refreshToken)),
					Expect().Status().Equal(http.StatusBadRequest),
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

func TestVerifyToken(t *testing.T) {
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
			name: "Success",
			args: goHitArgs{
				args: []IStep{
					Description("Verify Token Success"),
					Post(_verifyTokenPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{
              "token": "%s"
          }`, token)),
					Expect().Status().Equal(http.StatusOK),
				},
			},
		},
		{
			name: "WithoutToken",
			args: goHitArgs{
				args: []IStep{
					Description("Wrong Refresh Token"),
					Post(_refreshTokenPath),
					Send().Headers("Content-Type").Add("application/json"),
					Send().Body().String(fmt.Sprintf(`{}`)),
					Expect().Status().Equal(http.StatusBadRequest),
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
