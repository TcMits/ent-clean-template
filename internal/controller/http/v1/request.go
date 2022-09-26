package v1

type verifyTokenRequest struct {
	Token string `json:"token" validate:"required"`
}
