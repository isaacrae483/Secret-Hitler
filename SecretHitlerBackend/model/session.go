package model

import "time"

type Session struct {
	ID        uint
	CreatedAt time.Time
	UserID    uint
	Key       string
}
