package models

import "time"

type Review struct {
	ID        uint
	ItemID    uint
	UserID    uint
	Rating    int
	Comment   string
	CreatedAt time.Time
}
