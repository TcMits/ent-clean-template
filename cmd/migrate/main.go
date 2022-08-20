package main

import (
	"context"
	"log"

	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/ent"
	"github.com/TcMits/ent-clean-template/ent/migrate"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/datastore"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	client, err := datastore.NewClient(cfg)
	if err != nil {
		log.Fatalf("failed opening postgres client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
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
