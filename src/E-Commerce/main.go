package main

import (
	"ServerDrivenUI/src/E-Commerce/ec_constant"
	"ServerDrivenUI/src/E-Commerce/response"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	route.GET(ec_constant.ROUTE_HOME, response.Home)
	route.GET(ec_constant.ROUTE_EXPLORE, response.ExploreTopCategory)
	route.GET(ec_constant.ROUTE_PRODUCT, response.ExploreTopCategory)
	//route.GET("/category/explore", response.ExploreCategory)

	// Bakery Items
	route.Static(ec_constant.BakeryItemsRelativePath4x, ec_constant.BakeryItemsFolderPath4x)
	// Dairy Products
	route.Static(ec_constant.DairyProductsRelativePath4x, ec_constant.DairyProductsFolderPath4x)
	// Deal Of the Week
	route.Static(ec_constant.DealOfTheWeekRelativePath4x, ec_constant.DealOfTheWeekFolderPath4x)
	// Featured Items
	route.Static(ec_constant.FeaturedItemsRelativePath4x, ec_constant.FeaturedItemsFolderPath4x)
	// Fruits
	route.Static(ec_constant.FruitsRelativePath4x, ec_constant.FruitsFolderPath4x)
	// Groceries
	route.Static(ec_constant.GroceriesRelativePath4x, ec_constant.GroceriesFolderPath4x)
	// Top Categories
	route.Static(ec_constant.TopCategoriesRelativePath4x, ec_constant.TopCategoriesFolderPath4x)
	// Top Products
	route.Static(ec_constant.TopProductsRelativePath4x, ec_constant.TopProductsFolderPath4x)
	// Top Viewpager
	route.Static(ec_constant.TopViewpagerRelativePath4x, ec_constant.TopViewpagerFolderPath4x)
	// Vegetables
	route.Static(ec_constant.VegetablesRelativePath4x, ec_constant.VegetablesFolderPath4x)

	err := route.Run(ec_constant.IpAddress)
	if err != nil {
		return
	}
}
