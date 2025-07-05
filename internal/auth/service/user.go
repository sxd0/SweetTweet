package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/sxd0/SweetTweet/internal/auth/model"
	"github.com/sxd0/SweetTweet/internal/auth/repository"
)

type RegisterService struct {
	Repo *repository.UserRepository
}

func NewRegisterService(repo *repository.UserRepository) *RegisterService {
	return &RegisterService{Repo: repo}
}

func (s *RegisterService) Register(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Email:        email,
		PasswordHash: string(hash),
	}
	return s.Repo.Create(user)
}
