package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/helloworld", func(ctx *gin.Context) {
		// Menggunakan gin.H untuk mengirim response JSON
		data := gin.H{
			"message": "Hello World! My Name is",
			"name":    "Denny",
		}
		ctx.JSON(http.StatusOK, data)
	})
	r.Run(":8888")

}
