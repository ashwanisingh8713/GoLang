package cashback

import (
	"ServerDrivenUI/src/E-Commerce/data"
	"ServerDrivenUI/src/E-Commerce/ec_constant"
	"ServerDrivenUI/src/E-Commerce/ui_type"
)

func Cashback(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt04.String()
	var explore_api = ec_constant.IpAddress + ec_constant.ROUTE_EXPLORE + "?" + ec_constant.QUERY_KEY_CATEGORY + "="

	var header = data.Header{}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "soft_drink.png",
	}
	var productInfo = data.ProductInfo{
		DisplaySellingScale: "Get 25% Cashback",
		UnitPrice:           20,
		OfferPrice:          15,
		Quantity:            1,
		OfferPercentage:     "10",
	}

	var product1 = data.Product{
		Type:  templateType,
		Title: "on all baby products",
		Action: data.Action{
			API:         explore_api,
			Query:       ec_constant.QUERY_VALUE_DEAL_OF_WEEK,
			Method:      "Get",
			Destination: ec_constant.Destination_Explore,
		},
		Value:       value1,
		ProductInfo: productInfo,
	}

	viewGrpChildren = append(viewGrpChildren, product1)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.Box.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}
