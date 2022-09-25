package v1_test

import (
	"fmt"
	"testing"

	. "github.com/Eun/go-hit"
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/datastore"
	"github.com/stretchr/testify/require"
)

const (
	// Attempts connection
	_host       = "golang:8080"
	_healthPath = "http://" + _host + "/ping"
	_attempts   = 20

	// HTTP REST
	_v1SubPath           = "/v1"
	_loginSubPath        = _v1SubPath + "/login"
	_refreshTokenSubPath = _v1SubPath + "/refresh-token"
	_verifyTokenSubPath  = _v1SubPath + "/verify-token"
	_meSubPath           = _v1SubPath + "/me"

	_basePath         = "http://" + _host
	_loginPath        = _basePath + _loginSubPath
	_refreshTokenPath = _basePath + _refreshTokenSubPath
	_verifyTokenPath  = _basePath + _verifyTokenSubPath
	_mePath           = _basePath + _meSubPath

	// request num
	_requests = 10
)

type goHitArgs struct {
	args []IStep
}

func getConf(t *testing.T) *config.Config {
	t.Helper()
	conf, err := config.NewConfig()
	require.NoError(t, err)
	return conf
}

func getEntClient(t *testing.T) *ent.Client {
	t.Helper()
	conf := getConf(t)
	client, err := datastore.NewClient(conf.PG.URL, 1)
	require.NoError(t, err)
	return client
}

func getAccessTokenISteps(username string, password string, token *string) []IStep {
	return append(
		[]IStep{},
		Post(_loginPath),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(fmt.Sprintf(`{
        "username": "%s",
        "password": "%s"
    }`, username, password)),
		Store().Response().Body().JSON().JQ(".access_token").In(token),
	)
}

func getSendAuthenticationHeaderIStep(token string) IStep {
	return Send().Headers("Authorization").Add("JWT " + token)
}