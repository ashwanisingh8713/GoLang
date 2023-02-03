package properties

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/corner"
)

type Value struct {
	TextColor  Color  `json:"textColor"`
	TextSize   int    `json:"textSize,omitempty"`
	Weight     int    `json:"weight,omitempty"`
	Label      string `json:"label,omitempty"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	ImageUrl   string `json:"imageUrl,omitempty"`
	ImageScale string `json:"imageScale,omitempty"`
	FontStyle  string `json:"fontStyle,omitempty"`
}

type Color struct {
	Hue        int     `json:"hue"`
	Saturation float64 `json:"saturation"`
	Lighting   float64 `json:"lighting"`
	Alpha      float64 `json:"alpha"`
}

type Modifier struct {
	Label           string `json:"label"`
	BackgroundColor Color  `json:"backgroundColor"`
	BorderColor     Color  `json:"borderColor"`
	BorderWidth     int    `json:"borderWidth"`
	PaddingL        int    `json:"paddingL"`
	PaddingT        int    `json:"paddingT"`
	PaddingR        int    `json:"paddingR"`
	PaddingB        int    `json:"paddingB"`
	Corner          string `json:"cornerType"`
	IsEnable        bool   `json:"isEnable"`
	Radius          int    `json:"radius"`
	ColumnAlignment string `json:"columnAlignment"`
	TextAlignment   string `json:"textAlignment"`
}

type ProductInfo struct {
	OfferPercentage int    `json:"offerPercentage,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	Price           int    `json:"price,omitempty"`
	SellingScale    string `json:"sellingScale,omitempty"` // $3/kg
}

type Action struct {
	Url    string `json:"url,omitempty"`
	Method string `json:"method"`
}

func ColorObject(hue int, saturation float64, lighting float64, alpha float64) Color {
	return Color{Hue: hue, Saturation: saturation, Lighting: lighting, Alpha: alpha}
}

func ModifierObject(Label string, color Color, paddingLeft int, paddingTop int, paddingRight int, paddingBottom int, corner string) Modifier {
	return Modifier{Label: Label, BackgroundColor: color, PaddingL: paddingLeft,
		PaddingT: paddingTop, PaddingR: paddingRight, PaddingB: paddingBottom, Corner: corner}
}

func ColorDummy() Color {
	return ColorObject(0, 0.0, 0.0, 0.0)
}

func BorderColorDummy() Color {
	return ColorObject(100, 0.7, 0.7, 0.7)
}

func ModifierDummy(label string) Modifier {
	var col = ColorDummy()
	return ModifierObject(label, col, 5, 5, 5, 5, corner.CornerType.String(corner.ROUNDED_CORNER))
}
