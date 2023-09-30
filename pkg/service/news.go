package service

import (
	"time"

	news "github.com/Le0nar/crud_go"
	"github.com/Le0nar/crud_go/pkg/repository"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) CreateNews(news news.News) (int, error) {
	dateNow := time.Now()
	news.Date = dateNow.String()

	return s.repo.CreateNews(news)
}

func (s *NewsService) GetAllNews() ([]news.News, error) {
	return s.repo.GetAllNews()
}
