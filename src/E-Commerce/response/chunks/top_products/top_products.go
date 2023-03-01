package top_products

import (
	"ServerDrivenUI/src/E-Commerce/data"
	"ServerDrivenUI/src/E-Commerce/ec_constant"
	"ServerDrivenUI/src/E-Commerce/ui_type"
)

func TopProducts(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt03.String()
	var explore_api = ec_constant.IpAddress + ec_constant.ROUTE_EXPLORE + "?" + ec_constant.QUERY_KEY_CATEGORY + "="

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "Top Products",
		CtaTitle: "Explore All",
		Action: data.Action{
			API:         explore_api,
			Query:       ec_constant.QUERY_VALUE_TOPPRODUCT_EXPLORE,
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
		Title: "Fortune Rice",
		Action: data.Action{
			API:         "",
			Query:       ec_constant.QUERY_VALUE_PRODUCTID,
			Method:      "Get",
			Destination: ec_constant.Destination_Detail,
		},
		Value: value1,
	}

	var value2 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "tea.png",
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
		Value: value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "potato.png",
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
