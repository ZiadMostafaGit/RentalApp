package repository

import (
	"database/sql"
)

type repo struct {
	DB *sql.DB
}

func newrepo(DB *sql.DB) *repo {
	return &repo{
		DB,
	}
}
