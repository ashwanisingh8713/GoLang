package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type row struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Children []interface{}       `json:"children"`
	Value    properties.Value    `json:"value,omitempty"`
}

func Row(modifier properties.Modifier, children []interface{}) *row {
	row := new(row)
	row.Type = viewgrouptype.ROW.String()
	row.Modifier = modifier
	row.Children = children
	row.Value = properties.Value{Width: -1}
	return row
}
