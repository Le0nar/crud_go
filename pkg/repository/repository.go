package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type News interface {
}

type Repository struct {
	Authorization
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
