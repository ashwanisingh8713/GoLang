package controller

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context) {
	var scaffoldChildren []interface{}
	var appBarChildren []interface{}

	var verticalListModifier = properties.ModifierDummy("Vertical List Modifier")
	var verticalViewChildren []interface{}
	// Banner
	verticalViewChildren = Banner(verticalViewChildren)
	// Category Title
	verticalViewChildren = CategoryTitle("Top Categories", "Explore all", verticalViewChildren)
	// Horizontal ListView
	verticalViewChildren = HorizontalList_1(verticalViewChildren)

	var verticalListView = views.VerticalList(verticalListModifier, verticalViewChildren)

	// App Bar
	var appBar = views.AppBarDummy()
	appBarChildren = append(appBarChildren, appBar)

	// Scaffold
	scaffoldChildren = append(scaffoldChildren, verticalListView)
	var scaffold = views.Scaffold(scaffoldChildren, appBarChildren)

	// Outest Field
	var children []interface{}
	children = append(children, scaffold)
	//var m = make(map[string][]interface{})
	//m["data"] = children
	c.JSON(http.StatusOK, gin.H{"data": children})
}

func Banner(listChildren []interface{}) []interface{} {
	//var imgViewValue = properties.Value{ImageUrl: "https://cdn.pixabay.com/photo/2015/10/30/18/04/banner-1014539_1280.jpg",
	var imgViewValue = properties.Value{
		ImageUrl: "https://images.all-free-download.com/images/graphicwebp/rosa_red_beautiful_girl_219903.webp",
		Width:    500, Height: 100}
	var imgViewModifier = properties.Modifier{Color: properties.ColorDummy(), BorderColor: properties.BorderColorDummy()}
	var imgView = views.ImageView(imgViewModifier, imgViewValue)
	listChildren = append(listChildren, imgView)
	return listChildren
}

func CategoryTitle(title string, ctaTitle string, listChildren []interface{}) []interface{} {
	var rowChildren []interface{}
	var rowModifier = properties.ModifierDummy("Row Modifier")

	// Title
	var titleValue = properties.Value{Label: title, TextSize: 20, Weight: 700, FontStyle: types.FontStyleType.String(types.MontBold)}
	var titleColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.26, Alpha: 1.0}
	var titleTvmodifier = properties.Modifier{Color: titleColor}
	var titleTv = views.TextView(titleTvmodifier, titleValue)
	rowChildren = append(rowChildren, titleTv)

	// CTA Title
	var ctaValue = properties.Value{Label: ctaTitle, TextSize: 16, Weight: 600, FontStyle: types.FontStyleType.String(types.MontSemiBold)}
	var ctaColor = properties.Color{Hue: 128, Saturation: 0.34, Lighting: 0.5, Alpha: 1.0}
	var ctaTvmodifier = properties.Modifier{Color: ctaColor}
	var ctaTv = views.TextView(ctaTvmodifier, ctaValue)
	rowChildren = append(rowChildren, ctaTv)

	var row = views.Row(rowModifier, rowChildren)
	listChildren = append(listChildren, row)
	return listChildren
}

func HorizontalList_1(listChildren []interface{}) []interface{} {
	var horizontalChildren []interface{}
	for i := 0; i < 3; i++ {
		var columnChildren []interface{}
		var columnColor = properties.Color{Hue: 240, Saturation: 0.76, Lighting: 0.5, Alpha: 1.0}
		var columnModifier = properties.Modifier{Color: columnColor, PaddingL: 10, PaddingR: 10}

		// ImageView
		var imgViewValue = properties.Value{ImageUrl: "https://images.all-free-download.com/images/graphicwebp/rosa_red_beautiful_girl_219903.webp",
			Width: 100, Height: 100, ImageScale: "Crop"}
		var imgViewModifier = properties.Modifier{Color: properties.ColorDummy(), PaddingL: 10}
		var imgView = views.ImageView(imgViewModifier, imgViewValue)
		columnChildren = append(columnChildren, imgView)

		// Title
		var titleValue = properties.Value{Label: "Groceries", TextSize: 10, Weight: 400}
		var titleColor = properties.Color{Hue: 240, Saturation: 0.76, Lighting: 0.5, Alpha: 1.0}
		var titleModifier = properties.Modifier{Color: titleColor}
		var titleTv = views.TextView(titleModifier, titleValue)
		columnChildren = append(columnChildren, titleTv)

		// Column
		var columnItem = views.Column(columnModifier, columnChildren)
		// Horizontal List Items List

		horizontalChildren = append(horizontalChildren, columnItem)
	}
	var horizontalModifier = properties.Modifier{Color: properties.BorderColorDummy()}
	var horizontalListView = views.HorizontalList(horizontalModifier, horizontalChildren)
	listChildren = append(listChildren, horizontalListView)
	return listChildren
}
