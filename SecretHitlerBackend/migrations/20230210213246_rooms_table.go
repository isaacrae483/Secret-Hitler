package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upRoomsTable, downRoomsTable)
}

func upRoomsTable(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE IF NOT EXISTS rooms (id INTEGER PRIMARY KEY, created_at DATETIME NOT NULL, code TEXT NOT NULL, size INTEGER NOT NULL);")
	if err != nil {
		return err
	}
	return nil
}

func downRoomsTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS rooms;")
	if err != nil {
		return err
	}
	return nil
}
