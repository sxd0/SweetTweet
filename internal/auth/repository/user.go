package repository

import (
	"database/sql"
	"github.com/sxd0/SweetTweet/internal/auth/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *model.User) error {
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, user.Email, user.PasswordHash)
	return err
}
