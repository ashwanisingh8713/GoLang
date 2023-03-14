package userdata

import (
	"E-Commerce/data"
	"E-Commerce/ec_constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Profile(c *gin.Context) {

	var viewGrpCtaChildren []data.ProfileCta
	viewGrpCtaChildren = myOrder(viewGrpCtaChildren)
	viewGrpCtaChildren = MySubscription(viewGrpCtaChildren)
	viewGrpCtaChildren = MyAddresses(viewGrpCtaChildren)
	viewGrpCtaChildren = FAQ(viewGrpCtaChildren)
	viewGrpCtaChildren = ContactUs(viewGrpCtaChildren)
	viewGrpCtaChildren = About(viewGrpCtaChildren)
	viewGrpCtaChildren = LogOut(viewGrpCtaChildren)

	var profileInfo = data.Profile{
		ProfilePic: "https://cdn.pixabay.com/photo/2016/11/18/23/38/child-1837375__340.png",
		UserId:     "97a6c53e-b9a6-11ed-bb41-16ab030fa2fa",
		UserName:   "Ninestars",
		UserEmail:  "ninestars@gmail.com",
		CtaProfile: viewGrpCtaChildren,
	}
	c.JSON(http.StatusOK, gin.H{"data": profileInfo})
}

func myOrder(viewGrpCtaChildren []data.ProfileCta) []data.ProfileCta {
	var valueOrder = data.Value{
		Banner4x: ec_constant.ProfileImagePath4x + "order.png",
	}

	var actionOrder = data.Action{
		API:         "",
		Query:       "/order",
		Method:      "Get",
		Destination: "Order",
	}

	var profileData = data.ProfileCta{
		Value:    valueOrder,
		Action:   actionOrder,
		CtaTitle: "My Order",
	}

	viewGrpCtaChildren = append(viewGrpCtaChildren, profileData)
	return viewGrpCtaChildren
}
func MySubscription(viewGrpCtaChildren []data.ProfileCta) []data.ProfileCta {
	var valueSubscription = data.Value{
		Banner4x: ec_constant.ProfileImagePath4x + "subscription.png",
	}
	var actionSubscription = data.Action{
		API:         "",
		Query:       "/subscription",
		Method:      "Get",
		Destination: "Subscription",
	}

	var profileData = data.ProfileCta{
		Value:    valueSubscription,
		Action:   actionSubscription,
		CtaTitle: "My Subscription",
	}

	viewGrpCtaChildren = append(viewGrpCtaChildren, profileData)
	return viewGrpCtaChildren
}
func MyAddresses(viewGrpCtaChildren []data.ProfileCta) []data.ProfileCta {
	var valueAddress = data.Value{
		Banner4x: ec_constant.ProfileImagePath4x + "address.png",
	}
	var actionAddress = data.Action{
		API:         "",
		Query:       "/addresses",
		Method:      "Get",
		Destination: "Addresses",
	}

	var profileData = data.ProfileCta{
		Value:    valueAddress,
		Action:   actionAddress,
		CtaTitle: "My Addresses",
	}

	viewGrpCtaChildren = append(viewGrpCtaChildren, profileData)
	return viewGrpCtaChildren
}
func FAQ(viewGrpCtaChildren []data.ProfileCta) []data.ProfileCta {
	var valueFaq = data.Value{
		Banner4x: ec_constant.ProfileImagePath4x + "faq.png",
	}
	var actionFaq = data.Action{
		API:         "",
		Query:       "/faq",
		Method:      "Get",
		Destination: "FAQ",
	}

	var profileData = data.ProfileCta{
		Value:    valueFaq,
		Action:   actionFaq,
		CtaTitle: "FAQ",
	}

	viewGrpCtaChildren = append(viewGrpCtaChildren, profileData)
	return viewGrpCtaChildren
}
func ContactUs(viewGrpCtaChildren []data.ProfileCta) []data.ProfileCta {
	var valueContactUs = data.Value{
		Banner4x: ec_constant.ProfileImagePath4x + "contact.png",
	}
	var actionContactUs = data.Action{
		API:         "",
		Query:       "/contact_us",
		Method:      "Get",
		Destination: "Contact",
	}

	var profileData = data.ProfileCta{
		Value:    valueContactUs,
		Action:   actionContactUs,
		CtaTitle: "Contact Us",
	}

	viewGrpCtaChildren = append(viewGrpCtaChildren, profileData)
	return viewGrpCtaChildren
}
func About(viewGrpCtaChildren []data.ProfileCta) []data.ProfileCta {
	var valueAbout = data.Value{
		Banner4x: ec_constant.ProfileImagePath4x + "about.png",
	}
	var actionAbout = data.Action{
		API:         "",
		Query:       "/about",
		Method:      "Get",
		Destination: "About",
	}

	var profileData = data.ProfileCta{
		Value:    valueAbout,
		Action:   actionAbout,
		CtaTitle: "About",
	}

	viewGrpCtaChildren = append(viewGrpCtaChildren, profileData)
	return viewGrpCtaChildren
}
func LogOut(viewGrpCtaChildren []data.ProfileCta) []data.ProfileCta {
	var valueLogout = data.Value{
		Banner4x: ec_constant.ProfileImagePath4x + "logout.png",
	}
	var actionLogout = data.Action{
		API:         "",
		Query:       "/logout",
		Method:      "Get",
		Destination: "LogOut",
	}

	var profileData = data.ProfileCta{
		Value:    valueLogout,
		Action:   actionLogout,
		CtaTitle: "About",
	}

	viewGrpCtaChildren = append(viewGrpCtaChildren, profileData)
	return viewGrpCtaChildren
}
