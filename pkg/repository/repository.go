package repository

type Authorization interface {
}

type News interface {
}

type Repository struct {
	Authorization
	News
}

func NewRepository() *Repository {
	return &Repository{}
}