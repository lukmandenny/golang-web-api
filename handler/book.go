package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"golang-web-api/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func (h *bookHandler) CreateBook(ctx *gin.Context) {
	// title, price
	var bookRequest book.BookRequest

	err := ctx.ShouldBindJSON(&bookRequest)

	if err != nil {
		errorMessages := []string{}
		if vErr, ok := err.(validator.ValidationErrors); ok {
			// Lakukan iterasi untuk setiap error
			for _, e := range vErr {
				// Proses setiap error
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		} else {
			errorMessages = append(errorMessages, err.Error())
			// Jika err bukan validator.ValidationErrors, kirimkan pesan kesalahan umum
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}

	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func (h *bookHandler) UpdateBook(ctx *gin.Context) {

	// title, price
	var bookRequestUpdate book.BookRequestUpdate

	err := ctx.ShouldBindJSON(&bookRequestUpdate)

	if err != nil {
		errorMessages := []string{}
		if vErr, ok := err.(validator.ValidationErrors); ok {
			// Lakukan iterasi untuk setiap error
			for _, e := range vErr {
				// Proses setiap error
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		} else {
			errorMessages = append(errorMessages, err.Error())
			// Jika err bukan validator.ValidationErrors, kirimkan pesan kesalahan umum
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}

	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookRequestUpdate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func (h *bookHandler) DeleteBook(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})

}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		ID:          b.ID,
		Discount:    b.Discount,
	}
}
