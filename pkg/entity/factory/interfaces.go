package factory

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
)

type Generator[MutationType ent.Mutation, MutationInputType model.MutationInput[MutationType]] interface {
	Generate(context.Context, map[string]any) MutationInputType
}

type ModelFactory[
	ModelType any,
	MutationType ent.Mutation,
	MutationInputType model.MutationInput[MutationType],
	ModelCreatorType model.Creator[ModelType, MutationType],
] interface {
	Build(context.Context, map[string]any) MutationInputType
	Create(context.Context, ModelCreatorType, map[string]any) (ModelType, error)
}
