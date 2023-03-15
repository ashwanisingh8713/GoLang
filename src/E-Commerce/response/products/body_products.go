package products

type BodyProductInput struct {
	CategoryId     string  `json:"category_id"`     //2
	Dimension      string  `json:"dimension"`       //3
	ManufacturerId string  `json:"manufacturer_id"` //4
	Name           string  `json:"name"`            //5
	Picture        string  `json:"picture"`         //6
	Price          float32 `json:"price"`           //7
	Description    string  `json:"description"`     //8
	SellerId       string  `json:"seller_id"`       //9
	SellerName     string  `json:"seller_name"`     //10
	Sku            string  `json:"sku"`             //11
	SubCategoryId  string  `json:"sub_category_id"` //12
	Weight         float32 `json:"weight"`          //13
	Units          int     `json:"units"`           //14
	Subscribable   bool    `json:"subscribable"`    //15
}

type BodyProductIdInput struct {
	ProductId string `json:"productId" binding:"required"`
}
