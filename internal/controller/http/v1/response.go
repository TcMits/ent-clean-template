package v1

type (
	emptyResponse struct{}

	// error
	errorResponse struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}

	// login
	refreshTokenResponse struct {
		Token string `json:"token"`
	}
)
