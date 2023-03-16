package onboarding

import (
	"E-Commerce/data"
	"E-Commerce/ec_constant"
	"E-Commerce/ui_type"
)

func OnBoarding(parentChildren []interface{}) []interface{} {

	var viewGrpChildren []data.OnBoarding
	var value1 = data.Value{
		Banner4x: ec_constant.OnBoardingImagePath4x + "logo.png",
	}

	var value2 = data.Value{
		Banner4x: ec_constant.OnBoardingImagePath4x + "logo.png",
	}

	var value3 = data.Value{
		Banner4x: ec_constant.OnBoardingImagePath4x + "logo.png",
	}
	var product1 = data.OnBoarding{
		Value: value1,
		Title: "Your needs in just one place1",
	}

	var product2 = data.OnBoarding{
		Value: value2,
		Title: "Your needs in just one place2",
	}

	var product3 = data.OnBoarding{
		Value: value3,
		Title: "Your needs in just one place3",
	}

	viewGrpChildren = append(viewGrpChildren, product1)
	viewGrpChildren = append(viewGrpChildren, product2)
	viewGrpChildren = append(viewGrpChildren, product3)

	var viewGroup = data.ViewGroupOnboarding{
		ViewGroupType: ui_type.ViewPager.String(),
		Children:      viewGrpChildren,
	}
	return append(parentChildren, viewGroup)

}
