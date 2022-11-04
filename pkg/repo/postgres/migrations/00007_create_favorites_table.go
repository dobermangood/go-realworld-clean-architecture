package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00007, Down00007)
}

func Up00007(tx *sql.Tx) error {
	sql := `
		CREATE TABLE IF NOT EXISTS favorites (
			article_id BIGINT NOT NULL,
			user_id BIGINT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),

			PRIMARY KEY (article_id, user_id),
			CONSTRAINT fk_article FOREIGN KEY(article_id) REFERENCES articles(id) ON DELETE CASCADE,
			CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`

	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Down00007(tx *sql.Tx) error {
	sql := `
		DROP TABLE IF EXISTS favorites;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
