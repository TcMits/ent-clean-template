package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

func NewHandler() *iris.Application {
	handler := iris.New()

	// validatior
	handler.Validator = validator.New()

	// i18n
	handler.I18n.DefaultMessageFunc = func(
		langInput, langMatched, key string, args ...any,
	) string {
		return ""
	}
	err := handler.I18n.Load("./locales/*/*")
	if err != nil {
		panic(err)
	}
	handler.I18n.SetDefault("en-US")

	return handler
}
