package handler

import (
	"fmt"
	"net/http"

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

// func (h *bookHandler) RootHandler(ctx *gin.Context) {
// 	// Menggunakan gin.H untuk mengirim response JSON
// 	data := gin.H{
// 		"message": "Route URL",
// 	}
// 	ctx.JSON(http.StatusOK, data)
// }

// func (h *bookHandler) HelloHandler(ctx *gin.Context) {
// 	// Menggunakan gin.H untuk mengirim response JSON
// 	data := gin.H{
// 		"message": "Hello World! My Name is",
// 		"name":    "Denny",
// 	}
// 	ctx.JSON(http.StatusOK, data)
// }

// func (h *bookHandler) BooksHandler(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	title := ctx.Param("title")

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"title": title,
// 	})
// }

// func (h *bookHandler) QueryHandler(ctx *gin.Context) {
// 	// Menggunakan ctx.Query("title") untuk menangkap nilai dari parameter query "title"
// 	title := ctx.Query("title")
// 	harga := ctx.Query("harga")

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"title": title,
// 		"harga": harga,
// 	})
// }

func (h *bookHandler) GetBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *bookHandler) PostBooksHandler(ctx *gin.Context) {
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
