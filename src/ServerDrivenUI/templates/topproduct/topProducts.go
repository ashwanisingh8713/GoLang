package topproduct

import (
	"ServerDrivenUI/src/ServerDrivenUI/constants"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/alignment"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/corner"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/font"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/imgscale"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
	"fmt"
)

// TopProductTitle Category Title
func TopProductTitle(listChildren []interface{}) []interface{} {
	var rowChildren []interface{}
	var rowModifier = properties.Modifier{PaddingT: 10}

	var title = "Top Product"
	var ctaTitle = "Explore all"

	// Title
	var textColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}
	//var borderColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}
	//var backgroundColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
	var titleTvValue = views.TextViewValue{TextColor: textColor, Label: title, TextSize: 22,
		Weight: 700, FontStyle: font.FontStyleType.String(font.MontBold),
		BorderWidth: 1, PaddingL: 10}
	var titleTv = views.TextView(titleTvValue)
	rowChildren = append(rowChildren, titleTv)

	// CTA Title
	var ctaTextColor = properties.Color{Hue: 128, Saturation: 0.34, Lighting: 0.5, Alpha: 1.0}
	var ctaTvValue = views.TextViewValue{PaddingR: 15, TextColor: ctaTextColor,
		Label: ctaTitle, TextSize: 16, Weight: 600,
		FontStyle: font.FontStyleType.String(font.MontSemiBold)}
	var ctaTv = views.TextView(ctaTvValue)
	rowChildren = append(rowChildren, ctaTv)

	var row = views.Row(rowModifier, rowChildren)
	listChildren = append(listChildren, row)
	return listChildren
}

// VerticalGrid Vertical GridView
func VerticalGrid(listChildren []interface{}) []interface{} {
	var verticalGridChildren []interface{}
	for i := 0; i < 10; i++ {
		var columnChildren []interface{}

		// ImageView
		var imgViewValue = views.ImageValue{
			ImageUrl:   constants.FruitsImagePath4x + "banana.jpg",
			Width:      70,
			Height:     70,
			ImageScale: imgscale.ImgScale.String(imgscale.FillBounds)}
		var imgView = views.ImageView(imgViewValue)
		columnChildren = append(columnChildren, imgView)

		// Title
		var titleStr = fmt.Sprintf("Groceries %d", i)
		var textColor = properties.Color{Hue: 240, Saturation: 0.76, Lighting: 0.5, Alpha: 1.0}
		var titleValue = views.TextViewValue{TextColor: textColor, Label: titleStr, TextSize: 10, Weight: 400, Width: 100, PaddingL: 16, PaddingT: 5, PaddingB: 5, TextAlignment: alignment.ColumnAlignment.String(alignment.Start)}
		var titleTv = views.TextView(titleValue)
		columnChildren = append(columnChildren, titleTv)

		// Column Child
		var columnColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
		var columnModifier = properties.Modifier{BackgroundColor: columnColor, PaddingT: 10, ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally)}
		var columnItem = views.Column(columnModifier, columnChildren)
		//var columnValue = properties.Value{Width: 100}
		//columnItem.Value = columnValue

		// Horizontal Grid Items List
		var columnParentChildren []interface{}
		columnParentChildren = append(columnParentChildren, columnItem)

		var parentModifier = properties.Modifier{
			PaddingL: 55, PaddingT: 10, PaddingR: 15, PaddingB: 10,
			Corner:          corner.CornerType.String(corner.ROUNDED_CORNER),
			Radius:          10,
			ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally),
		}

		var columnParent = views.Column(parentModifier, columnParentChildren)

		verticalGridChildren = append(verticalGridChildren, columnParent)
	}
	var verticalGridValue = views.VerticalGridValue{GridColumn: 4, GridHeight: 350,
		VerticalArrangement: 10, HorizontalArrangement: 10, PaddingT: 10}
	var verticalGridView = views.VerticalGrid(verticalGridValue, verticalGridChildren)
	listChildren = append(listChildren, verticalGridView)
	return listChildren
}
