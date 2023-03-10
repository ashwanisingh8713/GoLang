package main

import (
	"ServerDrivenUI/src/book/controllers"
	"ServerDrivenUI/src/book/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run("192.168.13.22:8080")
}
