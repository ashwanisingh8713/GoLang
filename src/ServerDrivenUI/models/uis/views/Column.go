package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type column struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Children []interface{}       `json:"children"`
	Value    properties.Value    `json:"value,omitempty"`
}

func Column(modifier properties.Modifier, children []interface{}) *column {
	column := new(column)
	column.Type = viewgrouptype.COLUMN.String()
	column.Modifier = modifier
	column.Children = children
	column.Value = properties.Value{Width: -1}
	return column
}
