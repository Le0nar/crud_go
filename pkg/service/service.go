package service

import (
	news "github.com/Le0nar/crud_go"
	"github.com/Le0nar/crud_go/pkg/repository"
)

type Authorization interface {
	CreateUser(user news.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type News interface {
}

type Service struct {
	Authorization
	News
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}