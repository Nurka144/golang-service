package repository

import (
	"database/sql"

	"github.com/Nurka144/golang-service/internal/models"
)

type User interface {
	FindOne(id int) (*models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) User {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindOne(userId int) (*models.User, error) {
	var id int
	var name string
	var age int
	var email string

	err := r.db.QueryRow("select id, name, age, email from bh.user where id = $1", userId).Scan(&id, &name, &age, &email)

	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:       id,
		Username: name,
		Age:      age,
	}, nil
}
