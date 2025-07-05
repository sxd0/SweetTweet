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


func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	query := "SELECT id, email, password_hash FROM users WHERE email = $1"
	row := r.DB.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
