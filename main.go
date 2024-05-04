package main

import (
	"golang-web-api/book"
	"golang-web-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(localhost:3307)/golang_web_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	r := gin.Default()

	v1 := r.Group("/v1")

	// v1.GET("/", bookHandler.RootHandler)
	// v1.GET("/helloworld", bookHandler.HelloHandler)
	// v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	// v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.GET("/books", bookHandler.GetBooks)

	r.Run(":8888")

}
