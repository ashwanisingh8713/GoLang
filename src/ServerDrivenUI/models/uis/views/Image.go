package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types"
)

type imageView struct {
	Type     string              `json:"type" binding:"required"`
	Modifier properties.Modifier `json:"modifier"`
	Value    properties.Value    `json:"value,omitempty"`
}

func ImageView(modifier properties.Modifier, value properties.Value) *imageView {
	imgView := new(imageView)
	imgView.Type = types.IMAGE.String()
	imgView.Modifier = modifier
	imgView.Value = value
	return imgView
}

func ImageModifier() properties.Modifier {
	return properties.Modifier{Label: "ImageModifier", Color: properties.ColorDummy(), PaddingL: 5,
		PaddingT: 5, PaddingR: 5, PaddingB: 5, Corner: types.CornerType.String(types.CUT_CORNER)}
}

var image1 = "https://www.shutterstock.com/image-vector/group-abstract-diverse-people-friends-600w-2121864335.jpg"
