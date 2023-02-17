package featured_items

import (
	"ServerDrivenUI/src/E-Commerce/data"
	"ServerDrivenUI/src/E-Commerce/ec_constant"
	"ServerDrivenUI/src/E-Commerce/ui_type"
)

func FeaturedItems(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt06.String()

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "Featured Items",
		CtaTitle: "Explore All",
		Action:   data.Action{},
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

	var value2 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "tea.png",
	}
	var product2 = data.Product{
		Type:   templateType,
		Title:  "Fresh Avacado",
		Action: data.Action{},
		Value:  value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "potato.png",
	}
	var product3 = data.Product{
		Type:   templateType,
		Title:  "Orange",
		Action: data.Action{},
		Value:  value3,
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
