package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/corner"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewtype"
)

type imageView struct {
	Type string `json:"type" binding:"required"`
	//Modifier properties.Modifier `json:"modifier"`
	Value ImageValue `json:"value,omitempty"`
}

func ImageView(value ImageValue) *imageView {
	imgView := new(imageView)
	imgView.Type = viewtype.IMAGE.String()
	imgView.Value = value
	return imgView
}

type ImageValue struct {
	ImageUrl        string           `json:"imageUrl,omitempty"`
	ImageScale      string           `json:"imageScale,omitempty"`
	PaddingL        int              `json:"paddingL,omitempty"`
	PaddingT        int              `json:"paddingT,omitempty"`
	PaddingR        int              `json:"paddingR,omitempty"`
	PaddingB        int              `json:"paddingB,omitempty"`
	Corner          string           `json:"corner,omitempty"`
	BackgroundColor properties.Color `json:"backgroundColor,omitempty"`
	Radius          int              `json:"radius,omitempty"`
	BorderColor     properties.Color `json:"borderColor,omitempty"`
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	BorderWidth     int              `json:"borderWidth"`
}

func ImageModifier() properties.Modifier {
	return properties.Modifier{Label: "ImageModifier", BackgroundColor: properties.ColorDummy(), PaddingL: 5,
		PaddingT: 5, PaddingR: 5, PaddingB: 5, Corner: corner.CornerType.String(corner.CUT_CORNER)}
}

var image1 = "https://www.shutterstock.com/image-vector/group-abstract-diverse-people-friends-600w-2121864335.jpg"
