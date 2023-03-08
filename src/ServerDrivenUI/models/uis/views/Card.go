package views

import (
	"ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/models/uis/types/viewgrouptype"
)

type cardView struct {
	Type  string    `json:"type" binding:"required"`
	Value CardValue `json:"value,omitempty"`
}

func CardView(value CardValue) *cardView {
	card := new(cardView)
	card.Type = viewgrouptype.CARD.String()
	card.Value = value
	return card
}

type CardValue struct {
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
