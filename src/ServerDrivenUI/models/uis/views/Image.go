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

func ImageBanner() *imageView {
	var modifier = ImageModifier()
	var value = properties.Value{Width: -1, Height: -1}
	var imageView = ImageView(modifier, value)
	return imageView
}

func ImageThumb() *imageView {
	var modifier = ImageModifier()
	var value = properties.Value{Width: 100, Height: 100}
	var imageView = ImageView(modifier, value)
	return imageView
}

func ImageRect_100x200() *imageView {
	var modifier = ImageModifier()
	var value = properties.Value{Width: 100, Height: 200}
	var imageView = ImageView(modifier, value)
	return imageView
}

func ImageRect_200x200() *imageView {
	var modifier = ImageModifier()
	var value = properties.Value{Width: 200, Height: 200}
	var imageView = ImageView(modifier, value)
	return imageView
}

func ImageRect_300x200() *imageView {
	var modifier = ImageModifier()
	var value = properties.Value{Width: 300, Height: 200}
	var imageView = ImageView(modifier, value)
	return imageView
}
