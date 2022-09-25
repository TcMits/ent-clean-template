package factory

import (
	"context"

	"github.com/bluele/factory-go/factory"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
)

type generator[MutationType ent.Mutation, MutationInputType model.MutationInput[MutationType]] struct {
	client *factory.Factory
}

func (g *generator[_, MutationInputType]) Generate(
	ctx context.Context,
	opt map[string]any,
) MutationInputType {
	return g.client.MustCreateWithContextAndOption(ctx, opt).(MutationInputType)
}
