package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upRoomOccupantsTable, downRoomOccupantsTable)
}

func upRoomOccupantsTable(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE IF NOT EXISTS room_occupants (room_id INTEGER NOT NULL, user_id INTEGER NOT NULL, FOREIGN KEY (room_id) REFERENCES rooms(id), FOREIGN KEY (user_id) REFERENCES users(id));")
	if err != nil {
		return err
	}
	return nil
}

func downRoomOccupantsTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS room_occupants;")
	if err != nil {
		return err
	}
	return nil
}
