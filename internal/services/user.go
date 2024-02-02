package services

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
	url := "https://jsonplaceholder.typicode.com/users/1"
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)

	if er != nil {
		return nil, er
	}

	fmt.Println("Ответ сервера", string(body.id))

	return s.UserRepository.FindOne(id)
}

func (s *UserService) Create(user models.UserCreate) (int, error) {
	return s.UserRepository.Create(user)
}

func (s *UserService) FindMany() ([]models.User, error) {
	return s.UserRepository.FindMany()
}
