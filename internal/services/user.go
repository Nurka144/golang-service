package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

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

	var UserBodyResp UserResponse

	errorBodyUnmarshal := json.Unmarshal(body, &UserBodyResp)

	if errorBodyUnmarshal != nil {
		return nil, errorBodyUnmarshal
	}

	fmt.Println("Ответ сервера", UserBodyResp.ID)

	return s.UserRepository.FindOne(id)
}

func (s *UserService) Create(user models.UserCreate) (int, error) {
	return s.UserRepository.Create(user)
}

func fetchData(userId int, wg *sync.WaitGroup, ch chan<- UserResponse) {
	defer wg.Done()
	url := "https://jsonplaceholder.typicode.com/users/" + strconv.Itoa(userId)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)

	if er != nil {
		fmt.Println(er)
		return
	}

	var UserBodyResp UserResponse

	errorBodyUnmarshal := json.Unmarshal(body, &UserBodyResp)

	if errorBodyUnmarshal != nil {
		fmt.Println(errorBodyUnmarshal)
		return
	}

	ch <- UserBodyResp
}

func (s *UserService) FindMany() ([]models.User, error) {
	data, dataErr := s.UserRepository.FindMany()

	if dataErr != nil {
		return nil, dataErr
	}

	var wg sync.WaitGroup
	ch := make(chan UserResponse, len(data))

	for _, usr := range data {
		wg.Add(1)
		go fetchData(usr.ID, &wg, ch)
	}

	wg.Wait()
	close(ch)

	// dataCh, errCh := <-ch

	// if !errCh {
	// 	fmt.Println(errCh)
	// }

	// fmt.Println(dataCh)

	for dataCh := range ch {
		// Обрабатываем данные, например, распечатываем их
		fmt.Println(dataCh)
	}

	return data, nil
}
