package repository

import (
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
)

type urr = *ent.UserReadRepository
type ucr = *ent.UserCreateRepository
type uur = *ent.UserUpdateRepository
type udr = *ent.UserDeleteRepository

type userRepository struct {
	urr
	ucr
	uur
	udr
}

func NewUserRepository(client *ent.Client) ModelRepository[
	*model.User, *model.UserOrderInput, *model.UserWhereInput, *model.UserCreateInput, *model.UserUpdateInput,
] {
	if client == nil {
		panic("client is required")
	}
	return &userRepository{
		ent.NewUserReadRepository(client, nil, nil, nil),
		ent.NewUserCreateRepository(client, nil, nil, false),
		ent.NewUserUpdateRepository(client, nil, nil, false),
		ent.NewUserDeleteRepository(client, nil, nil, false),
	}
}
