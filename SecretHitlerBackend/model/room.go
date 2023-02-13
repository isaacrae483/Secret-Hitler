package model

import (
	"SecretHitlerBackend/utils"
	"database/sql"
	"time"
)

type Room struct {
	ID        uint
	CreatedAt time.Time
	Code      string
	Size      int
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
