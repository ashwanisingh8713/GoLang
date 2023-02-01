package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type textView struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Value    properties.Value    `json:"value,omitempty"`
}

func TextView(modifier properties.Modifier, value properties.Value) *textView {
	tv := new(textView)
	tv.Type = types.TEXT.String()
	tv.Modifier = modifier
	tv.Value = value
	return tv
}
