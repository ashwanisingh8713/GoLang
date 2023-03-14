package subscription

import (
	"E-Commerce/data"
	"E-Commerce/ec_constant"
	"E-Commerce/ui_type"
)

func SubscriptionInfo(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt10.String()
	var viewGrpChildren []data.Product
	var subscriptions []data.SubscriptionInfo

	var subcription1 = data.SubscriptionInfo{
		SubscriptionId: "",
		ProductId:      "",
		OrderId:        "",
		CreatedAt:      "",
		StartDate:      "",
		EndDate:        "",
		Monday:         true,
		Thursday:       true,
		Wednesday:      false,
		Tuesday:        true,
		Friday:         true,
		Saturday:       false,
		Sunday:         true,
		PauseDates:     "",
	}

	var subcription2 = data.SubscriptionInfo{
		SubscriptionId: "",
		ProductId:      "",
		OrderId:        "",
		CreatedAt:      "",
		StartDate:      "",
		EndDate:        "",
		Monday:         true,
		Thursday:       true,
		Wednesday:      false,
		Tuesday:        true,
		Friday:         true,
		Saturday:       false,
		Sunday:         true,
		PauseDates:     "",
	}

	var subcription3 = data.SubscriptionInfo{
		SubscriptionId: "",
		ProductId:      "",
		OrderId:        "",
		CreatedAt:      "",
		StartDate:      "",
		EndDate:        "",
		Monday:         true,
		Thursday:       true,
		Wednesday:      false,
		Tuesday:        true,
		Friday:         true,
		Saturday:       false,
		Sunday:         true,
		PauseDates:     "",
	}

	var subcription4 = data.SubscriptionInfo{
		SubscriptionId: "",
		ProductId:      "",
		OrderId:        "",
		CreatedAt:      "",
		StartDate:      "",
		EndDate:        "",
		Monday:         true,
		Thursday:       true,
		Wednesday:      false,
		Tuesday:        true,
		Friday:         true,
		Saturday:       false,
		Sunday:         true,
		PauseDates:     "",
	}

	var value1 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "jaggery.png",
	}

	var product1 = data.Product{
		ProductId: "10",
		Type:      templateType,
		Title:     "Fortune Rice",
		Action: data.Action{
			API:         "",
			Query:       ec_constant.QUERY_VALUE_PRODUCTID,
			Method:      "Get",
			Destination: ec_constant.Destination_Detail,
		},
		Value:            value1,
		SubscriptionInfo: subcription1,
	}
	var value2 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "jaggery.png",
	}

	var product2 = data.Product{
		Type:  templateType,
		Title: "Fresh Avacado",
		Action: data.Action{
			API:         "",
			Query:       ec_constant.QUERY_VALUE_PRODUCTID,
			Method:      "Get",
			Destination: ec_constant.Destination_Detail,
		},
		Value:            value2,
		SubscriptionInfo: subcription2,
	}
	var value3 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "jaggery.png",
	}

	var product3 = data.Product{
		Type:  templateType,
		Title: "Orange",
		Action: data.Action{
			API:         "",
			Query:       ec_constant.QUERY_VALUE_PRODUCTID,
			Method:      "Get",
			Destination: ec_constant.Destination_Detail,
		},
		Value:            value3,
		SubscriptionInfo: subcription3,
	}
	var value4 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "jaggery.png",
	}

	var product4 = data.Product{
		Type:  templateType,
		Title: "Potato",
		Action: data.Action{
			API:         "",
			Query:       ec_constant.QUERY_VALUE_PRODUCTID,
			Method:      "Get",
			Destination: ec_constant.Destination_Detail,
		},
		Value:            value4,
		SubscriptionInfo: subcription4,
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)
	viewGrpChildren = append(viewGrpChildren, product4)

	var viewGroup = data.ViewGroupSubscription{
		ViewGroupType: ui_type.VerticalGrid.String(),
		Children:      viewGrpChildren,
		Subscription:  subscriptions,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}
