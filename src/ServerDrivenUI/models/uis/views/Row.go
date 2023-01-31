package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type row struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Children []interface{}       `json:"children"`
	Value    properties.Value    `json:"value,omitempty"`
}

func Row(modifier properties.Modifier, children []interface{}) *row {
	row := new(row)
	row.Type = types.ROW.String()
	row.Modifier = modifier
	row.Children = children
	row.Value = properties.Value{Width: -1}
	return row
}

func RowDummy_Count1() *row {
	var modifier = properties.ModifierDummy("Row Modifier")
	var children []interface{}
	var subTitleTextView = SubTitleTextView("Row Item 1")
	children = append(children, subTitleTextView)
	return Row(modifier, children)
}

func RowDummy_Count2() *row {
	var modifier = properties.ModifierDummy("Row Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("Row Item 1")
	var subTitleTextView2 = SubTitleTextView("Row Item 2")
	var subTitleTextView3 = SubTitleTextView("Row Item 3")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, subTitleTextView3)
	return Row(modifier, children)
}

func RowDummy_Count3() *row {
	var modifier = properties.ModifierDummy("Row Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("Row Item 1")
	var subTitleTextView2 = SubTitleTextView("Row Item 2")
	var subTitleTextView3 = SubTitleTextView("Row Item 3")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, subTitleTextView3)
	return Row(modifier, children)
}
