package repository

import (
	news "github.com/Le0nar/crud_go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user news.User) (int, error)
}

type News interface {
}

type Repository struct {
	Authorization
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
