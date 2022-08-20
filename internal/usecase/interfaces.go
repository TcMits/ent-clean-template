// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	LoginUseCase interface {
		Login(context.Context, *useCaseModel.LoginInput) (*useCaseModel.JWTAuthenticatedPayload, error)
		RefreshToken(context.Context, *useCaseModel.RefreshTokenInput) (string, error)
		VerifyToken(context.Context, string) (*model.User, error)
	}
)
