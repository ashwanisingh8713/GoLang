package main

import (
	"ServerDrivenUI/src/ServerDrivenUI/constants"
	"ServerDrivenUI/src/ServerDrivenUI/models/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.GET("/ui", controller.Response)
	route.Static(constants.Image_Relative_Path_4x, constants.DrawableFolderPath_4x)
	route.Static(constants.Image_Relative_Path_3x, constants.DrawableFolderPath_3x)
	route.Static(constants.Image_Relative_Path_2x, constants.DrawableFolderPath_2x)
	err := route.Run(constants.IP_Address)
	if err != nil {
		return
	}

}
