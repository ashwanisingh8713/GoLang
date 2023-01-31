package main

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.GET("/ui", controller.Response)
	route.Run("192.168.13.87:8080")
	//route.Run()

}
