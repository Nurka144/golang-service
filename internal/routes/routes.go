package routes

import (
	users "github.com/Nurka144/golang-service/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {

	router := gin.New()

	routes := router.Group("/api/v1.0")
	{
		usersApi := routes.Group("/user")
		{
			usersApi.GET("/:id", users.FindOne)
		}
	}

	return router
}
