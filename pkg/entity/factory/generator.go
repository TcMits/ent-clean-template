package factory

import (
	"context"
	"reflect"

	"github.com/TcMits/ent-clean-template/ent"
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
