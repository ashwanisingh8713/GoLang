package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type column struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Children []interface{}       `json:"children"`
}

func Column(modifier properties.Modifier, children []interface{}) *column {
	row := new(column)
	row.Type = types.ROW.String()
	row.Modifier = modifier
	row.Children = children
	return row
}

func ColumnDummy_Count1() *column {
	var modifier = properties.ModifierDummy("Row Modifier")
	var children []interface{}
	var subTitleTextView = SubTitleTextView("Row Item 1")
	children = append(children, subTitleTextView)
	return Column(modifier, children)
}

func ColumnDummy_Count2() *column {
	var modifier = properties.ModifierDummy("Row Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("Row Item 1")
	var subTitleTextView2 = SubTitleTextView("Row Item 2")
	var subTitleTextView3 = SubTitleTextView("Row Item 3")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, subTitleTextView3)
	return Column(modifier, children)
}

func ColumnDummy_Count3() *column {
	var modifier = properties.ModifierDummy("Row Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("Row Item 1")
	var subTitleTextView2 = SubTitleTextView("Row Item 2")
	var subTitleTextView3 = SubTitleTextView("Row Item 3")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, subTitleTextView3)
	return Column(modifier, children)
}
