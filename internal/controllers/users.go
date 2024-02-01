package controllers

import (
	"net/http"
	"strconv"

	"github.com/Nurka144/golang-service/internal/models"
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

func (c *UserController) FindMany(ctx *gin.Context) {
	users, err := c.UserService.FindMany()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "data": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}

func (c *UserController) FindOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "data": err})
		return
	}

	res, errFind := c.UserService.FindOne(id)

	if errFind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": errFind.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
}

func (c *UserController) Create(ctx *gin.Context) {
	var body models.UserCreate

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err})
		return
	}

	userId, err := c.UserService.Create(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": userId})
}
