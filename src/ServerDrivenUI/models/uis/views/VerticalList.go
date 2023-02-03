package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type verticalList struct {
	Type     string            `json:"type,omitempty" binding:"required"`
	Children []interface{}     `json:"children,omitempty"`
	Value    VerticalListValue `json:"verticalListValue,omitempty"`
}

func VerticalList(verticalListValue VerticalListValue, children []interface{}) *verticalList {
	verticalList := new(verticalList)
	verticalList.Type = viewgrouptype.VERTICAL_LIST.String()
	verticalList.Children = children
	verticalList.Value = verticalListValue
	return verticalList
}

type VerticalListValue struct {
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	BackgroundColor properties.Color `json:"backgroundColor,omitempty"`
	PaddingL        int              `json:"paddingL,omitempty"`
	PaddingT        int              `json:"paddingT,omitempty"`
	PaddingR        int              `json:"paddingR,omitempty"`
	PaddingB        int              `json:"paddingB,omitempty"`
	IsEnable        bool             `json:"isEnable,omitempty"`
}
