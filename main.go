package main

import (
	"fmt"
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

	findall, _ := bookService.FindAll()

	for _, book := range findall {
		fmt.Println("Title:", book.Title)
	}

	findbyid, _ := bookService.FindByID(2)

	fmt.Println("Title:", findbyid.Title)

	bookRequest := book.BookRequest{
		Title: "Spiderman",
		Price: 120000,
	}

	newBook, _ := bookService.Create(bookRequest)

	fmt.Println("Title:", newBook.Title)

	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/helloworld", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	r.Run(":8888")

}
