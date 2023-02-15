package data

type Header struct {
	Type            string `json:"type"`
	Title           string `json:"title"`
	CtaTitle        string `json:"ctaTitle"`
	BackgroundColor Color  `json:"background_color"`
	Action          Action `json:"action"`
}

type Value struct {
	Thumb1x  string `json:"thumb_1x"`
	Thumb2x  string `json:"thumb_2x"`
	Thumb3x  string `json:"thumb_3x"`
	Thumb4x  string `json:"thumb_4x"`
	Banner1x string `json:"banner_1x"`
	Banner2x string `json:"banner_2x"`
	Banner3x string `json:"banner_3x"`
	Banner4x string `json:"banner_4x"`
}

type Action struct {
	Method string `json:"method"`
	API    string `json:"api"`
}

type Color struct {
	Hue        int     `json:"hue"`
	Saturation float64 `json:"saturation"`
	Lighting   float64 `json:"lighting"`
	Alpha      float64 `json:"alpha"`
}

type ProductInfo struct {
	OfferPercentage     string `json:"offer_percentage"`
	Quantity            int    `json:"quantity"`
	DisplaySellingScale string `json:"display_selling_scale"`
	UnitPrice           int    `json:"unit_price"`
	OfferPrice          int    `json:"offer_price"`
}

type Product struct {
	ProductId       string      `json:"product_id"`
	Type            string      `json:"type"`
	Title           string      `json:"title"`
	BackgroundColor Color       `json:"background_color"`
	Value           Value       `json:"value"`
	ProductInfo     ProductInfo `json:"product_info"`
	Action          Action      `json:"action"`
}

type ViewGroup struct {
	ViewGroupType string    `json:"view_group_type"`
	Header        Header    `json:"header"`
	Children      []Product `json:"children"`
	GridColumn    int       `json:"gridColumn" binding:"required"`
}
