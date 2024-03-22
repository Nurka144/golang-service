package services

import (
	"github.com/example/test/internal/models"
	"github.com/example/test/internal/repository"
)

type Test struct {
	TestRepository repository.Test
}

func InitTestService(r repository.Test) *Test {
	return &Test{
		TestRepository: r,
	}
}

func (s Test) Create(body models.TestCreateBody, AuthData models.AuthMiddlewareData) (int, error) {
	return s.TestRepository.Create(body, AuthData)
}

func (s Test) Find(id int) (*models.TestFind, error) {
	return s.TestRepository.Find(id)
}
