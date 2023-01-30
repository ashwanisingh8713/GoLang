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

func TextViewValue(label string, size int, weight int) properties.Value {
	return properties.Value{Label: label, TextSize: size, Weight: weight}
}

func NormalTextValue(label string) properties.Value {
	return TextViewValue(label, 14, 300)
}

func NormalTextView(label string) *textView {
	var modifier = properties.ModifierDummy(label)
	return TextView(modifier, NormalTextValue(label))
}

func TitleTextViewValue(label string) properties.Value {
	return TextViewValue(label, 25, 600)
}
func TitleTextView(label string) *textView {
	var modifier = properties.ModifierDummy(label)
	return TextView(modifier, TitleTextViewValue(label))
}

func SubTitleTextViewValue(label string) properties.Value {
	return TextViewValue(label, 18, 450)
}
func SubTitleTextView(label string) *textView {
	var modifier = properties.ModifierDummy(label)
	return TextView(modifier, SubTitleTextViewValue(label))
}

func PriceTextViewValue(label string) properties.Value {
	return TextViewValue(label, 10, 500)
}
func PriceTextView(label string) *textView {
	var modifier = properties.ModifierDummy(label)
	return TextView(modifier, PriceTextViewValue(label))
}
