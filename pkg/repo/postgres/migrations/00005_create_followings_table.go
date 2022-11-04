package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00005, Down00005)
}

func Up00005(tx *sql.Tx) error {
	sql := `
		CREATE TABLE IF NOT EXISTS followings (
			following_id BIGINT NOT NULL,
			follower_id BIGINT NOT NULL,
			followed_on TIMESTAMP NOT NULL DEFAULT NOW(),

			PRIMARY KEY (following_id, follower_id),
			CONSTRAINT fk_following FOREIGN KEY(following_id) REFERENCES users(id),
			CONSTRAINT fk_follower FOREIGN KEY(follower_id) REFERENCES users(id)
		);
	`

	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func Down00005(tx *sql.Tx) error {
	sql := `
		DROP TABLE IF EXISTS followings;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
