package postgres

import (
	"context"

	"github.com/BTurner218/gifter/api/gifter"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	args := []interface{}{user.Username, user.Email, string(hashedPassword)}
	_, err := db.Exec(context.Background(), query, args...)

	return err
}
