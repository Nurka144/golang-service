package services

import (
	"github.com/Nurka144/golang-service/internal/models"
	"github.com/Nurka144/golang-service/internal/repository"
)

type UserService struct {
	UserRepository repository.User
}

type UserResponse struct {
	ID int `json:"id"`
}

func NewUserService(r repository.User) *UserService {
	return &UserService{
		UserRepository: r,
	}
}

func (s *UserService) FindOne(id int) (*models.User, error) {
	return s.UserRepository.FindOne(id)
}

func (s *UserService) Create(user models.UserCreate) (int, error) {
	return s.UserRepository.Create(user)
}

func (s *UserService) FindMany() ([]models.User, error) {
	data, dataErr := s.UserRepository.FindMany()

	if dataErr != nil {
		return nil, dataErr
	}

	return data, nil
}
