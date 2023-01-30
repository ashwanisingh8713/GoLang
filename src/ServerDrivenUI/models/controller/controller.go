package controller

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context) {
	var scaffoldChildren []interface{}
	var appBarChildren []interface{}

	var verticalListView = views.VerticalListDummy_Count2_Horizontal_Count2()
	scaffoldChildren = append(scaffoldChildren, verticalListView)

	var appBar = views.AppBarDummy()
	appBarChildren = append(appBarChildren, appBar)
	var scaffold = views.Scaffold(scaffoldChildren, appBarChildren)

	// Outest Field
	var children []interface{}
	children = append(children, scaffold)
	//var m = make(map[string][]interface{})
	//m["data"] = children
	c.JSON(http.StatusOK, gin.H{"data": children})
}
