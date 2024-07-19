package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upUrl, downUrl)
}

func upUrl(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL)`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downUrl(tx *sql.Tx) error {
	query := `DROP TABLE url`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
