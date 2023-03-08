package bakeryitems

import (
	"ServerDrivenUI/constants"
	"ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/models/uis/types/alignment"
	"ServerDrivenUI/models/uis/types/corner"
	"ServerDrivenUI/models/uis/types/font"
	"ServerDrivenUI/models/uis/types/imgscale"
	"ServerDrivenUI/models/uis/views"
	"fmt"
)

// BakeryTitle Title
func BakeryTitle(listChildren []interface{}) []interface{} {
	var rowChildren []interface{}
	var rowModifier = properties.Modifier{PaddingT: 10}

	var title = "Bakery Items"
	var ctaTitle = "Explore all"

	// Title
	var titleTextColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.26, Alpha: 1.0}
	var titleValue = properties.Value{TextColor: titleTextColor, Label: title, TextSize: 22, Weight: 700, FontStyle: font.FontStyleType.String(font.MontBold)}
	//var borderColor = properties.BackgroundColor{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}
	var titleTvModifier = properties.Modifier{BorderWidth: 1, PaddingL: 10}
	var titleTv = views.TextView(titleTvModifier, titleValue)
	rowChildren = append(rowChildren, titleTv)

	// CTA Title
	var ctaTextColor = properties.Color{Hue: 128, Saturation: 0.34, Lighting: 0.5, Alpha: 1.0}
	var ctaValue = properties.Value{TextColor: ctaTextColor, Label: ctaTitle, TextSize: 16, Weight: 600, FontStyle: font.FontStyleType.String(font.MontSemiBold)}
	var ctaTvModifier = properties.Modifier{PaddingR: 15}
	var ctaTv = views.TextView(ctaTvModifier, ctaValue)
	rowChildren = append(rowChildren, ctaTv)

	var row = views.Row(rowModifier, rowChildren)
	listChildren = append(listChildren, row)
	return listChildren
}

// BakeryHorizontalList Horizontal ListView
func BakeryHorizontalList(listChildren []interface{}) []interface{} {
	var horizontalChildren []interface{}
	for i := 0; i < 10; i++ {
		var columnChildren []interface{}

		// ImageView
		var imgViewValue = views.ImageValue{
			ImageUrl:   constants.BakeryItemsImagePath4x + "banana.jpg",
			Width:      70,
			Height:     70,
			ImageScale: imgscale.ImgScale.String(imgscale.FillBounds)}
		var imgView = views.ImageView(imgViewValue)
		columnChildren = append(columnChildren, imgView)

		// Title
		var titleStr = fmt.Sprintf("Groceries %d", i)
		var textColor = properties.Color{Hue: 240, Saturation: 0.76, Lighting: 0.5, Alpha: 1.0}
		var titleValue = properties.Value{TextColor: textColor, Label: titleStr, TextSize: 10, Weight: 400, Width: 100}
		var titleModifier = properties.Modifier{PaddingL: 16, PaddingT: 5, PaddingB: 5, TextAlignment: alignment.ColumnAlignment.String(alignment.Start)}
		var titleTv = views.TextView(titleModifier, titleValue)
		columnChildren = append(columnChildren, titleTv)

		// Column Child
		var columnColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
		var columnModifier = properties.Modifier{BackgroundColor: columnColor, PaddingT: 10, ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally)}
		var columnItem = views.Column(columnModifier, columnChildren)
		var columnValue = properties.Value{Width: 100}
		columnItem.Value = columnValue

		// Horizontal List Items List
		var columnParentChildren []interface{}
		columnParentChildren = append(columnParentChildren, columnItem)

		var parentModifier = properties.Modifier{
			PaddingL: 5, PaddingT: 10, PaddingR: 5, PaddingB: 10,
			Corner:          corner.CornerType.String(corner.ROUNDED_CORNER),
			Radius:          10,
			ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally),
		}

		var columnParent = views.Column(parentModifier, columnParentChildren)

		horizontalChildren = append(horizontalChildren, columnParent)
	}
	var horizontalModifier = properties.Modifier{}
	var horizontalListView = views.HorizontalList(horizontalModifier, horizontalChildren)
	listChildren = append(listChildren, horizontalListView)
	return listChildren
}
