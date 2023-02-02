package categoryTitle

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/font"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
)

func CategoryTitle(title string, ctaTitle string, listChildren []interface{}) []interface{} {
	var rowChildren []interface{}
	var rowModifier = properties.Modifier{PaddingT: 10}

	// Title
	var titleValue = properties.Value{Label: title, TextSize: 22, Weight: 700, FontStyle: font.FontStyleType.String(font.MontBold)}
	var titleColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.26, Alpha: 1.0}
	//var borderColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}
	var titleTvModifier = properties.Modifier{Color: titleColor, BorderWidth: 1, PaddingL: 10}
	var titleTv = views.TextView(titleTvModifier, titleValue)
	rowChildren = append(rowChildren, titleTv)

	// CTA Title
	var ctaValue = properties.Value{Label: ctaTitle, TextSize: 16, Weight: 600, FontStyle: font.FontStyleType.String(font.MontSemiBold)}
	var ctaColor = properties.Color{Hue: 128, Saturation: 0.34, Lighting: 0.5, Alpha: 1.0}
	var ctaTvModifier = properties.Modifier{Color: ctaColor, PaddingR: 15}
	var ctaTv = views.TextView(ctaTvModifier, ctaValue)
	rowChildren = append(rowChildren, ctaTv)

	var row = views.Row(rowModifier, rowChildren)
	listChildren = append(listChildren, row)
	return listChildren
}
