package deal_of_the_week

import (
	"E-Commerce/data"
	"E-Commerce/ec_constant"
	"E-Commerce/ui_type"
)

func DealOfTheWeek(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt05.String()
	var explore_api = ec_constant.IpAddress + ec_constant.ROUTE_EXPLORE + "?" + ec_constant.QUERY_KEY_CATEGORY + "="

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "Deals Of the week",
		CtaTitle: "Explore All",
		Action: data.Action{
			API:         explore_api,
			Query:       ec_constant.QUERY_VALUE_DEAL_OF_WEEK,
			Method:      "Get",
			Destination: ec_constant.Destination_Explore,
		},
	}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "soft_drink.png",
	}

	var product1 = data.Product{
		Type:  templateType,
		Title: "Soft Drinks",
		Action: data.Action{
			API:         explore_api,
			Query:       ec_constant.QUERY_VALUE_DEAL_OF_WEEK,
			Method:      "Get",
			Destination: ec_constant.Destination_Explore,
		},
		Value: value1,
	}

	var value2 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "tea.png",
	}
	var product2 = data.Product{
		Type:  templateType,
		Title: "Tea",
		Action: data.Action{
			API:         explore_api,
			Query:       ec_constant.QUERY_VALUE_DEAL_OF_WEEK,
			Method:      "Get",
			Destination: ec_constant.Destination_Explore,
		},
		Value: value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "potato.png",
	}
	var product3 = data.Product{
		Type:  templateType,
		Title: "Potato",
		Action: data.Action{
			API:         explore_api,
			Query:       ec_constant.QUERY_VALUE_DEAL_OF_WEEK,
			Method:      "Get",
			Destination: ec_constant.Destination_Explore,
		},
		Value: value3,
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.HorizontalGrid.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}
