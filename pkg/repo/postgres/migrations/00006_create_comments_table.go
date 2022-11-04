package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00006, Down00006)
}

func Up00006(tx *sql.Tx) error {
	sql := `
		CREATE TABLE IF NOT EXISTS comments (
			id BIGSERIAL PRIMARY KEY,
			article_id BIGINT NOT NULL,
			author_id BIGINT NOT NULL,
			body text NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),

			CONSTRAINT fk_article FOREIGN KEY(article_id) REFERENCES articles(id) ON DELETE CASCADE,
			CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`

	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Down00006(tx *sql.Tx) error {
	sql := `
		DROP TABLE IF EXISTS comments;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
