package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type verticalList struct {
	Type     string              `json:"type,omitempty" binding:"required"`
	Modifier properties.Modifier `json:"modifier,omitempty"`
	Children []interface{}       `json:"children,omitempty"`
	Value    properties.Value    `json:"value,omitempty"`
}

func VerticalList(modifier properties.Modifier, children []interface{}) *verticalList {
	verticalList := new(verticalList)
	verticalList.Type = viewgrouptype.VERTICAL_LIST.String()
	verticalList.Modifier = modifier
	verticalList.Children = children
	verticalList.Value = properties.Value{Width: -1, Height: -1}
	return verticalList
}
