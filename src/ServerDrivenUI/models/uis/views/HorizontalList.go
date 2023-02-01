package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type horizontalList struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Children []interface{}       `json:"children"`
	Value    properties.Value    `json:"value,omitempty"`
}

func HorizontalList(modifier properties.Modifier, children []interface{}) *horizontalList {
	horizontalListView := new(horizontalList)
	horizontalListView.Type = types.HORIZONTAL_LIST.String()
	horizontalListView.Modifier = modifier
	horizontalListView.Children = children
	horizontalListView.Value = properties.Value{Width: -1}
	return horizontalListView
}
