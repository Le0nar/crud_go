package repository

import (
	"fmt"
	"strings"

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
	query := fmt.Sprintf("SELECT * FROM %s", newsTable)
	err := r.db.Select(&newsList, query)

	return newsList, err
}

func (r *NewsPostgres) GetNewsById(newsId int) (news.News, error) {
	var newsItem news.News
	query := fmt.Sprintf("SELECT * FROM %s where id = %d", newsTable, newsId)
	err := r.db.Get(&newsItem, query)

	return newsItem, err
}

func (r *NewsPostgres) UpdateNewsById(input news.UpdateNewsInput, newsId int) error {
	changedFieldList := make([]string, 0)

	if input.Title != nil {
		value := fmt.Sprintf("title='%s'", *input.Title)
		changedFieldList = append(changedFieldList, value)
	}

	if input.Description != nil {
		value := fmt.Sprintf("description='%s'", *input.Description)
		changedFieldList = append(changedFieldList, value)
	}

	queryFields := strings.Join(changedFieldList, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d", newsTable, queryFields, newsId)

	_, err := r.db.Exec(query)

	return err
}

func (r *NewsPostgres) DeleteNewsById(newsId int) error {
	query := fmt.Sprintf("DELETE FROM %s where id = %d", newsTable, newsId)
	_, err := r.db.Exec(query)
	
	return err
}
