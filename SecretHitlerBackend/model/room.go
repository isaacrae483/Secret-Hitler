package model

import (
	"SecretHitlerBackend/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type Room struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"created_at"`
	Code      string    `json:"code"`
	Size      int64     `json:"size"`
}

func (room *Room) MarshalJSON() ([]byte, error) {
	type Alias Room
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(room),
		CreatedAt: room.CreatedAt.Format("2006-01-02T15:04:05+0000"),
	})
}

func CreateRoom(user User, db *sql.DB) (string, error) {
	room := Room{
		CreatedAt: time.Now(),
		Code:      utils.RandUpperString(5),
		Size:      10,
	}

	if err := room.Create(db); err != nil {
		return "", err
	}

	if err := room.AddPerson(user.ID, db); err != nil {
		return "", err
	}

	return room.Code, nil
}

type JoinRoomInput struct {
	Code string `json:"code" binding:"required"`
}

func (jri *JoinRoomInput) Join(db *sql.DB) error {
	room, err := GetRoomByCode(jri.Code, db)
	if err != nil {
		return err
	}

	occupantCount, err := room.OccupantCount(db)
	if occupantCount >= room.Size {
		return errors.New("room full")
	}

	return room.AddPerson(1, db)
}
