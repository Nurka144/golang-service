package routes

import (
	"database/sql"

	"github.com/Nurka144/golang-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.HttpLogger())
	router.Use(middleware.AuthMiddleware())
	UserRoutes(db, router)
	BookRoutes(db, router)

	return router
}
