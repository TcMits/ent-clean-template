// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=usecase

type (
	LoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType any] interface {
		Login(context.Context, LoginInputType) (JWTAuthenticatedPayloadType, error)
		RefreshToken(context.Context, RefreshTokenInputType) (string, error)
		VerifyToken(context.Context, string) (UserType, error)
	}
)
