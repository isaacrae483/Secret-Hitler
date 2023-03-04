package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upUsersTable, downUsersTable)
}

func upUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, created_at DATETIME NOT NULL, name TEXT NOT NULL, guest BOOLEAN default false);")
	if err != nil {
		return err
	}
	return nil
}

func downUsersTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		return err
	}
	return nil
}
