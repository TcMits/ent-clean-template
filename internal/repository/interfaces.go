package repository

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=repository

type (
	TransactionRepository interface {
		Start(context.Context) (*ent.Tx, error)
		Commit(*ent.Tx) error
		Rollback(*ent.Tx) error
	}
	GetModelRepository[ModelType, WhereInputType any] interface {
		Get(context.Context, WhereInputType) (ModelType, error)
	}
	CountModelRepository[WhereInputType any] interface {
		Count(context.Context, WhereInputType) (int, error)
	}
	GetWithClientModelRepository[ModelType, WhereInputType any] interface {
		GetWithClient(context.Context, *ent.Client, WhereInputType, bool) (ModelType, error)
	}
	ListModelRepository[ModelType, OrderInputType, WhereInputType any] interface {
		List(context.Context, int, int, OrderInputType, WhereInputType) ([]ModelType, error)
	}
	ListWithClientModelRepository[ModelType, OrderInputType, WhereInputType any] interface {
		ListWithClient(context.Context, *ent.Client, int, int, OrderInputType, WhereInputType, bool) ([]ModelType, error)
	}
	CreateModelRepository[ModelType, CreateInputType any] interface {
		Create(context.Context, CreateInputType) (ModelType, error)
	}
	CreateWithClientModelRepository[ModelType, CreateInputType any] interface {
		CreateWithClient(context.Context, *ent.Client, CreateInputType) (ModelType, error)
	}
	UpdateModelRepository[ModelType, UpdateInputType any] interface {
		Update(context.Context, ModelType, UpdateInputType) (ModelType, error)
	}
	UpdateWithClientModelRepository[ModelType, UpdateInputType any] interface {
		UpdateWithClient(context.Context, *ent.Client, ModelType, UpdateInputType) (ModelType, error)
	}
	DeleteModelRepository[ModelType any] interface {
		Delete(context.Context, ModelType) error
	}
	DeleteWithClientModelRepository[ModelType any] interface {
		DeleteWithClient(context.Context, *ent.Client, ModelType) error
	}

	ModelRepository[ModelType, OrderInputType, WhereInputType, CreateInputType, UpdateInputType any] interface {
		GetModelRepository[ModelType, WhereInputType]
		GetWithClientModelRepository[ModelType, WhereInputType]
		CountModelRepository[WhereInputType]
		ListModelRepository[ModelType, OrderInputType, WhereInputType]
		ListWithClientModelRepository[ModelType, OrderInputType, WhereInputType]
		CreateModelRepository[ModelType, CreateInputType]
		CreateWithClientModelRepository[ModelType, CreateInputType]
		UpdateModelRepository[ModelType, UpdateInputType]
		UpdateWithClientModelRepository[ModelType, UpdateInputType]
		DeleteModelRepository[ModelType]
		DeleteWithClientModelRepository[ModelType]
	}

	LoginRepository[UserType, WhereInputType, LoginInputType any] interface {
		GetModelRepository[UserType, WhereInputType]

		Login(context.Context, LoginInputType) (UserType, error)
	}
)
