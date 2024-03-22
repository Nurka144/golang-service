package controllers

import (
	"net/http"
	"strconv"

	"github.com/example/test/internal/models"
	"github.com/example/test/internal/services"
	"github.com/gin-gonic/gin"
)

type CreateTestResponse struct {
	ID int `json:"id"`
}

type Test struct {
	TestService *services.Test
}

func InitTestController(s *services.Test) *Test {
	return &Test{
		TestService: s,
	}
}

func (c Test) Create(ctx *gin.Context) {

	AuthData := ctx.MustGet("AuthData").(*models.AuthMiddlewareData)

	var body models.TestCreateBody

	if errBody := ctx.BindJSON(&body); errBody != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": errBody.Error()})
		return
	}

	id, err := c.TestService.Create(body, *AuthData)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &gin.H{"errCode": "0", "errMsg": "Успешно", "data": &CreateTestResponse{ID: id}})
}

func (c Test) Find(ctx *gin.Context) {

	idparam, errParam := strconv.Atoi(ctx.Param("todoId"))

	if errParam != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "data": errParam.Error()})
		return
	}

	result, error := c.TestService.TestRepository.Find(idparam)

	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": error.Error()})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Не найдено"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
