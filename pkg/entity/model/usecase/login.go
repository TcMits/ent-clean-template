package usecase

type LoginInput struct {
	Username string `json:"username" validate:"alphanum,required,min=1,max=128"`
	Password string `json:"password" validate:"required,min=1"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
	RefreshKey   string `json:"refresh_key" validate:"required"`
}

type JWTAuthenticatedPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	RefreshKey   string `json:"refresh_key"`
}
