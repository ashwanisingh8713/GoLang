package appbar

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/font"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/views"
)

func AppBar() []interface{} {
	var listChildren []interface{}
	var modifier = properties.ModifierDummy("App Bar Dummy")
	var children []interface{}
	var titleValue = properties.Value{Label: "App Bar Title ", TextSize: 25, Weight: 700, FontStyle: font.FontStyleType.String(font.MontBold)}
	var titleColor = properties.Color{Hue: 0, Saturation: 0.0, Lighting: 0.26, Alpha: 1.0}
	//var borderColor = properties.BackgroundColor{Hue: 0, Saturation: 0.0, Lighting: 0.0, Alpha: 1.0}
	var titleTvModifier = properties.Modifier{BackgroundColor: titleColor}
	var titleTv = views.TextView(titleTvModifier, titleValue)
	children = append(children, titleTv)
	var appBar = views.AppBar(modifier, children)
	listChildren = append(listChildren, appBar)
	return listChildren
}
