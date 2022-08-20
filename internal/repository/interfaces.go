package repository

import (
	"context"

	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=repository_test

type (
	GetUserRepository interface {
		// ctx, where
		Get(context.Context, ...model.PredicateUser) (*model.User, error)
	}

	LoginRepository interface {
		GetUserRepository

		Login(context.Context, *useCaseModel.LoginInput) (*model.User, error)
	}
)
