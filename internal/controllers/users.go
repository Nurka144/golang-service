package controllers

import (
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "data": err})
	}

	res, _ := c.UserService.FindOne(id)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}
