package user

import (
	"database/sql"

	"github.com/Nurka144/golang-service/internal/controllers"
	"github.com/Nurka144/golang-service/internal/repository"
	"github.com/Nurka144/golang-service/internal/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(db *sql.DB, route *gin.Engine) *gin.Engine {
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	usersApi := route.Group("/user")
	{
		usersApi.GET("/", userController.FindMany)
		usersApi.GET("/:id", userController.FindOne)
		usersApi.POST("/", userController.Create)
	}

	return route
}
