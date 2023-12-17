package services

import (
	"github.com/Nurka144/golang-service/internal/models"
	"github.com/Nurka144/golang-service/internal/repository"
)

type UserService struct {
	UserRepository repository.User
}

func NewUserService(r repository.User) *UserService {
	return &UserService{
		UserRepository: r,
	}
}

func (s *UserService) FindOne(id int) (*models.User, error) {
	return s.UserRepository.FindOne(id)
}
