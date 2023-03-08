package response

import (
	"E-Commerce/ec_constant"
	"E-Commerce/response/chunks/cashback"
	"E-Commerce/response/chunks/deal_of_the_week"
	"E-Commerce/response/chunks/explore_category"
	"E-Commerce/response/chunks/featured_items"
	"E-Commerce/response/chunks/top_categogies"
	"E-Commerce/response/chunks/top_products"
	"E-Commerce/response/chunks/top_viewpager"
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
	} else if queryKey == ec_constant.QUERY_VALUE_TOPPRODUCT_EXPLORE {
		parentChildren = top_categogies.ExploreGroceries(parentChildren)
		parentChildren = top_categogies.ExploreVegetables(parentChildren)
		parentChildren = top_categogies.ExploreFruits(parentChildren)
		parentChildren = top_categogies.ExploreDairyProduct(parentChildren)
		parentChildren = top_categogies.ExploreBakeryProduct(parentChildren)
		c.JSON(http.StatusOK, gin.H{"data": parentChildren})
	} else if queryKey == ec_constant.QUERY_VALUE_DEAL_OF_WEEK {
		parentChildren = top_categogies.ExploreGroceries(parentChildren)
		parentChildren = top_categogies.ExploreVegetables(parentChildren)
		parentChildren = top_categogies.ExploreFruits(parentChildren)
		parentChildren = top_categogies.ExploreDairyProduct(parentChildren)
		parentChildren = top_categogies.ExploreBakeryProduct(parentChildren)
		c.JSON(http.StatusOK, gin.H{"data": parentChildren})
	} else if queryKey == ec_constant.QUERY_VALUE_FEATURED_ITEM {
		parentChildren = top_categogies.ExploreGroceries(parentChildren)
		parentChildren = top_categogies.ExploreVegetables(parentChildren)
		parentChildren = top_categogies.ExploreFruits(parentChildren)
		parentChildren = top_categogies.ExploreDairyProduct(parentChildren)
		parentChildren = top_categogies.ExploreBakeryProduct(parentChildren)
		c.JSON(http.StatusOK, gin.H{"data": parentChildren})
	}
}

func ExploreCategory(c *gin.Context) {
	var parentChildren []interface{}
	parentChildren = explore_category.ExploreCategory(parentChildren)
	c.JSON(http.StatusOK, gin.H{"data": parentChildren})
}
