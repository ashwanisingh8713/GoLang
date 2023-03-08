package views

import (
	"ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/models/uis/types/viewtype"
)

type editText struct {
	Type  string        `json:"type" binding:"required"`
	Value EditTextValue `json:"value,omitempty"`
}

func EditText(value EditTextValue) *editText {
	editText := new(editText)
	editText.Type = viewtype.WEBVIEW.String()
	editText.Value = value
	return editText
}

type EditTextValue struct {
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	FontStyle       string           `json:"fontStyle,omitempty"`
	Weight          int              `json:"weight,omitempty"`
	Label           string           `json:"label,omitempty"`
	TextSize        int              `json:"textSize,omitempty"`
	TextColor       properties.Color `json:"textColor,omitempty"`
	BackgroundColor properties.Color `json:"backgroundColor"`
	BorderColor     properties.Color `json:"borderColor"`
	BorderWidth     int              `json:"borderWidth"`
	PaddingL        int              `json:"paddingL"`
	PaddingT        int              `json:"paddingT"`
	PaddingR        int              `json:"paddingR"`
	PaddingB        int              `json:"paddingB"`
	Corner          string           `json:"corner"`
	IsEnable        bool             `json:"isEnable"`
	Radius          int              `json:"radius"`
	TextAlignment   string           `json:"textAlignment"`
}
