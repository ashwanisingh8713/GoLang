package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type appBar struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Children []interface{}       `json:"children"`
}

func AppBar(modifier properties.Modifier, children []interface{}) *appBar {
	appBar := new(appBar)
	appBar.Type = types.APP_BAR.String()
	appBar.Modifier = modifier
	appBar.Children = children
	return appBar
}

func AppBarDummy() *appBar {
	var modifier = properties.ModifierDummy("App Bar Dummy")
	var children []interface{}
	var titleTextView = TitleTextView("App Bar Title")
	children = append(children, titleTextView)
	return AppBar(modifier, children)
}
