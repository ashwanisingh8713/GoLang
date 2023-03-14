package explore_category

import (
	"E-Commerce/data"
	"E-Commerce/ec_constant"
	"E-Commerce/ui_type"
)

func ExploreCategory(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt08.String()

	var header = data.Header{
		Type:     "",
		Title:    "Groceries",
		CtaTitle: "See all",
		Action:   data.Action{},
	}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "jaggery.png",
	}

	var product1 = data.Product{
		Type:   templateType,
		Title:  "Jaggery Powder",
		Action: data.Action{},
		Value:  value1,
	}

	var value2 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "banana.jpg",
	}
	var product2 = data.Product{
		Type:   templateType,
		Title:  "Banana",
		Action: data.Action{},
		Value:  value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "potato.png",
	}
	var product3 = data.Product{
		Type:   templateType,
		Title:  "Potato",
		Action: data.Action{},
		Value:  value3,
	}

	var value4 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "soft_drink.png",
	}
	var product4 = data.Product{
		Type:   templateType,
		Title:  "Soft Drink",
		Action: data.Action{},
		Value:  value4,
	}

	var value5 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "tea.png",
	}
	var product5 = data.Product{
		Type:   templateType,
		Title:  "Tea",
		Action: data.Action{},
		Value:  value5,
	}

	var value6 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "banana.jpg",
	}
	var product6 = data.Product{
		Type:   templateType,
		Title:  "Banana",
		Action: data.Action{},
		Value:  value6,
	}

	var value7 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "potato.png",
	}
	var product7 = data.Product{
		Type:   templateType,
		Title:  "Potato",
		Action: data.Action{},
		Value:  value7,
	}

	var value8 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "tea.png",
	}
	var product8 = data.Product{
		Type:   templateType,
		Title:  "Tea",
		Action: data.Action{},
		Value:  value8,
	}

	var value9 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "tataarhardal.png",
	}
	var product9 = data.Product{
		Type:   templateType,
		Title:  "TATA Arhar Dal",
		Action: data.Action{},
		Value:  value9,
	}

	var value10 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "tataarhardal.png",
	}
	var product10 = data.Product{
		Type:   templateType,
		Title:  "TATA Arhar Dal",
		Action: data.Action{},
		Value:  value10,
	}

	var value11 = data.Value{
		Banner4x: ec_constant.GroceriesImagePath4x + "saffolagoldoil.png",
	}
	var product11 = data.Product{
		Type:   templateType,
		Title:  "Suffola Gold Oil",
		Action: data.Action{},
		Value:  value11,
	}

	var value12 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "tea.png",
	}
	var product12 = data.Product{
		Type:   templateType,
		Title:  "Tea",
		Action: data.Action{},
		Value:  value12,
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)
	viewGrpChildren = append(viewGrpChildren, product4)
	viewGrpChildren = append(viewGrpChildren, product5)
	viewGrpChildren = append(viewGrpChildren, product6)
	viewGrpChildren = append(viewGrpChildren, product7)
	viewGrpChildren = append(viewGrpChildren, product8)
	viewGrpChildren = append(viewGrpChildren, product9)
	viewGrpChildren = append(viewGrpChildren, product10)
	viewGrpChildren = append(viewGrpChildren, product11)
	viewGrpChildren = append(viewGrpChildren, product12)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.VerticalGrid.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    2,
	}

	return append(parentChildren, viewGroup)
}
