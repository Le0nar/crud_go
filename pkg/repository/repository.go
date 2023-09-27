package repository

import (
	news "github.com/Le0nar/crud_go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user news.User) (int, error)
	GetUser(username, password string) (news.User, error)
}

type News interface {
	CreateNews(news news.News) (int, error)
}

type Repository struct {
	Authorization
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		News: NewNewsPostgres(db),
	}
}
