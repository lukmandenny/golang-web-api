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
	// ++++++++++++++++++++++++++++++++++++++++++
	// CRUD OPERATION
	// ++++++++++++++++++++++++++++++++++++++++++
	// // CREATE
	// ++++++++++++++++++++++++++++++++++++++++++

	// book := book.Book{}
	// book.Title = "Kura-kura Ninja"
	// book.Description = "Kura-kura yang bisa bertarung seperti ninja, keren!"
	// book.Price = 150000
	// book.Rating = 5
	// book.Discount = 5

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// 	fmt.Println("Error creating book records")
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// }

	// ++++++++++++++++++++++++++++++++++++++++++
	// READ
	// ++++++++++++++++++++++++++++++++++++++++++

	// var book book.Book

	// err = db.Debug().First(&book, 2).Error
	// if err != nil {
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// 	fmt.Println("Error finding book records")
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// }

	// fmt.Println("Title:", book.Title)
	// fmt.Println("book object %v", book)

	// ++++++++++++++++++++++++++++++++++++++++++

	// var books []book.Book

	// err = db.Debug().Find(&books).Error
	// if err != nil {
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// 	fmt.Println("Error finding book records")
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	// ++++++++++++++++++++++++++++++++++++++++++

	// var books []book.Book

	// err = db.Debug().Where("title LIKE ?", "%a%").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// 	fmt.Println("Error finding book records")
	// 	fmt.Println("+++++++++++++++++++++++++++")
	// }

	// for _, b := range books {
	// 	fmt.Println("Title:", b.Title)
	// 	fmt.Println("book object %v", b)
	// }

	// ++++++++++++++++++++++++++++++++++++++++++
	// UPDATE
	// ++++++++++++++++++++++++++++++++++++++++++

	var book book.Book

	err = db.Debug().Where("id = ?", "1").First(&book).Error
	if err != nil {
		fmt.Println("+++++++++++++++++++++++++++")
		fmt.Println("Error finding book records")
		fmt.Println("+++++++++++++++++++++++++++")
	}

	book.Title = "Tahilalats x Tokyobike"
	err = db.Debug().Save(&book).Error
	if err != nil {
		fmt.Println("+++++++++++++++++++++++++++")
		fmt.Println("Error updating book records")
		fmt.Println("+++++++++++++++++++++++++++")
	}

	r := gin.Default()

	v1 := r.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/helloworld", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	r.Run(":8888")

}
