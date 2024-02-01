package routes

import (
	"database/sql"

	"github.com/Nurka144/golang-service/internal/routes/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	user.UserRoutes(db, router)

	return router
}
