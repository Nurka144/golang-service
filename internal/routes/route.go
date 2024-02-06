package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	UserRoutes(db, router)
	BookRoutes(db, router)

	return router
}
