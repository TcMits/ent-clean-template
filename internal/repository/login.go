package repository

import (
	"context"
	"fmt"

	"github.com/TcMits/ent-clean-template/copygen"
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/tool/password"
)

type loginRepository struct {
	client *ent.Client
}

func NewLoginRepository(client *ent.Client) LoginRepository[
	*model.User, *model.UserWhereInput, *useCaseModel.LoginInput,
] {
	return &loginRepository{client: client}
}

func (repo *loginRepository) Get(
	ctx context.Context, userWhereInput *model.UserWhereInput) (*model.User, error) {

	query, err := userWhereInput.Filter(repo.client.User.Query())
	if err != nil {
		return nil, fmt.Errorf(
			"loginRepository - Login - userWhereInput.Filter: %w", err,
		)
	}
	u, err := query.Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("loginRepository - Login - query.Only: %w", err)
	}
	return u, nil
}

func (repo *loginRepository) Login(
	ctx context.Context, loginInput *useCaseModel.LoginInput) (*model.User, error) {
	isActive := true
	userWhereInput := &model.UserWhereInput{IsActive: &isActive}
	copygen.LoginInputToUserWhereInput(userWhereInput, loginInput)
	user, err := repo.Get(ctx, userWhereInput)
	if err != nil {
		return nil, fmt.Errorf("loginRepository - Login - repo.Get: %w", err)
	}
	if !password.ValidatePassword(user.Password, loginInput.Password) {
		return nil, fmt.Errorf("loginRepository - Login: %w", err)
	}
	return user, nil
}
