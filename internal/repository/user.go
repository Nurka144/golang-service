package repository

import "github.com/Nurka144/golang-service/internal/models"

type User interface {
	FindOne(id int) (*models.User, error)
}

type UserRepository struct {
	User
}

func NewUserRepository() User {
	return &UserRepository{}
}

func (r *UserRepository) FindOne(id int) (*models.User, error) {
	return &models.User{
		ID:       1,
		Username: "trst",
		Age:      12,
	}, nil
}
