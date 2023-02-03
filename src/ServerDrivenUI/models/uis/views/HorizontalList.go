package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type horizontalList struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Children []interface{}       `json:"children"`
	Value    HorizontalListValue `json:"value,omitempty"`
}

func HorizontalList(modifier properties.Modifier, children []interface{}) *horizontalList {
	horizontalListView := new(horizontalList)
	horizontalListView.Type = viewgrouptype.HORIZONTAL_LIST.String()
	horizontalListView.Modifier = modifier
	horizontalListView.Children = children
	horizontalListView.Value = HorizontalListValue{Width: -1}
	return horizontalListView
}

type HorizontalListValue struct {
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	PaddingL        int              `json:"paddingL"`
	PaddingT        int              `json:"paddingT"`
	PaddingR        int              `json:"paddingR"`
	PaddingB        int              `json:"paddingB"`
	BackgroundColor properties.Color `json:"backgroundColor"`
	Corner          string           `json:"corner"`
	IsEnable        bool             `json:"isEnable"`
	Radius          int              `json:"radius"`
}
