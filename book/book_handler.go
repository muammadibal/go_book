package book

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"pustaka-api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// bookHandler => private
// BookHandler => public

type handler struct {
	bookService Service
}

func AssignHandler(bookService Service) *handler {
	return &handler{bookService}
}

func (h *handler) BookHandler(ctx *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []BookResponse
	for _, v := range books {
		bookResponse := BooksResponse(v)
		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *handler) DetailBookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "Not valid type",
		})
		return
	}

	if helpers.IsNumeric(id) {
		book, err := h.bookService.FindById(intVar)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": BooksResponse(book),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "Not valid type",
		})
	}
}

func (h *handler) AddBookHandler(ctx *gin.Context) {
	var bookData BookRequestType

	err := ctx.ShouldBindJSON(&bookData)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errors := []string{}

			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errors = append(errors, errorMessage)
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errors,
			})
		} else {
			isNumber := helpers.IsNumeric(string(bookData.Price))
			fmt.Println("isNumber", isNumber)
			if !isNumber {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"errors": "Not valid type",
				})
			}
		}
		return
	}

	book, err := h.bookService.Create(bookData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *handler) UpdateBookHandler(ctx *gin.Context) {
	var bookData BookRequestUpdateType

	err := ctx.ShouldBindJSON(&bookData)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errors := []string{}

			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errors = append(errors, errorMessage)
			}

			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errors,
			})
		} else {
			isNumber := helpers.IsNumeric(string(bookData.Price))
			if !isNumber {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"errors": "Not valid type",
				})
			}
		}
		return
	}

	book, err := h.bookService.Update(bookData.ID, bookData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": BooksResponse(book),
	})
}

func (h *handler) DeleteBookHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	intVar, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "Not valid type",
		})
		return
	}

	if helpers.IsNumeric(id) {
		_, err := h.bookService.FindById(intVar)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
			return
		}

		intVar, err := strconv.Atoi(id)
		book, err := h.bookService.Delete(intVar)

		ctx.JSON(http.StatusOK, gin.H{
			"data":    book,
			"message": "delete success",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": "Not valid type",
		})
	}
}

func (h *handler) QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	price := ctx.Query("price")
	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (h *handler) TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *handler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":    "muhammad iqbal",
		"address": "jakarta barat kedoya",
	})
}
