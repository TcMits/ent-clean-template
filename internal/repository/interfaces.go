package repository

import (
	"context"
	"io"

	"github.com/TcMits/ent-clean-template/ent"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=repository

type (
	// database.
	TransactionRepository interface {
		// return client, commit function, rollback function, error
		Start(context.Context) (*ent.Client, func() error, func() error, error)
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
		List(context.Context, *int, *int, OrderInputType, WhereInputType) ([]ModelType, error)
	}
	ListWithClientModelRepository[ModelType, OrderInputType, WhereInputType any] interface {
		ListWithClient(
			context.Context,
			*ent.Client,
			*int,
			*int,
			OrderInputType,
			WhereInputType,
			bool,
		) ([]ModelType, error)
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
		UpdateWithClient(
			context.Context,
			*ent.Client,
			ModelType,
			UpdateInputType,
		) (ModelType, error)
	}
	DeleteModelRepository[ModelType any] interface {
		Delete(context.Context, ModelType) error
	}
	DeleteWithClientModelRepository[ModelType any] interface {
		DeleteWithClient(context.Context, *ent.Client, ModelType) error
	}

	// files.
	ReadFileRepository interface {
		Read(context.Context, string, io.Writer, int64, int64) (int64, error)
	}
	ExistFileRepository interface {
		Exist(context.Context, string) (bool, error)
	}
	WriteFileRepository interface {
		Write(context.Context, string, io.Reader, int64) (int64, error)
	}
	DeleteFileRepository interface {
		Delete(context.Context, string) error
	}

	FileRepository interface {
		ReadFileRepository
		ExistFileRepository
		WriteFileRepository
		DeleteFileRepository
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
)
