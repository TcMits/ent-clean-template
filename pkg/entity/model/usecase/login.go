package usecase

type LoginInput struct {
	Username string `json:"username" form:"username" validate:"alphanum,required,min=1,max=128"`
	Password string `json:"password" form:"password" validate:"required,min=1"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
	RefreshKey   string `json:"refresh_key"   form:"refresh_key"   validate:"required"`
}

type JWTAuthenticatedPayload struct {
	AccessToken  string `json:"access_token"  form:"access_token"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
	RefreshKey   string `json:"refresh_key"   form:"refresh_key"`
}
