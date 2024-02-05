package services

import (
	"github.com/Nurka144/golang-service/internal/models"
	"github.com/Nurka144/golang-service/internal/repository"
)

type BookService struct {
	BookRepository repository.Book
}

func NewBookService(r repository.Book) *BookService {
	return &BookService{
		BookRepository: r,
	}
}

func (s *BookService) Create(book models.BookCreate) (int, error) {
	return s.BookRepository.Create(book)
}