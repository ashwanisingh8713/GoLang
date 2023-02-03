package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type viewPager struct {
	Type     string         `json:"type" binding:"required"`
	Children []interface{}  `json:"children"`
	Value    ViewPagerValue `json:"value,omitempty"`
}

func ViewPager(value ViewPagerValue, children []interface{}) *viewPager {
	viewPager := new(viewPager)
	viewPager.Type = viewgrouptype.PAGER.String()
	viewPager.Children = children
	viewPager.Value = value
	return viewPager
}

type ViewPagerValue struct {
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	BackgroundColor properties.Color `json:"backgroundColor,omitempty"`
	BorderColor     properties.Color `json:"borderColor,omitempty"`
	BorderWidth     int              `json:"borderWidth,omitempty"`
	PaddingL        int              `json:"paddingL,omitempty"`
	PaddingT        int              `json:"paddingT,omitempty"`
	PaddingR        int              `json:"paddingR,omitempty"`
	PaddingB        int              `json:"paddingB,omitempty"`
	Corner          string           `json:"corner,omitempty"`
	IsEnable        bool             `json:"isEnable,omitempty"`
}
