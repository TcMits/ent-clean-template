package factory

import (
	"fmt"

	"github.com/bluele/factory-go/factory"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/TcMits/ent-clean-template/pkg/tool/password"
)

var _userFactory = modelFactory[
	*model.User,
	*model.UserMutation,
	*model.UserCreateInput,
	*model.UserCreate,
]{generator: newUserGenerator[*model.UserMutation, *model.UserCreateInput]()}

func newUserGenerator[
	MutationType ent.Mutation,
	PMutationInputType interface {
		model.MutationInput[MutationType]

		*MutationInputType
	},
	MutationInputType any,
]() Generator[MutationType, PMutationInputType] {
	g := &generator[MutationType, PMutationInputType]{}
	factoryClient := factory.NewFactory(
		PMutationInputType(new(MutationInputType)),
	).Attr(
		"Password", func(a factory.Args) (any, error) {
			password, err := password.GetHashPassword("12345678")

			return &password, err
		},
	).SeqInt(
		"Username", func(n int) (any, error) {
			username := fmt.Sprintf("username%d", n)

			return username, nil
		},
	).SeqInt(
		"FirstName", func(n int) (any, error) {
			firstName := fmt.Sprintf("first-name%d", n)

			return &firstName, nil
		},
	).SeqInt(
		"LastName", func(n int) (any, error) {
			lastName := fmt.Sprintf("last-name%d", n)

			return &lastName, nil
		},
	).SeqInt(
		"Email", func(n int) (any, error) {
			email := fmt.Sprintf("email%d@gmail.com", n)

			return email, nil
		},
	).Attr(
		"IsStaff", func(a factory.Args) (any, error) {
			isStaff := true

			return &isStaff, nil
		},
	).Attr(
		"IsActive", func(a factory.Args) (any, error) {
			isActive := true

			return &isActive, nil
		},
	)
	g.client = factoryClient

	return g
}

func GetUserFactory() ModelFactory[
	*model.User,
	*model.UserMutation,
	*model.UserCreateInput,
	*model.UserCreate,
] {
	return &_userFactory
}
