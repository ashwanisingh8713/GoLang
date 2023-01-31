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

func HorizontalListDummy_Count1() *horizontalList {
	var modifier = properties.ModifierDummy("Horizontal List Modifier")
	var children []interface{}
	var subTitleTextView = SubTitleTextView("H-Row Item 1")
	children = append(children, subTitleTextView)
	return HorizontalList(modifier, children)
}

func HorizontalListDummy_Count2() *horizontalList {
	var modifier = properties.ModifierDummy("Horizontal List Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("H-Row Item 1")
	var subTitleTextView2 = SubTitleTextView("H-Row Item 2")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	return HorizontalList(modifier, children)
}

func HorizontalListDummy_Count3() *horizontalList {
	var modifier = properties.ModifierDummy("Horizontal List Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("H-Row Item 1")
	var subTitleTextView2 = SubTitleTextView("H-Row Item 2")
	var subTitleTextView3 = SubTitleTextView("H-Row Item 3")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, subTitleTextView3)
	return HorizontalList(modifier, children)
}

func HorizontalListDummy_Count2_ImageView() *horizontalList {
	var modifier = properties.ModifierDummy("Horizontal List Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("H-Row Item 1")
	var subTitleTextView2 = SubTitleTextView("H-Row Item 2")
	var imageView1 = ImageRect_300x200()
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, imageView1)
	return HorizontalList(modifier, children)
}
