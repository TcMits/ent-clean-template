package main

import (
	"context"
	"flag"
	"log"

	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/pkg/infrastructure/datastore"
	"github.com/TcMits/ent-clean-template/pkg/tool/password"
)

func main() {
	isStaff := flag.Bool("isStaff", true, "Create staff user")
	isSuperuser := flag.Bool("isSuperuser", true, "Create superuser")
	email := flag.String("email", "eddiebrock2001@gmail.com", "User's email")
	username := flag.String("username", "TcMits", "User's username")
	pass := flag.String("password", "We love golang", "User's password")
	flag.Parse()
	hashedPassword, err := password.GetHashPassword(*pass)
	if err != nil {
		log.Fatalf("Hashing password error: %s", err)
	}

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	client, err := datastore.NewClient(cfg.PG.URL, cfg.PG.PoolMax)
	if err != nil {
		log.Fatalf("failed opening postgres client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	client.User.Create().SetIsStaff(
		*isStaff,
	).SetIsSuperuser(
		*isSuperuser,
	).SetEmail(
		*email,
	).SetUsername(
		*username,
	).SetPassword(
		hashedPassword,
	).Save(ctx)
}
