package factory

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
)

type Generator[ModelType any, MutationType ent.Mutation] interface {
	Generate(context.Context, map[string]any) ModelType
	SetMutation(MutationType)
}

type ModelCreator[ModelType any, MutationType ent.Mutation] interface {
	Mutation() MutationType
	Save(context.Context) (ModelType, error)
}

type ModelFactory[ModelType any, MutationType ent.Mutation, ModelCreatorType ModelCreator[ModelType, MutationType]] interface {
	Build(context.Context, MutationType, map[string]any) ModelType
	Create(context.Context, ModelCreatorType, map[string]any) (ModelType, error)
}
