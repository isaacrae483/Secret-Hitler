package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upSessionsTable, downSessionsTable)
}

func upSessionsTable(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE IF NOT EXISTS sessions (id INTEGER PRIMARY KEY, created_at DATETIME NOT NULL, user_id INTEGER, key TEXT UNIQUE NOT NULL, FOREIGN KEY (user_id) REFERENCES users(id));")
	if err != nil {
		return err
	}
	return nil
}

func downSessionsTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS sessions;")
	if err != nil {
		return err
	}
	return nil
}
