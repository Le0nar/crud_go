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
	news.Date = time.Now()

	return s.repo.CreateNews(news)
}

func (s *NewsService) GetAllNews() ([]news.News, error) {
	return s.repo.GetAllNews()
}

func (s *NewsService) GetNewsById(newsId int) (news.News, error) {
	return s.repo.GetNewsById(newsId)
}

func (s *NewsService) UpdateNewsById(input news.UpdateNewsInput, newId int) error {
	return s.repo.UpdateNewsById(input, newId)
}

func (s *NewsService) DeleteNewsById(newsId int) error {
	return s.repo.DeleteNewsById(newsId)
}
