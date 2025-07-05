package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/sxd0/SweetTweet/internal/auth/repository"
	"github.com/sxd0/SweetTweet/pkg/jwt"
)

type LoginService struct {
	Repo *repository.UserRepository
}

func NewLoginService(repo *repository.UserRepository) *LoginService {
	return &LoginService{Repo: repo}
}

func (s *LoginService) Login(email, password string) (string, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
