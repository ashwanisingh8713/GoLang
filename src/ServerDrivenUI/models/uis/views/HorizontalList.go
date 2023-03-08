package views

import (
	"ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/models/uis/types/viewgrouptype"
)

type horizontalList struct {
	Type     string              `json:"type" binding:"required"`
	Children []interface{}       `json:"children"`
	Value    HorizontalListValue `json:"horizontalListValue,omitempty"`
}

func HorizontalList(horizontalListValue HorizontalListValue, children []interface{}) *horizontalList {
	horizontalListView := new(horizontalList)
	horizontalListView.Type = viewgrouptype.HORIZONTAL_LIST.String()
	horizontalListView.Children = children
	horizontalListView.Value = horizontalListValue
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
	Corner          string           `json:"cornerType"`
	IsEnable        bool             `json:"isEnable"`
	Radius          int              `json:"radius"`
}
