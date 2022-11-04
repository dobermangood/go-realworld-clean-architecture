package postgres

import (
	"database/sql"

	"github.com/dobermangood/go-realworld-clean-architecture/config"
	_ "github.com/dobermangood/go-realworld-clean-architecture/pkg/repo/postgres/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func RunMigrations(conf config.Postgres) error {
	db, err := sql.Open("postgres", conf.ConnString)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "."); err != nil {
		return err
	}

	return nil
}
