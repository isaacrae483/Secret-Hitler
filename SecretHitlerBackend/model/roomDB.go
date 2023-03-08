package model

import (
	"database/sql"
	"errors"
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

	return nil
}

func GetAvailableRooms(db *sql.DB) ([]Room, error) {
	//TODO: doesn't get room if 0 occupants, tho that shouldn't happen
	stmt, err := db.Prepare("select * from (select rooms.*, count(ro.room_id) as occupant_count from rooms join room_occupants ro on rooms.id = ro.room_id group by rooms.id) as rooms where rooms.size > rooms. occupant_count")
	if err != nil {
		log.Fatal(err)
	}

	var rooms []Room

	rows, err := stmt.Query()
	for rows.Next() {
		var room Room
		var count int64
		err := rows.Scan(&room.ID, &room.CreatedAt, &room.Code, &room.Size, &count)
		if err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, room)
	}
	return rooms, err
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

func (room *Room) OccupantCount(db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM room_occupants WHERE room_id = ?")
	if err != nil {
		log.Fatal(err)
	}

	var count int64

	// Execute the prepared statement, passing in an id value for the
	// parameter whose placeholder is ?
	err = stmt.QueryRow(room.ID).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("record not found")
		}
		return 0, err
	}
	return count, nil
}

func (room *Room) Delete() error {
	return nil
}

func (room *Room) AddPerson(userID uint, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO room_occupants (room_id, user_id) VALUES (?,?);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(room.ID, userID)
	if err != nil {
		return err
	}

	return nil
}
