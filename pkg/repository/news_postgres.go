package repository

import (
	"fmt"

	news "github.com/Le0nar/crud_go"
	"github.com/jmoiron/sqlx"
)

type NewsPostgres struct {
	db *sqlx.DB
}

func NewNewsPostgres(db *sqlx.DB) *NewsPostgres {
	return &NewsPostgres{db: db}
}

func (r *NewsPostgres) CreateNews(news news.News) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (date, description, title, user_id) values ($1, $2, $3, $4) RETURNING id", newsTable)

	row := r.db.QueryRow(query, news.Date, news.Description, news.Title, news.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *NewsPostgres) GetAllNews() ([]news.News, error) {
	var newsList []news.News

	// TODO: ошибка вроде бы здесь в query

	err := r.db.Select(&newsList, "SELECT * FROM news ")

	return newsList, err
}
