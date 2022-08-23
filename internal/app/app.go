// Package app configures and runs application.
package app

import (
	"fmt"
	"net"

	"github.com/TcMits/ent-clean-template/config"
	v1 "github.com/TcMits/ent-clean-template/internal/controller/http/v1"
	"github.com/TcMits/ent-clean-template/internal/repository"
	"github.com/TcMits/ent-clean-template/internal/usecase"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/datastore"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	client, err := datastore.NewClient(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer client.Close()

	// Usecase
	loginUseCase := usecase.NewLoginUseCase(
		repository.NewLoginRepository(client),
		cfg.LoginUseCase.Secret,
	)

	// HTTP Server
	handler := iris.New()
	handler.Validator = validator.New()
	handler.I18n.Load("./locales/*/*")
	handler.I18n.SetDefault("en-US")

	v1.RegisterLoginController(handler, loginUseCase, l)
	handler.Listen(net.JoinHostPort("", cfg.HTTP.Port))
	// httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	// l.Info("Listening and serving HTTP on %s\n", httpServer.Addr())
	//
	// // Waiting signal
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	//
	// select {
	// case s := <-interrupt:
	// 	l.Info("app - Run - signal: " + s.String())
	// case err = <-httpServer.Notify():
	// 	l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	// }
	//
	// // Shutdown
	// err = httpServer.Shutdown()
	// if err != nil {
	// 	l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	// }
}
