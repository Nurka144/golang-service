package book

import (
	"database/sql"

	"github.com/Nurka144/golang-service/internal/controllers"
	"github.com/Nurka144/golang-service/internal/repository"
	"github.com/Nurka144/golang-service/internal/services"
	"github.com/gin-gonic/gin"
)

func BookRoutes(db *sql.DB, route *gin.Engine) *gin.Engine {
	bookRepository := repository.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookController := controllers.NewBookController(bookService)

	bookApi := route.Group("/book")
	{
		bookApi.POST("/", bookController.Create)
	}

	return route
}
