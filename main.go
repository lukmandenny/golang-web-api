package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)
	r.GET("/helloworld", helloHandler)
	r.GET("/books/:id/:title", booksHandler)
	r.GET("/query", queryHandler)
	r.POST("/books", postBooksHandler)

	r.Run(":8888")

}

func rootHandler(ctx *gin.Context) {
	// Menggunakan gin.H untuk mengirim response JSON
	data := gin.H{
		"message": "Route URL",
	}
	ctx.JSON(http.StatusOK, data)
}

func helloHandler(ctx *gin.Context) {
	// Menggunakan gin.H untuk mengirim response JSON
	data := gin.H{
		"message": "Hello World! My Name is",
		"name":    "Denny",
	}
	ctx.JSON(http.StatusOK, data)
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")

	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(ctx *gin.Context) {
	// Menggunakan ctx.Query("title") untuk menangkap nilai dari parameter query "title"
	title := ctx.Query("title")
	harga := ctx.Query("harga")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"harga": harga,
	})
}

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Harga    int    `json:"harga" binding:"required,number"`
	SubTitle string `json:"sub_title"` //directive
}

func postBooksHandler(ctx *gin.Context) {
	// title, price
	var bookInput BookInput

	err := ctx.ShouldBindJSON(&bookInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		log.Println("Error:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"harga":     bookInput.Harga,
		"sub_title": bookInput.SubTitle,
	})

}
