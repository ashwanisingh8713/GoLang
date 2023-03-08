package controller

import (
	"ServerDrivenUI/models/uis/views"
	"ServerDrivenUI/templates/appbar"
	"ServerDrivenUI/templates/banner"
	"ServerDrivenUI/templates/topcategories"
	"ServerDrivenUI/templates/topproduct"
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

	// Main Listing Children
	var verticalViewChildren []interface{}

	// Banner
	verticalViewChildren = banner.Banner(verticalViewChildren)
	// Top Category - Title
	verticalViewChildren = topcategories.TopCategoryTitle(verticalViewChildren)
	// Top Category  - Horizontal ListView
	verticalViewChildren = topcategories.HorizontalList(verticalViewChildren)
	// Top Category - Horizontal GridView
	verticalViewChildren = topcategories.HorizontalGrid(verticalViewChildren)

	// Top Product - Title
	verticalViewChildren = topproduct.TopProductTitle(verticalViewChildren)
	// Top Product  - Vertical GridView
	verticalViewChildren = topproduct.VerticalGrid(verticalViewChildren)

	// Bakery - Title
	//verticalViewChildren = bakeryitems.BakeryTitle(verticalViewChildren)
	// Bakery  - Horizontal ListView
	//verticalViewChildren = bakeryitems.BakeryHorizontalList(verticalViewChildren)

	// Below this, It should not be Touched
	var verticalListValue = views.VerticalListValue{Width: -1, Height: -1}
	var verticalListView = views.VerticalList(verticalListValue, verticalViewChildren)

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
