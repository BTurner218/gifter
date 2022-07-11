package postgres

import (
	"context"

	"github.com/BTurner218/gifter/api/gifter"
)

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db}
}

func (us *UserService) CreateUser(user *gifter.User) error {
	err := createUser(user, us.db)

	if err != nil {
		return err
	}

	return nil
}

func createUser(user *gifter.User, db *DB) error {
	query := `
	INSERT INTO users (username, email, password)
	VALUES ($1, $2, $3)
	`
	args := []interface{}{user.Username, user.Email, user.Password}
	_, err := db.Exec(context.Background(), query, args...)

	return err
}
