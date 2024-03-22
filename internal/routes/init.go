package routes

import (
	"github.com/example/test/internal/controllers"
	"github.com/example/test/internal/repository"
	"github.com/example/test/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRoutes(db *sqlx.DB, route *gin.RouterGroup) *gin.RouterGroup {
	repository := repository.InitTestRepository(db)
	service := services.InitTestService(repository)
	controller := controllers.InitTestController(service)

	bookApi := route.Group("/todo")
	{
		bookApi.POST("/", controller.Create)
		bookApi.GET("/:todoId", controller.Find)
	}

	return route
}
