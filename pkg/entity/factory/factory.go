package factory

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
)

var UserFactory = NewFactory[*model.User, *model.UserMutation, *model.UserCreate](NewUserGenerator[*model.User, *model.UserMutation])

type getGeneratorFunc[ModelType any, MutationType ent.Mutation] func() Generator[ModelType, MutationType]

type modelFactory[ModelType any, MutationType ent.Mutation, ModelCreatorType ModelCreator[ModelType, MutationType]] struct {
	getGeneratorFunc[ModelType, MutationType]

	generator Generator[ModelType, MutationType]
}

func NewFactory[ModelType any, MutationType ent.Mutation, ModelCreatorType ModelCreator[ModelType, MutationType]](
	genFunc getGeneratorFunc[ModelType, MutationType],
) ModelFactory[ModelType, MutationType, ModelCreatorType] {
	return &modelFactory[ModelType, MutationType, ModelCreatorType]{getGeneratorFunc: genFunc}
}

func (f *modelFactory[ModelType, MutationType, _]) Build(
	ctx context.Context, mutation MutationType, opt map[string]any) ModelType {
	f.ensureGenerator(mutation)
	return f.generator.Generate(ctx, opt)
}

func (f *modelFactory[ModelType, _, ModelCreatorType]) Create(
	ctx context.Context, creator ModelCreatorType, opt map[string]any) (ModelType, error) {
	f.ensureGenerator(creator.Mutation())
	f.generator.Generate(ctx, opt)
	return creator.Save(ctx)
}

func (f *modelFactory[_, MutationType, _]) ensureGenerator(mutation MutationType) {
	if f.generator == nil {
		f.generator = f.getGeneratorFunc()
	}
	f.generator.SetMutation(mutation)
}
