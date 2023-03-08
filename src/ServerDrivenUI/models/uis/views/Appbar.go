package views

import (
	"ServerDrivenUI/models/uis/types/viewgrouptype"
)

type appBar struct {
	Type     string        `json:"type" binding:"required"`
	Children []interface{} `json:"children"`
}

func AppBar(children []interface{}) *appBar {
	appBar := new(appBar)
	appBar.Type = viewgrouptype.APP_BAR.String()
	appBar.Children = children
	return appBar
}
