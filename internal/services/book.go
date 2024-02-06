package services

import (
	"github.com/Nurka144/golang-service/internal/models"
	"github.com/Nurka144/golang-service/internal/repository"
)

type BookService struct {
	BookRepository repository.BookRepository
}

func NewBookService(r repository.BookRepository) *BookService {
	return &BookService{
		BookRepository: r,
	}
}

func (s *BookService) Create(book models.BookCreate) (int, error) {
	return s.BookRepository.Create(book)
}
