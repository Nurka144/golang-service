package routes

import (
	"github.com/example/test/docs"
	"github.com/example/test/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func InitRoutes(db *sqlx.DB) *gin.Engine {

	router := gin.Default()

	router.Use(docs.InitApiDocs())

	routerGroup := router.Group("/api/test-be/" + viper.GetString("apiVersion"))
	routerGroup.Use(middleware.AuthMiddleware())
	InitBookRoutes(db, routerGroup)

	return router
}
