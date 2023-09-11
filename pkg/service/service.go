package service

import (
	"github.com/Le0nar/crud_go/pkg/repository"
)

type Authorization interface {
}

type News interface {
}

type Service struct {
	Authorization
	News
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}