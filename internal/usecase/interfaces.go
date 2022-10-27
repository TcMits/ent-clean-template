// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=usecase

type (
	SerializeModelUseCase[ModelType, SerializedType any] interface {
		Serialize(context.Context, ModelType) SerializedType
	}
	ListModelUseCase[ModelType, OrderInput, WhereInput any] interface {
		List(context.Context, *int, *int, OrderInput, WhereInput) ([]ModelType, error)
	}
	GetModelUseCase[ModelType, WhereInput any] interface {
		Get(context.Context, WhereInput) (ModelType, error)
	}
	CountModelUseCase[WhereInput any] interface {
		Count(context.Context, WhereInput) (int, error)
	}
	CreateModelUseCase[ModelType, CreateInput any] interface {
		Create(context.Context, CreateInput) (ModelType, error)
	}
	GetAndUpdateModelUseCase[ModelType, WhereInput, UpdateInput any] interface {
		GetAndUpdate(context.Context, WhereInput, UpdateInput) (ModelType, error)
	}
	GetAndDeleteModelUseCase[ModelType, WhereInput any] interface {
		GetAndDelete(context.Context, WhereInput) error
	}
	UserPermissionCheckerUseCase[UserType any] interface {
		Check(context.Context, UserType) error
		Or(UserPermissionCheckerUseCase[UserType]) UserPermissionCheckerUseCase[UserType]
		And(UserPermissionCheckerUseCase[UserType]) UserPermissionCheckerUseCase[UserType]
	}

	LoginUseCase[LoginInputType, JWTAuthenticatedPayloadType, RefreshTokenInputType, UserType any] interface {
		Login(context.Context, LoginInputType) (JWTAuthenticatedPayloadType, error)
		RefreshToken(context.Context, RefreshTokenInputType) (string, error)
		VerifyToken(context.Context, string) (UserType, error)
	}
)
