package models

import (
	"time"
)

type Rental struct {
	Id               uint
	Item_id          uint
	User_id          uint
	Start_date       time.Time
	End_date         time.Time
	Current_state    string
	Esstimated_time  int
	Delivery_address string
	Created_at       time.Time
}
