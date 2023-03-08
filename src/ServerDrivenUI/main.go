package main

import (
	"ServerDrivenUI/constants"
	"ServerDrivenUI/models/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.GET("/ui", controller.Response)

	// Bakery Items
	route.Static(constants.BakeryItemsRelativePath4x, constants.BakeryItemsFolderPath4x)
	// Dairy Products
	route.Static(constants.DairyProductsRelativePath4x, constants.DairyProductsFolderPath4x)
	// Deal Of the Week
	route.Static(constants.DealOfTheWeekRelativePath4x, constants.DealOfTheWeekFolderPath4x)
	// Featured Items
	route.Static(constants.FeaturedItemsRelativePath4x, constants.FeaturedItemsFolderPath4x)
	// Fruits
	route.Static(constants.FruitsRelativePath4x, constants.FruitsFolderPath4x)
	// Groceries
	route.Static(constants.GroceriesRelativePath4x, constants.GroceriesFolderPath4x)
	// Top Categories
	route.Static(constants.TopCategoriesRelativePath4x, constants.TopCategoriesFolderPath4x)
	// Top Products
	route.Static(constants.TopProductsRelativePath4x, constants.TopProductsFolderPath4x)
	// Vegetables
	route.Static(constants.VegetablesRelativePath4x, constants.VegetablesFolderPath4x)

	err := route.Run(constants.IpAddress)
	if err != nil {
		return
	}

}
