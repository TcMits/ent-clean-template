package datastore

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/TcMits/ent-clean-template/config"
	"github.com/TcMits/ent-clean-template/ent"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Open new connection
func Open(databaseUrl string, maxPoolSize int) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxPoolSize)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

// NewClient returns an orm client
func NewClient(cfg *config.Config) (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	return Open(cfg.PG.URL, cfg.PG.PoolMax)
}
