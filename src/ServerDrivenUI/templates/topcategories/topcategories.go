package topcategories

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

// TopCategoryTitle Category Title
func TopCategoryTitle(listChildren []interface{}) []interface{} {
	var rowChildren []interface{}
	var rowModifier = properties.Modifier{PaddingT: 10}

	var title = "Top Categories"
	var ctaTitle = "Explore all"

	// Title
	var textColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}
	var titleValue = properties.Value{TextColor: textColor, Label: title, TextSize: 22, Weight: 700, FontStyle: font.FontStyleType.String(font.MontBold)}
	var borderColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}

	// Modifier
	var backgroundColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
	var titleTvModifier = properties.Modifier{BorderColor: borderColor, BackgroundColor: backgroundColor, BorderWidth: 1, PaddingL: 10}
	var titleTv = views.TextView(titleTvModifier, titleValue)
	rowChildren = append(rowChildren, titleTv)

	// CTA Title
	var ctaValue = properties.Value{Label: ctaTitle, TextSize: 16, Weight: 600, FontStyle: font.FontStyleType.String(font.MontSemiBold)}
	var ctaColor = properties.Color{Hue: 128, Saturation: 0.34, Lighting: 0.5, Alpha: 1.0}
	var ctaTvModifier = properties.Modifier{BackgroundColor: ctaColor, PaddingR: 15}
	var ctaTv = views.TextView(ctaTvModifier, ctaValue)
	rowChildren = append(rowChildren, ctaTv)

	var row = views.Row(rowModifier, rowChildren)
	listChildren = append(listChildren, row)
	return listChildren
}

// HorizontalList Horizontal ListView
func HorizontalList(listChildren []interface{}) []interface{} {
	var horizontalChildren []interface{}
	for i := 0; i < 10; i++ {
		var columnChildren []interface{}

		// ImageView
		var imgViewValue = properties.Value{
			ImageUrl:   constants.FruitsImagePath4x + "banana.jpg",
			Width:      70,
			Height:     70,
			ImageScale: imgscale.ImgScale.String(imgscale.FillBounds)}
		var imgViewModifier = properties.Modifier{BackgroundColor: properties.ColorDummy()}
		var imgView = views.ImageView(imgViewModifier, imgViewValue)
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

// HorizontalGrid Horizontal GridView
func HorizontalGrid(listChildren []interface{}) []interface{} {
	var horizontalChildren []interface{}
	for i := 0; i < 10; i++ {
		var columnChildren []interface{}

		// ImageView
		var imgViewValue = properties.Value{
			ImageUrl:   constants.FruitsImagePath4x + "banana.jpg",
			Width:      70,
			Height:     70,
			ImageScale: imgscale.ImgScale.String(imgscale.FillBounds)}
		var imgViewModifier = properties.Modifier{BackgroundColor: properties.ColorDummy()}
		var imgView = views.ImageView(imgViewModifier, imgViewValue)
		columnChildren = append(columnChildren, imgView)

		// Title
		var titleStr = fmt.Sprintf("Groceries %d", i)
		var titleValue = properties.Value{Label: titleStr, TextSize: 10, Weight: 400, Width: 100}
		var titleColor = properties.Color{Hue: 240, Saturation: 0.76, Lighting: 0.5, Alpha: 1.0}
		var titleModifier = properties.Modifier{BackgroundColor: titleColor, PaddingL: 16, PaddingT: 5, PaddingB: 5, TextAlignment: alignment.ColumnAlignment.String(alignment.Start)}
		var titleTv = views.TextView(titleModifier, titleValue)
		columnChildren = append(columnChildren, titleTv)

		// Column Child
		var columnColor = properties.Color{Hue: 132, Saturation: 0.63, Lighting: 0.97, Alpha: 0.3}
		var columnModifier = properties.Modifier{BackgroundColor: columnColor, PaddingT: 10, ColumnAlignment: alignment.ColumnAlignment.String(alignment.CenterHorizontally)}
		var columnItem = views.Column(columnModifier, columnChildren)
		var columnValue = properties.Value{Width: 100}
		columnItem.Value = columnValue

		// Horizontal Grid Items List
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
	var horizontalGridView = views.HorizontalGrid(2, 250, horizontalModifier, horizontalChildren)
	listChildren = append(listChildren, horizontalGridView)
	return listChildren
}
