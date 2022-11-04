package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00004, Down00004)
}

func Up00004(tx *sql.Tx) error {
	sql := `
		CREATE TABLE IF NOT EXISTS article_tags (
			article_id BIGINT NOT NULL,
			tag_id BIGINT NOT NULL,

			PRIMARY KEY (article_id, tag_id),
			CONSTRAINT fk_article FOREIGN KEY(article_id) REFERENCES articles(id) ON DELETE CASCADE,
			CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
		);
	`

	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Down00004(tx *sql.Tx) error {
	sql := `
		DROP TABLE IF EXISTS article_tags;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
