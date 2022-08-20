package factory

import (
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	"github.com/bluele/factory-go/factory"
)

var UserFactory = factory.NewFactory(&model.User{})
