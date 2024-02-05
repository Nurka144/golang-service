package controllers

import (
	"net/http"

	"github.com/Nurka144/golang-service/internal/models"
	"github.com/Nurka144/golang-service/internal/services"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *services.BookService
}

func NewBookController(s *services.BookService) *BookController {
	return &BookController{
		BookService: s,
	}
}

func (c BookController) Create(ctx *gin.Context) {
	var body models.BookCreate

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{"status": "error", "error": err})
		return
	}

	bookId, error := c.BookService.Create(body)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{"status": "error", "error": error})
		return
	}

	ctx.JSON(http.StatusCreated, &gin.H{"status": "success", "data": bookId})

}
