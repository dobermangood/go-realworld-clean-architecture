package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) error {
	sql := `
		CREATE TABLE IF NOT EXISTS articles (
			id BIGSERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			body TEXT NOT NULL,
			description TEXT,
			slug VARCHAR(255) NOT NULL UNIQUE,
			author_id  BIGINT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

			CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES users(id)
		);
	`

	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Down00002(tx *sql.Tx) error {
	sql := `
		DROP TABLE IF EXISTS articles;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
