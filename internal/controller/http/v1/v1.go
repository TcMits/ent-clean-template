package v1

import (
	"github.com/TcMits/ent-clean-template/internal/controller/http/v1/middleware"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
)

const _v1SubPath = "/api/v1"

func NewHandler() *iris.Application {
	handler := iris.New()

	// validator
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func RegisterV1HTTPServices(
	handler iris.Party,
	// adding more usecases here
	loginUseCase usecase.LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User],
	publicMeUseCase interface {
		usecase.GetModelUseCase[*model.User, *struct{}]
		usecase.GetAndUpdateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput]
		usecase.SerializeModelUseCase[*model.User, map[string]any]
	},
	// logger
	l logger.Interface,
) {
	handler.UseRouter(middleware.Recover(handleError, l))
	RegisterHealthCheckController(handler)

	// HTTP middlewares
	h := handler.Party(
		_v1SubPath,
		cors.New().Handler(),
		middleware.Logger(l),
		middleware.Common(),
		middleware.Auth(loginUseCase),
	)
	// routes
	{
		RegisterLoginController(h, loginUseCase, l)
		RegisterPublicMeController(h, publicMeUseCase, publicMeUseCase, publicMeUseCase, l)
		RegisterDocsController(h, l)
	}
}
