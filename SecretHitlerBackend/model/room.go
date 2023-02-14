package model

import (
	"SecretHitlerBackend/utils"
	"database/sql"
	"errors"
	"time"
)

type Room struct {
	ID        uint
	CreatedAt time.Time
	Code      string
	Size      int64
}

func CreateRoom(db *sql.DB) (string, error) {
	room := Room{
		CreatedAt: time.Now(),
		Code:      utils.RandString(5),
		Size:      6,
	}

	if err := room.Create(db); err != nil {
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
