// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TcMits/ent-clean-template/config"
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/datastore"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/httpserver"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/TcMits/ent-clean-template/pkg/tool"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	client, err := datastore.NewClient(cfg.PG.URL, cfg.PG.PoolMax)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer client.Close()

	// repository
	userRepository := repository.NewUserRepository(client)

	// HTTP Server
	handler := v1.NewHandler()

	// Usecase
	loginUseCase := usecase.NewLoginUseCase(
		userRepository,
		cfg.LoginUseCase.Secret,
	)
	publicMeUseCase := usecase.NewPublicMeUseCase(
		userRepository, userRepository, tool.GetIrisReverseFunc("publicMe", handler),
	)

	v1.RegisterV1HTTPServices(handler, loginUseCase, publicMeUseCase, l)

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
