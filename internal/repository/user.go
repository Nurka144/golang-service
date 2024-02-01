package repository

import (
	"database/sql"
	"errors"

	"github.com/Nurka144/golang-service/internal/models"
)

type User interface {
	FindOne(id int) (*models.User, error)
	Create(user models.UserCreate) (int, error)
	FindMany() ([]models.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) User {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindMany() ([]models.User, error) {

	sql := "select id, username, age, email from bh.user"
	rows, err := r.db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Age, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindOne(userId int) (*models.User, error) {
	var id int
	var name string
	var age int
	var email string

	err := r.db.QueryRow("select id, username, age, email from bh.user where id = $1", userId).Scan(&id, &name, &age, &email)

	if err != nil {
		return nil, errors.New("Получена ошибка при получение данных : " + err.Error())
	}

	return &models.User{
		ID:       id,
		Username: name,
		Age:      age,
		Email:    email,
	}, nil
}

func (r *UserRepository) Create(user models.UserCreate) (int, error) {
	var userId int = 0
	sql := `insert into bh.user(username, age, email) values ($1, $2, $3) returning id`
	err := r.db.QueryRow(sql, user.Username, user.Age, user.Email).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
