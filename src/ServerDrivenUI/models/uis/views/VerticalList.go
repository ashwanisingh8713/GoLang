package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type verticalList struct {
	Type     string              `json:"type,omitempty" binding:"required"`
	Modifier properties.Modifier `json:"modifier,omitempty"`
	Children []interface{}       `json:"children,omitempty"`
	Value    properties.Value    `json:"value,omitempty"`
}

func VerticalList(modifier properties.Modifier, children []interface{}) *verticalList {
	verticalList := new(verticalList)
	verticalList.Type = types.VERTICAL_LIST.String()
	verticalList.Modifier = modifier
	verticalList.Children = children
	verticalList.Value = properties.Value{Width: -1, Height: -1}
	return verticalList
}

func VerticalListEmptyChild(modifier properties.Modifier) *verticalList {
	verticalList := new(verticalList)
	verticalList.Type = types.VERTICAL_LIST.String()
	verticalList.Modifier = modifier
	return verticalList
}

func VerticalListDummy_Count1() *verticalList {
	var modifier = properties.ModifierDummy("Vertical List Modifier")
	var children []interface{}
	var subTitleTextView = SubTitleTextView("V-Row Item 1")
	children = append(children, subTitleTextView)
	return VerticalList(modifier, children)
}

func VerticalListDummy_Count2() *verticalList {
	var modifier = properties.ModifierDummy("Vertical List Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("V-Row Item 1")
	var subTitleTextView2 = SubTitleTextView("V-Row Item 2")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	return VerticalList(modifier, children)
}

func VerticalListDummy_Count3() *verticalList {
	var modifier = properties.ModifierDummy("Vertical List Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("V-Row Item 1")
	var subTitleTextView2 = SubTitleTextView("V-Row Item 2")
	var subTitleTextView3 = SubTitleTextView("V-Row Item 3")
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, subTitleTextView3)
	return VerticalList(modifier, children)
}

func VerticalListDummy_Count2_Verticle_Count2() *verticalList {
	var modifier = properties.ModifierDummy("Vertical List Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("V-Row Item 1")
	var subTitleTextView2 = SubTitleTextView("V-Row Item 2")
	var horizontal = HorizontalListDummy_Count2_ImageView()
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, horizontal)
	return VerticalList(modifier, children)
}

func VerticalListDummy_Count2_Horizontal_Count2() *verticalList {
	var modifier = properties.ModifierDummy("Vertical List Modifier")
	var children []interface{}
	var subTitleTextView1 = SubTitleTextView("V-Row Item 1")
	var subTitleTextView2 = SubTitleTextView("V-Row Item 2")
	var horizontal = HorizontalListDummy_Count2_ImageView()
	children = append(children, subTitleTextView1)
	children = append(children, subTitleTextView2)
	children = append(children, horizontal)
	return VerticalList(modifier, children)
}
