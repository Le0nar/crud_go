package news

import "time"

type UpdateNewsInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type News struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date   time.Time `json:"date"`
	UserId int       `json:"userId" binding:"required" db:"user_id"`
}
