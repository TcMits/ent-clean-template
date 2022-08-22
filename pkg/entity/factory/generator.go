package factory

import (
	"context"
	"fmt"
	"reflect"

	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/pkg/tool/password"
	"github.com/bluele/factory-go/factory"
)

type generator[ModelType any, MutationType ent.Mutation] struct {
	client   *factory.Factory
	mutation MutationType
}

func (g *generator[ModelType, _]) Generate(ctx context.Context, opt map[string]any) ModelType {
	return g.client.MustCreateWithContextAndOption(ctx, opt).(ModelType)
}

func (g *generator[_, MutationType]) SetMutation(mutation MutationType) {
	g.mutation = mutation
}

func prepareModelType(model any) any {
	reflectValue := reflect.ValueOf(model)
	for reflectValue.Kind() == reflect.Ptr {
		if reflectValue.IsNil() && reflectValue.CanAddr() {
			reflectValue.Set(reflect.New(reflectValue.Type().Elem()))
		}
		reflectValue = reflectValue.Elem()
	}
	return reflect.New(reflectValue.Type()).Interface()
}

func NewUserGenerator[UserType any, UserMutationType ent.Mutation]() Generator[UserType, UserMutationType] {
	g := &generator[UserType, UserMutationType]{}
	factoryClient := factory.NewFactory(prepareModelType(new(UserType))).Attr(
		"Password", func(a factory.Args) (any, error) {
			password, _ := password.GetHashPassword("12345678")
			g.mutation.SetField("password", password)
			return password, nil
		},
	).SeqInt(
		"Username", func(n int) (any, error) {
			username := fmt.Sprintf("username%d", n)
			g.mutation.SetField("username", username)
			return username, nil
		},
	).SeqInt(
		"FirstName", func(n int) (any, error) {
			firstName := fmt.Sprintf("first-name%d", n)
			g.mutation.SetField("first_name", firstName)
			return firstName, nil
		},
	).SeqInt(
		"LastName", func(n int) (any, error) {
			lastName := fmt.Sprintf("last-name%d", n)
			g.mutation.SetField("last_name", lastName)
			return lastName, nil
		},
	).SeqInt(
		"LastName", func(n int) (any, error) {
			email := fmt.Sprintf("email%d@gmail.com", n)
			g.mutation.SetField("email", email)
			return email, nil
		},
	).Attr(
		"IsStaff", func(a factory.Args) (any, error) {
			g.mutation.SetField("is_staff", true)
			return true, nil
		},
	).Attr(
		"IsActive", func(a factory.Args) (any, error) {
			g.mutation.SetField("is_active", true)
			return true, nil
		},
	)
	g.client = factoryClient
	return g
}
