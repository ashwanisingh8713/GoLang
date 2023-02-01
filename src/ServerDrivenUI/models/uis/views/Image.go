package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/corner"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewtype"
)

type imageView struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Value    properties.Value    `json:"value,omitempty"`
}

func ImageView(modifier properties.Modifier, value properties.Value) *imageView {
	imgView := new(imageView)
	imgView.Type = viewtype.IMAGE.String()
	imgView.Modifier = modifier
	imgView.Value = value
	return imgView
}

func ImageModifier() properties.Modifier {
	return properties.Modifier{Label: "ImageModifier", Color: properties.ColorDummy(), PaddingL: 5,
		PaddingT: 5, PaddingR: 5, PaddingB: 5, Corner: corner.CornerType.String(corner.CUT_CORNER)}
}

var image1 = "https://www.shutterstock.com/image-vector/group-abstract-diverse-people-friends-600w-2121864335.jpg"
