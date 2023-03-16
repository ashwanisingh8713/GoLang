package response

type ViewGroupMockResponse struct {
	Data []struct {
		ViewGroupType string `json:"view_group_type"`
		Header        struct {
			Type            string `json:"type"`
			Title           string `json:"title"`
			CtaTitle        string `json:"ctaTitle"`
			BackgroundColor struct {
				Hue        int `json:"hue"`
				Saturation int `json:"saturation"`
				Lighting   int `json:"lighting"`
				Alpha      int `json:"alpha"`
			} `json:"background_color"`
			Action struct {
				Method      string `json:"method"`
				Api         string `json:"api"`
				Destination string `json:"destination"`
				Query       string `json:"query"`
			} `json:"action"`
		} `json:"header"`
		Children []struct {
			ProductId       string `json:"product_id"`
			Type            string `json:"type"`
			Title           string `json:"title"`
			BackgroundColor struct {
				Hue        int `json:"hue"`
				Saturation int `json:"saturation"`
				Lighting   int `json:"lighting"`
				Alpha      int `json:"alpha"`
			} `json:"background_color"`
			Value struct {
				Thumb1X  string `json:"thumb_1x"`
				Thumb2X  string `json:"thumb_2x"`
				Thumb3X  string `json:"thumb_3x"`
				Thumb4X  string `json:"thumb_4x"`
				Banner1X string `json:"banner_1x"`
				Banner2X string `json:"banner_2x"`
				Banner3X string `json:"banner_3x"`
				Banner4X string `json:"banner_4x"`
			} `json:"value"`
			ProductInfo struct {
				OfferPercentage     string `json:"offer_percentage"`
				Quantity            int    `json:"quantity"`
				DisplaySellingScale string `json:"display_selling_scale"`
				UnitPrice           int    `json:"unit_price"`
				OfferPrice          int    `json:"offer_price"`
			} `json:"product_info"`
			Action struct {
				Method      string `json:"method"`
				Api         string `json:"api"`
				Destination string `json:"destination"`
				Query       string `json:"query"`
			} `json:"action"`
			SubscriptionInfo struct {
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
			} `json:"subscription_info"`
		} `json:"children"`
		GridColumn int `json:"gridColumn"`
	} `json:"data"`
}

type SubscriptionPageMockResponse struct {
	Data []struct {
		ViewGroupType string `json:"view_group_type"`
		Children      []struct {
			ProductId       string `json:"product_id"`
			Type            string `json:"type"`
			Title           string `json:"title"`
			BackgroundColor struct {
				Hue        int `json:"hue"`
				Saturation int `json:"saturation"`
				Lighting   int `json:"lighting"`
				Alpha      int `json:"alpha"`
			} `json:"background_color"`
			Value struct {
				Thumb1X  string `json:"thumb_1x"`
				Thumb2X  string `json:"thumb_2x"`
				Thumb3X  string `json:"thumb_3x"`
				Thumb4X  string `json:"thumb_4x"`
				Banner1X string `json:"banner_1x"`
				Banner2X string `json:"banner_2x"`
				Banner3X string `json:"banner_3x"`
				Banner4X string `json:"banner_4x"`
			} `json:"value"`
			ProductInfo struct {
				OfferPercentage     string `json:"offer_percentage"`
				Quantity            int    `json:"quantity"`
				DisplaySellingScale string `json:"display_selling_scale"`
				UnitPrice           int    `json:"unit_price"`
				OfferPrice          int    `json:"offer_price"`
			} `json:"product_info"`
			Action struct {
				Method      string `json:"method"`
				Api         string `json:"api"`
				Destination string `json:"destination"`
				Query       string `json:"query"`
			} `json:"action"`
			SubscriptionInfo struct {
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
			} `json:"subscription_info"`
		} `json:"children"`
		Subscription interface{} `json:"subscription"`
		GridColumn   int         `json:"gridColumn"`
	} `json:"data"`
}
