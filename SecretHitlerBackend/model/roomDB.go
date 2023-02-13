package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func (room *Room) Create(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO rooms (created_at, code, size) VALUES (?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Exec(room.CreatedAt, room.Code, room.Size)
	roomID, err := result.LastInsertId()
	room.ID = uint(roomID)
	if err != nil {
		return err
	}

	fmt.Println(room.ID)

	return nil
}

func GetRoomByCode(code string, db *sql.DB) (Room, error) {
	stmt, err := db.Prepare("SELECT * FROM rooms WHERE code = ?")
	if err != nil {
		log.Fatal(err)
	}

	var room Room

	// Execute the prepared statement, passing in an id value for the
	// parameter whose placeholder is ?
	err = stmt.QueryRow(code).Scan(&room.ID, &room.CreatedAt, &room.Code, &room.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return Room{}, errors.New("record not found")
		}
		return room, err
	}
	return room, nil
}

func (room *Room) Save(db *sql.DB) error {
	return nil
}

func (room *Room) Delete() error {
	return nil
}
