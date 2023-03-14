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
	Method      string `json:"method"`
	API         string `json:"api"`
	Destination string `json:"destination"` //Screen destination
	Query       string `json:"query"`       // API Query
}

type Profile struct {
	UserId     string       `json:"user_id"`
	ProfilePic string       `json:"profile_pic"`
	UserName   string       `json:"user_name"`
	UserEmail  string       `json:"user_email"`
	CtaProfile []ProfileCta `json:"cta_profile"`
}
type ProfileCta struct {
	Value    Value  `json:"value"`
	CtaTitle string `json:"title"`
	Action   Action `json:"action"`
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

type SubscriptionInfo struct {
	SubscriptionId string `json:"subscription_id"`
	ProductId      string `json:"product_id"`
	OrderId        string `json:"order_id"`
	UserId         string `json:"user_id"`
	CreatedAt      string `json:"created_at"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Monday         bool   `json:"monday"`
	Tuesday        bool   `json:"tuesday"`
	Wednesday      bool   `json:"wednesday"`
	Thursday       bool   `json:"thursday"`
	Friday         bool   `json:"friday"`
	Saturday       bool   `json:"saturday"`
	Sunday         bool   `json:"sunday"`
	PauseDates     string `json:"pause_dates"`
}

type Product struct {
	ProductId        string           `json:"product_id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	BackgroundColor  Color            `json:"background_color"`
	Value            Value            `json:"value"`
	ProductInfo      ProductInfo      `json:"product_info"`
	Action           Action           `json:"action"`
	SubscriptionInfo SubscriptionInfo `json:"subscription_info,omitempty"`
}

type OnBoarding struct {
	Value Value  `json:"value"`
	Title string `json:"title"`
}
type ViewGroupOnboarding struct {
	ViewGroupType string       `json:"view_group_type"`
	Children      []OnBoarding `json:"children"`
}

type ViewGroup struct {
	ViewGroupType string    `json:"view_group_type"`
	Header        Header    `json:"header"`
	Children      []Product `json:"children"`
	GridColumn    int       `json:"gridColumn" binding:"required"`
}

type ViewGroupSubscription struct {
	ViewGroupType string             `json:"view_group_type"`
	Children      []Product          `json:"children"`
	Subscription  []SubscriptionInfo `json:"subscription"`
	GridColumn    int                `json:"gridColumn" binding:"required"`
}
type ViewGroupProfile struct {
	ViewGroupType string  `json:"view_group_type"`
	Profile       Profile `json:"profile"`
}
