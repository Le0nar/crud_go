package service

import (
	news "github.com/Le0nar/crud_go"
	"github.com/Le0nar/crud_go/pkg/repository"
)

type Authorization interface {
	CreateUser(user news.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type News interface {
	CreateNews(news news.News) (int, error)
	GetAllNews() ([]news.News, error)
	GetNewsById(newsId int) (news.News, error)
}

type Service struct {
	Authorization
	News
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		News: NewNewsService(repos.News),
	}
}
