package main

import (
	"context"
	"log"

	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/ent/migrate"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/datastore"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Info("migrating...")
	client, err := datastore.NewClient(cfg.PG.URL, cfg.PG.PoolMax)
	if err != nil {
		log.Fatalf("failed opening postgres client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
	l.Info("finish migrations")
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
