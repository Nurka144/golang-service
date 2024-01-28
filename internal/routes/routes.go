package routes

import (
	"github.com/Nurka144/golang-service/internal/controllers"
	"github.com/Nurka144/golang-service/internal/repository"
	"github.com/Nurka144/golang-service/internal/services"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	router := gin.New()

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	usersApi := router.Group("/user")
	{
		usersApi.GET("/:id", userController.FindOne)
	}

	return router
}
