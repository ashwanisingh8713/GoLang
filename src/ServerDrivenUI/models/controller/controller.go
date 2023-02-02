package controller

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
	"ServerDrivenUI/src/ServerDrivenUI/templates/appbar"
	"ServerDrivenUI/src/ServerDrivenUI/templates/banner"
	"ServerDrivenUI/src/ServerDrivenUI/templates/topcategories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func page2(c *gin.Context) {

	// Main Listing Children
	var verticalViewChildren []interface{}
	// Horizontal GridView
	verticalViewChildren = topcategories.HorizontalGrid(verticalViewChildren)

	c.JSON(http.StatusOK, gin.H{"data": verticalViewChildren})
}

func Response(c *gin.Context) {

	var verticalListModifier = properties.ModifierDummy("Vertical List Modifier")

	// Main Listing Children
	var verticalViewChildren []interface{}

	// Banner
	verticalViewChildren = banner.Banner(verticalViewChildren)
	// Top Category - Title
	verticalViewChildren = topcategories.CategoryTitle("Top Categories", "Explore all", verticalViewChildren)
	// Top Category  - Horizontal ListView
	verticalViewChildren = topcategories.HorizontalList(verticalViewChildren)
	// Top Category - Horizontal GridView
	verticalViewChildren = topcategories.HorizontalGrid(verticalViewChildren)

	var verticalListView = views.VerticalList(verticalListModifier, verticalViewChildren)

	// App Bar
	var appBarChildren = appbar.AppBar()

	// Scaffold
	var scaffoldChildren []interface{}
	scaffoldChildren = append(scaffoldChildren, verticalListView)
	var scaffold = views.Scaffold(scaffoldChildren, appBarChildren)

	// Outest Field
	var children []interface{}
	children = append(children, scaffold)
	c.JSON(http.StatusOK, gin.H{"data": children})

}
