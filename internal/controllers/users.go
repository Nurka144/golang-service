package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
