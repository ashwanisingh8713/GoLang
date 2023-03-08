package topcategories

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

// TopCategoryTitle Category Title
func TopCategoryTitle(listChildren []interface{}) []interface{} {
	var rowChildren []interface{}
	var rowViewValue = views.RowViewValue{PaddingT: 10, Width: -1}

	var title = "Top Categories"
	var ctaTitle = "Explore all"

	// Title
	var textColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}
	var borderColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}

	// Modifier
	var backgroundColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
	var titleTvValue = views.TextViewValue{BorderColor: borderColor, BackgroundColor: backgroundColor,
		BorderWidth: 1, PaddingL: 10, TextColor: textColor, Label: title, TextSize: 22, Weight: 700, FontStyle: font.FontStyleType.String(font.MontBold)}
	var titleTv = views.TextView(titleTvValue)
	rowChildren = append(rowChildren, titleTv)

	// CTA Title
	var ctaTextColor = properties.Color{Hue: 128, Saturation: 0.34, Lighting: 0.5, Alpha: 1.0}
	var ctaTvValue = views.TextViewValue{PaddingR: 15, TextColor: ctaTextColor, Label: ctaTitle, TextSize: 16, Weight: 600, FontStyle: font.FontStyleType.String(font.MontSemiBold)}
	var ctaTv = views.TextView(ctaTvValue)
	rowChildren = append(rowChildren, ctaTv)

	var row = views.Row(rowViewValue, rowChildren)
	listChildren = append(listChildren, row)
	return listChildren
}

// HorizontalList Horizontal ListView
func HorizontalList(listChildren []interface{}) []interface{} {
	var horizontalChildren []interface{}
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
		var titleValue = views.TextViewValue{PaddingL: 16, PaddingT: 5, PaddingB: 5,
			TextAlignment: alignment.ColumnAlignment.String(alignment.Start), TextColor: textColor, Label: titleStr, TextSize: 10, Weight: 400, Width: 100}
		var titleTv = views.TextView(titleValue)
		columnChildren = append(columnChildren, titleTv)

		// Column Child
		var columnColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
		var columnModifier = views.ColumnViewValue{Width: 100, BackgroundColor: columnColor, PaddingT: 10, ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally)}
		var columnItem = views.Column(columnModifier, columnChildren)

		// Horizontal List Items List
		var columnParentChildren []interface{}
		columnParentChildren = append(columnParentChildren, columnItem)

		var parentModifier = views.ColumnViewValue{
			PaddingL: 5, PaddingT: 10, PaddingR: 5, PaddingB: 10,
			Corner:          corner.CornerType.String(corner.ROUNDED_CORNER),
			Radius:          10,
			ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally),
		}

		var columnParent = views.Column(parentModifier, columnParentChildren)

		horizontalChildren = append(horizontalChildren, columnParent)
	}
	var horizontalListValue = views.HorizontalListValue{}
	var horizontalListView = views.HorizontalList(horizontalListValue, horizontalChildren)
	listChildren = append(listChildren, horizontalListView)
	return listChildren
}

// HorizontalGrid Horizontal GridView
func HorizontalGrid(listChildren []interface{}) []interface{} {
	var horizontalChildren []interface{}
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
		var titleValue = views.TextViewValue{PaddingL: 16, PaddingT: 5, PaddingB: 5,
			TextAlignment: alignment.ColumnAlignment.String(alignment.Start), TextColor: textColor, Label: titleStr, TextSize: 10, Weight: 400, Width: 100}
		var titleTv = views.TextView(titleValue)
		columnChildren = append(columnChildren, titleTv)

		// Column Child
		//var columnBgColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
		var columnModifier = views.ColumnViewValue{ /*BackgroundColor: columnBgColor,*/ PaddingT: 10, ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally)}
		var columnItem = views.Column(columnModifier, columnChildren)

		// Horizontal Grid Items List
		var columnParentChildren []interface{}
		columnParentChildren = append(columnParentChildren, columnItem)

		var parentModifier = views.ColumnViewValue{
			PaddingL: 5, PaddingT: 10, PaddingR: 5, PaddingB: 10,
			Corner:          corner.CornerType.String(corner.ROUNDED_CORNER),
			Radius:          10,
			ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally),
		}

		var columnParent = views.Column(parentModifier, columnParentChildren)

		horizontalChildren = append(horizontalChildren, columnParent)
	}
	var horizontalGridValue = views.HorizontalGridValue{GridHeight: 220, GridColumn: 2,
		VerticalArrangement: 5, HorizontalArrangement: 0,
		BackgroundColor: properties.Color{
			Hue: 300, Saturation: .76, Lighting: 0.72, Alpha: 0.3,
		}}
	var horizontalGridView = views.HorizontalGrid(horizontalGridValue, horizontalChildren)
	listChildren = append(listChildren, horizontalGridView)
	return listChildren
}
