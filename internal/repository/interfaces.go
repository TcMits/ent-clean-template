package repository

import (
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=repository

type (
	GetModelRepository[ModelType, PredicateModelType any] interface {
		Get(context.Context, ...PredicateModelType) (ModelType, error)
	}

	LoginRepository[UserType, PredicateUserType, LoginInputType any] interface {
		Get(context.Context, ...PredicateUserType) (UserType, error)
		Login(context.Context, LoginInputType) (UserType, error)
	}
)
