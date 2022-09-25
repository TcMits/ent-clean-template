package datastore

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/TcMits/ent-clean-template/ent"
)

// Open new connection.
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

// NewClient returns an orm client.
func NewClient(url string, poolMax int) (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	return Open(url, poolMax)
}
