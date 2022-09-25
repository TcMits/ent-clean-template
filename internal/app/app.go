// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/kataras/iris/v12/middleware/recover"

	"github.com/TcMits/ent-clean-template/config"
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/controller/http/v1/middleware"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/entity/model"
	useCaseModel "github.com/TcMits/ent-clean-template/pkg/entity/model/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/datastore"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/httpserver"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	client, err := datastore.NewClient(cfg.PG.URL, cfg.PG.PoolMax)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer client.Close()

	// repository
	userRepository := repository.NewUserRepository(client)
	loginRepository := repository.NewLoginRepository(client)

	// HTTP Server
	handler := v1.NewHandler()

	// Usecase
	loginUseCase := usecase.NewLoginUseCase(
		loginRepository,
		userRepository,
		cfg.LoginUseCase.Secret,
	)
	publicMeUseCase := usecase.NewPublicMeUseCase(
		userRepository, userRepository, tool.GetIrisReverseFunc("publicMe", handler),
	)

	// RegisterV1HTTPServices
	registerV1HTTPServices(handler, loginUseCase, publicMeUseCase, l)

	if err := handler.Build(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Build: %w", err))
	}
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Info("Listening and serving HTTP on %s", httpServer.Addr())

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func registerV1HTTPServices(
	handler iris.Party,
	// login route
	loginUseCase usecase.LoginUseCase[*useCaseModel.LoginInput, *useCaseModel.JWTAuthenticatedPayload, *useCaseModel.RefreshTokenInput, *model.User],
	// publicMe route
	publicMeUseCase interface {
		usecase.GetModelUseCase[*model.User, *struct{}]
		usecase.GetAndUpdateModelUseCase[*model.User, *struct{}, *useCaseModel.PublicMeUseCaseUpdateInput]
		usecase.SerializeModelUseCase[*model.User, map[string]any]
	},
	// adding more usecases here
	// logger
	l logger.Interface,
) {
	handler.UseRouter(recover.New())
	v1.RegisterHealthCheckController(handler)

	// HTTP middlewares
	h := handler.Party(
		"/v1",
		cors.New().Handler(),
		middleware.Logger(l),
		middleware.Common(),
		middleware.Auth(loginUseCase),
	)
	{
		// routes
		v1.RegisterLoginController(h, loginUseCase, l)
		// protected routes
		v1.RegisterPublicMeController(h, publicMeUseCase, publicMeUseCase, publicMeUseCase, l)
	}
}
