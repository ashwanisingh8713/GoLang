package top_categogies

import (
	"E-Commerce/data"
	"E-Commerce/ec_constant"
	"E-Commerce/ui_type"
)

func ExploreGroceries(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt08.String()

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
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
		Banner4x: ec_constant.GroceriesImagePath4x + "saffolagoldoil.png",
	}
	var product2 = data.Product{
		Type:   templateType,
		Title:  "Saffola gold oil",
		Action: data.Action{},
		Value:  value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.DealOfTheWeekImagePath4x + "tataarhardal.png",
	}
	var product3 = data.Product{
		Type:   templateType,
		Title:  "TATA Arhar Dal",
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

func ExploreVegetables(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt07.String()

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "Vegetables",
		CtaTitle: "See All",
		Action:   data.Action{},
	}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.VegetablesImagePath4x + "tomato.png",
	}

	var product1 = data.Product{
		Type:   templateType,
		Title:  "Tomato",
		Action: data.Action{},
		Value:  value1,
	}

	var value2 = data.Value{
		Banner4x: ec_constant.VegetablesImagePath4x + "potato.png",
	}
	var product2 = data.Product{
		Type:   templateType,
		Title:  "Potato",
		Action: data.Action{},
		Value:  value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.FruitsRelativePath4x + "onion.png",
	}
	var product3 = data.Product{
		Type:   templateType,
		Title:  "Onion",
		Action: data.Action{},
		Value:  value3,
	}

	var value4 = data.Value{
		Banner4x: ec_constant.VegetablesImagePath4x + "potato.png",
	}
	var product4 = data.Product{
		Type:   templateType,
		Title:  "Potato",
		Action: data.Action{},
		Value:  value4,
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)
	viewGrpChildren = append(viewGrpChildren, product4)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.HorizontalGrid.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}

func ExploreFruits(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt07.String()

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "Fruits",
		CtaTitle: "See All",
		Action:   data.Action{},
	}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.FruitsImagePath4x + "banana.png",
	}

	var product1 = data.Product{
		Type:   templateType,
		Title:  "Banana",
		Action: data.Action{},
		Value:  value1,
	}

	var value2 = data.Value{
		Banner4x: ec_constant.FruitsImagePath4x + "kiwi.png",
	}
	var product2 = data.Product{
		Type:   templateType,
		Title:  "Kiwi",
		Action: data.Action{},
		Value:  value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.FruitsRelativePath4x + "straberry.png",
	}
	var product3 = data.Product{
		Type:   templateType,
		Title:  "Straw berry",
		Action: data.Action{},
		Value:  value3,
	}

	var value4 = data.Value{
		Banner4x: ec_constant.VegetablesImagePath4x + "kiwi.png",
	}
	var product4 = data.Product{
		Type:   templateType,
		Title:  "Kiwi",
		Action: data.Action{},
		Value:  value4,
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)
	viewGrpChildren = append(viewGrpChildren, product4)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.HorizontalGrid.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}
func ExploreDairyProduct(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt07.String()

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "Dairy Products",
		CtaTitle: "See All",
		Action:   data.Action{},
	}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.DairyProductsImagePath4x + "butter.png",
	}

	var product1 = data.Product{
		Type:   templateType,
		Title:  "Amul Butter",
		Action: data.Action{},
		Value:  value1,
	}

	var value2 = data.Value{
		Banner4x: ec_constant.FruitsImagePath4x + "milk.png",
	}
	var product2 = data.Product{
		Type:   templateType,
		Title:  "A2MATE milk",
		Action: data.Action{},
		Value:  value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.FruitsRelativePath4x + "softdrink.png",
	}
	var product3 = data.Product{
		Type:   templateType,
		Title:  "Sofit soya milk",
		Action: data.Action{},
		Value:  value3,
	}

	var value4 = data.Value{
		Banner4x: ec_constant.VegetablesImagePath4x + "milk.png",
	}
	var product4 = data.Product{
		Type:   templateType,
		Title:  "A2MATE milk",
		Action: data.Action{},
		Value:  value4,
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)
	viewGrpChildren = append(viewGrpChildren, product4)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.HorizontalGrid.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}
func ExploreBakeryProduct(parentChildren []interface{}) []interface{} {
	var templateType = ui_type.Vt07.String()

	var header = data.Header{
		Type:     ui_type.Header_01.String(),
		Title:    "Bakery Items",
		CtaTitle: "See All",
		Action:   data.Action{},
	}

	var viewGrpChildren []data.Product

	var value1 = data.Value{
		Banner4x: ec_constant.BannerImagePath4x + "rusk.png",
	}

	var product1 = data.Product{
		Type:   templateType,
		Title:  "Parle Rusk",
		Action: data.Action{},
		Value:  value1,
	}

	var value2 = data.Value{
		Banner4x: ec_constant.FruitsImagePath4x + "muffin.png",
	}
	var product2 = data.Product{
		Type:   templateType,
		Title:  "Choco muffin",
		Action: data.Action{},
		Value:  value2,
	}

	var value3 = data.Value{
		Banner4x: ec_constant.FruitsRelativePath4x + "bar.png",
	}
	var product3 = data.Product{
		Type:   templateType,
		Title:  "Harshey’s Bar",
		Action: data.Action{},
		Value:  value3,
	}

	var value4 = data.Value{
		Banner4x: ec_constant.VegetablesImagePath4x + "muffin.png",
	}
	var product4 = data.Product{
		Type:   templateType,
		Title:  "Choco muffin",
		Action: data.Action{},
		Value:  value4,
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)
	viewGrpChildren = append(viewGrpChildren, product4)

	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.HorizontalGrid.String(),
		Header:        header,
		Children:      viewGrpChildren,
		GridColumn:    1,
	}

	return append(parentChildren, viewGroup)
}
