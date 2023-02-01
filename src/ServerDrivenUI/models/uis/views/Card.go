package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type cardView struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Value    properties.Value    `json:"value,omitempty"`
}

func CardView(modifier properties.Modifier, value properties.Value) *cardView {
	tv := new(cardView)
	tv.Type = types.TEXT.String()
	tv.Modifier = modifier
	tv.Value = value
	return tv
}
