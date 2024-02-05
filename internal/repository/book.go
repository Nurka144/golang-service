package repository

import (
	"database/sql"

	"github.com/Nurka144/golang-service/internal/models"
)

type Book interface {
	Create(book models.BookCreate) (int, error)
}

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) Book {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Create(book models.BookCreate) (int, error) {
	return 1, nil
}
