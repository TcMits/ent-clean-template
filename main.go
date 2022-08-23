package main

import (
	"log"

	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
