package models

import (
	"time"
)

type User struct {
	Id         uint
	First_name string
	Last_name  string
	Role       string
	Gender     string
	State      string
	City       string
	Street     string
	Score      int
	Email      string
	Password   string
}

type User_phone struct {
	User_id    uint
	User_phone string
}
