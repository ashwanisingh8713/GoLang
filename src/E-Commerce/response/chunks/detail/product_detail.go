package detail

import (
	"ServerDrivenUI/src/E-Commerce/data"
	"ServerDrivenUI/src/E-Commerce/ec_constant"
	"ServerDrivenUI/src/E-Commerce/ui_type"
)

func ProductDetail(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt06.String()
	var explore_api = ec_constant.IpAddress + ec_constant.ROUTE_EXPLORE + "?" + ec_constant.QUERY_KEY_CATEGORY + "="

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "",
		CtaTitle: "",
		Action: data.Action{
			API:         explore_api,
			Query:       ec_constant.QUERY_VALUE_FEATURED_ITEM,
			Method:      "Get",
			Destination: ec_constant.Destination_Explore,
		},
	}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "soft_drink.png",
	}

	var product1 = data.Product{
		Type:   templateType,
		Title:  "Fortune Rice",
		Action: data.Action{},
		Value:  value1,
	}

	viewGrpChildren = append(viewGrpChildren, product1)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.HorizontalGrid.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}
