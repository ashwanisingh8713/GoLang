package properties

type Color struct {
	Hue        int     `json:"hue"`
	Saturation float64 `json:"saturation"`
	Lighting   float64 `json:"lighting"`
	Alpha      float64 `json:"alpha"`
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

func ColorDummy() Color {
	return ColorObject(0, 0.0, 0.0, 0.0)
}

func BorderColorDummy() Color {
	return ColorObject(100, 0.7, 0.7, 0.7)
}
