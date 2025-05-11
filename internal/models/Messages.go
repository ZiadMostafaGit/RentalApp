package models

import (
	"time"
)

type Messages struct {
	Id              uint
	Conversation_id uint
	Content         string
	Sender_id       uint
	Created_at      time.Time
}
