package repository

import (
	"context"
	"fmt"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/ent/user"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/password"
)

type loginRepository struct {
	client *ent.Client
}

func NewLoginRepository(client *ent.Client) LoginRepository {
	return &loginRepository{client: client}
}

func (repo *loginRepository) Get(
	ctx context.Context, predicateUsers ...model.PredicateUser) (*model.User, error) {
	u, err := repo.client.User.Query().Where(predicateUsers...).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"loginRepository - Login - repo.client.User.Query.Where.Only: %w", err,
		)
	}
	return u, nil
}

func (repo *loginRepository) Login(
	ctx context.Context, loginInput *useCaseModel.LoginInput) (*model.User, error) {
	user, err := repo.Get(
		ctx, user.UsernameEQ(loginInput.Username), user.IsActiveEQ(true),
	)
	if err != nil {
		return nil, fmt.Errorf("loginRepository - Login - repo.Get: %w", err)
	}
	if !password.ValidatePassword(user.Password, loginInput.Password) {
		return nil, fmt.Errorf("loginRepository - Login: %w", err)
	}
	return user, nil
}
