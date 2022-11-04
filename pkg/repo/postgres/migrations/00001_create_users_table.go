package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	sql := `
		CREATE TABLE IF NOT EXISTS users(
			id BIGSERIAL PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			bio TEXT,
			image TEXT, 
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`

	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Down00001(tx *sql.Tx) error {
	sql := `
		DROP TABLE IF EXISTS users;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
