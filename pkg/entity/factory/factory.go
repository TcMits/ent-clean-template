package factory

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
)

var UserFactory = modelFactory[
	*model.User,
	*model.UserMutation,
	*model.UserCreateInput,
	*model.UserCreate,
]{generator: newUserGenerator[*model.UserMutation, *model.UserCreateInput]()}

type modelFactory[
	ModelType any,
	MutationType ent.Mutation,
	MutationInputType model.MutationInput[MutationType],
	ModelCreatorType model.Creator[ModelType, MutationType],
] struct {
	generator Generator[MutationType, MutationInputType]
}

func (f *modelFactory[_, _, MutationInputType, _]) Build(
	ctx context.Context, opt map[string]any) MutationInputType {
	return f.generator.Generate(ctx, opt)
}

func (f *modelFactory[ModelType, _, _, ModelCreatorType]) Create(
	ctx context.Context, creator ModelCreatorType, opt map[string]any) (ModelType, error) {
	mutaitonInput := f.generator.Generate(ctx, opt)
	mutaitonInput.Mutate(creator.Mutation())
	return creator.Save(ctx)
}
