package model

import (
	"context"

	"github.com/TcMits/ent-clean-template/ent"
)

type (
	MutationInput[MutationType ent.Mutation] interface {
		Mutate(MutationType)
	}

	Creator[ModelType any, MutationType ent.Mutation] interface {
		Mutation() MutationType
		Save(context.Context) (ModelType, error)
	}
)
