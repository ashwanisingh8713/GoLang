package main

import (
	"E-Commerce/database"
	_ "E-Commerce/docs"
	"E-Commerce/ec_constant"
	"E-Commerce/response"
	"E-Commerce/response/userdata"
	"E-Commerce/userData"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	//httpSwagger "github.com/swaggo/http-swagger"
)

// @title APIs of E-Commerce Project
// @version 1.0
// @description This is E-Commerce API Server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support (Ashwani)
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.13.119:8080
// @BasePath /
// @schemes http
func main() {

	// It creates singleton cassandra db instance
	database.CreateDBSession()

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Println("this is port", port)
	// Force log's color
	gin.ForceConsoleColor()

	route := gin.Default()
	route.Use(CORSMiddleware())

	// Swagger UI http://192.168.13.119:8080/swagger/index.html
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.GET(ec_constant.ROUTE_HOME, response.Home)
	route.GET(ec_constant.ROUTE_EXPLORE, response.ExploreTopCategory)
	route.GET(ec_constant.ROUTE_PRODUCT, response.ExploreTopCategory)
	route.GET(ec_constant.ROUTE_PROFILE, userdata.Profile)
	route.GET(ec_constant.ROUTE_SUBSCRIPTION, response.Subscription)
	route.GET(ec_constant.ROUTE_ONBOARDING, response.OnBoarding)

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
	// Profile
	route.Static(ec_constant.ProfileRelativePath4x, ec_constant.ProfileFolderPath4x)

	//OnBoarding
	route.Static(ec_constant.OnBoardingRelativePath4x, ec_constant.OnBoardingFolderPath4x)

	// User Related Details
	route.POST("/user", userData.GetUserInfo)
	route.POST("/address", userData.GetUserAddress)
	route.POST("/cart_item", userData.GetUserCartItem)
	route.POST("/subscription", userData.GetUserAllSubscription)
	route.POST("/c_wishlist", userData.CreateWishlist)
	route.POST("/r_wishlist", userData.GetWishlist)
	route.DELETE("/d_wishlist", userData.DeleteWishlist)

	err := route.Run(ec_constant.IpAddress)
	if err != nil {
		return
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}
}
