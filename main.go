package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/ping", testHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id", bookHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", addBookHandler)

	// router.Run() // default
	router.Run(":8888") // changed port
}

func bookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type BookInput struct {
	Title       string      `json:"title" binding:"required"`
	SubTitle    string      `json:"sub_title"`
	Price       interface{} `json:"price" binding:"required,numeric"`
	Description string
}

func addBookHandler(ctx *gin.Context) {
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {
		errors := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errors = append(errors, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":       bookInput.Title,
		"sub_title":   bookInput.SubTitle,
		"price":       bookInput.Price,
		"description": bookInput.Description,
	})
}

func queryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func testHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    "muhammad iqbal",
		"address": "jakarta barat kedoya",
	})
}
