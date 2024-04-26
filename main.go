package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", rootHandler)
	r.GET("/helloworld", helloHandler)

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
