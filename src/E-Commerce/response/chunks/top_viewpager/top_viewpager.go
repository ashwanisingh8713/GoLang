package top_viewpager

import (
	"E-Commerce/data"
	"E-Commerce/ec_constant"
	"E-Commerce/res/drawable/top_viewpager"
	"E-Commerce/ui_type"
)

func TopViewpager(parentChildren []interface{}) []interface{} {

	var viewGrpChildren []data.Product

	var productValue = data.Value{
		Banner4x: ec_constant.TopViewpagerImagePath4x + top_viewpager.AnimatedImage,
	}
	var action = data.Action{}

	var product = data.Product{
		Type:   ui_type.Vt01.String(),
		Title:  "EveryDay Essential",
		Value:  productValue,
		Action: action,
		ProductInfo: data.ProductInfo{
			DisplaySellingScale: "10% Off",
			UnitPrice:           20,
			OfferPrice:          15,
			Quantity:            1,
			OfferPercentage:     "10",
		},
	}
	var product1 = data.Product{
		Type:   ui_type.Vt01.String(),
		Title:  "EveryDay Essential",
		Value:  productValue,
		Action: action,
		ProductInfo: data.ProductInfo{
			DisplaySellingScale: "10% Off",
			UnitPrice:           20,
			OfferPrice:          15,
			Quantity:            1,
			OfferPercentage:     "10",
		},
	}
	var product2 = data.Product{
		Type:   ui_type.Vt01.String(),
		Title:  "EveryDay Essential",
		Value:  productValue,
		Action: action,
		ProductInfo: data.ProductInfo{
			DisplaySellingScale: "10% Off",
			UnitPrice:           20,
			OfferPrice:          15,
			Quantity:            1,
			OfferPercentage:     "10",
		},
	}

	viewGrpChildren = append(viewGrpChildren, product)
	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)

	var header = data.Header{
		Type:   "",
		Action: data.Action{},
		Title:  "Header View Group",
	}
	var viewGroup = data.ViewGroup{
		ViewGroupType: ui_type.ViewPager.String(),
		Header:        header,
		Children:      viewGrpChildren,
	}

	return append(parentChildren, viewGroup)
}
