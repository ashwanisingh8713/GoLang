package response

import (
	"ServerDrivenUI/src/E-Commerce/ec_constant"
	"ServerDrivenUI/src/E-Commerce/response/chunks/cashback"
	"ServerDrivenUI/src/E-Commerce/response/chunks/deal_of_the_week"
	"ServerDrivenUI/src/E-Commerce/response/chunks/explore_category"
	"ServerDrivenUI/src/E-Commerce/response/chunks/featured_items"
	"ServerDrivenUI/src/E-Commerce/response/chunks/top_categogies"
	"ServerDrivenUI/src/E-Commerce/response/chunks/top_products"
	"ServerDrivenUI/src/E-Commerce/response/chunks/top_viewpager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	var parentChildren []interface{}
	parentChildren = top_viewpager.TopViewpager(parentChildren)
	parentChildren = top_categogies.TopCategories(parentChildren)
	parentChildren = top_products.TopProducts(parentChildren)
	parentChildren = cashback.Cashback(parentChildren)
	parentChildren = deal_of_the_week.DealOfTheWeek(parentChildren)
	parentChildren = featured_items.FeaturedItems(parentChildren)
	c.JSON(http.StatusOK, gin.H{"data": parentChildren})
}

func ExploreTopCategory(c *gin.Context) {
	var parentChildren []interface{}

	var queryKey = c.Query(ec_constant.QUERY_KEY_CATEGORY)
	if queryKey == ec_constant.QUERY_VALUE_TOPCATEGORY {
		parentChildren = top_categogies.ExploreGroceries(parentChildren)
		parentChildren = top_categogies.ExploreVegetables(parentChildren)
		parentChildren = top_categogies.ExploreFruits(parentChildren)
		parentChildren = top_categogies.ExploreDairyProduct(parentChildren)
		parentChildren = top_categogies.ExploreBakeryProduct(parentChildren)
		c.JSON(http.StatusOK, gin.H{"data": parentChildren})
	} else if queryKey == ec_constant.QUERY_VALUE_TOPPRODUCT {
		var parentChildren []interface{}
		parentChildren = explore_category.ExploreCategory(parentChildren)
		c.JSON(http.StatusOK, gin.H{"data": parentChildren})
	} else {

	}

}

func ExploreCategory(c *gin.Context) {
	var parentChildren []interface{}
	parentChildren = explore_category.ExploreCategory(parentChildren)
	c.JSON(http.StatusOK, gin.H{"data": parentChildren})
}
