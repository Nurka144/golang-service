package controllers

import (
	"net/http"

	"github.com/Nurka144/golang-service/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) FindOne(ctx *gin.Context) {
	res, _ := c.UserService.FindOne(1)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}
