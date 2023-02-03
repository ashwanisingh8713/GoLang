package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewtype"
)

type textView struct {
	Type  string        `json:"type" binding:"required"`
	Value TextViewValue `json:"textViewValue,omitempty"`
}

func TextView(textViewValue TextViewValue) *textView {
	tv := new(textView)
	tv.Type = viewtype.TEXT.String()
	tv.Value = textViewValue
	return tv
}

type TextViewValue struct {
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	FontStyle       string           `json:"fontStyle,omitempty"`
	Weight          int              `json:"weight,omitempty"`
	Label           string           `json:"label,omitempty"`
	TextSize        int              `json:"textSize,omitempty"`
	TextColor       properties.Color `json:"textColor,omitempty"`
	BackgroundColor properties.Color `json:"backgroundColor,omitempty"`
	BorderColor     properties.Color `json:"borderColor,omitempty"`
	BorderWidth     int              `json:"borderWidth,omitempty"`
	PaddingL        int              `json:"paddingL,omitempty"`
	PaddingT        int              `json:"paddingT,omitempty"`
	PaddingR        int              `json:"paddingR,omitempty"`
	PaddingB        int              `json:"paddingB,omitempty"`
	Corner          string           `json:"cornerType,omitempty"`
	IsEnable        bool             `json:"isEnable,omitempty"`
	Radius          int              `json:"radius,omitempty"`
	TextAlignment   string           `json:"textAlignment,omitempty"`
}
