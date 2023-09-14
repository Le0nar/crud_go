package service

import (
	"crypto/sha1"
	"fmt"

	news "github.com/Le0nar/crud_go"
	"github.com/Le0nar/crud_go/pkg/repository"
)

// TODO: move salt to .env file
const salt = "asdas98das23dds22332a"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user news.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash (password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
