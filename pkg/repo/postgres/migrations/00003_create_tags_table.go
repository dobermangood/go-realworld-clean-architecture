package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00003, Down00003)
}

func Up00003(tx *sql.Tx) error {
	sql := `
		CREATE TABLE IF NOT EXISTS tags (
			id BIGSERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL
		);
	`

	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Down00003(tx *sql.Tx) error {
	sql := `
		DROP TABLE IF EXISTS tags;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
