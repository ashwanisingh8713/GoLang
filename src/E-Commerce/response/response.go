package response

import (
	"E-Commerce/ec_constant"
	"E-Commerce/response/chunks/cashback"
	"E-Commerce/response/chunks/deal_of_the_week"
	"E-Commerce/response/chunks/explore_category"
	"E-Commerce/response/chunks/faq"
	"E-Commerce/response/chunks/featured_items"
	"E-Commerce/response/chunks/subscription"
	"E-Commerce/response/chunks/top_categogies"
	"E-Commerce/response/chunks/top_products"
	"E-Commerce/response/chunks/top_viewpager"
	"E-Commerce/response/onboarding"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Home godoc
// @Summary Provides HomePage Views Groups
// @Description To get Home's Views Group
// @Tags VIEW_GROUP
// @Produce json
// @Success 200 {object} ViewGroupMockResponse
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /ec [get]
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

// ExploreTopCategory godoc
// @Summary Explore Click To Action
// @Description To get response of Explore Btn CTA
// @Tags VIEW_GROUP
// @Produce json
// @Success 200 {object} ViewGroupMockResponse
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /explore [get]
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

// Subscription godoc
// @Summary Subscription Page ViewGroup
// @Description To get response of Subscription Page ViewGroup
// @Tags VIEW_GROUP
// @Produce json
// @Success 200 {object} SubscriptionPageMockResponse
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /subscription [get]
func Subscription(c *gin.Context) {
	var parentChildren []interface{}
	parentChildren = subscription.SubscriptionInfo(parentChildren)
	c.JSON(http.StatusOK, gin.H{"data": parentChildren})
}

// OnBoarding godoc
// @Summary OnBoarding Page ViewGroup
// @Description To get response of OnBoarding Page ViewGroup
// @Tags VIEW_GROUP
// @Produce json
// @Success 200 {object} SubscriptionPageMockResponse
// @Failure      400  string Bad Request
// @Failure      404  string Page Not found
// @Failure      500  string Internal Server Error
// @Router /onboarding [get]
func OnBoarding(c *gin.Context) {
	var parentChildren []interface{}
	parentChildren = onboarding.OnBoarding(parentChildren)
	c.JSON(http.StatusOK, gin.H{"data": parentChildren})
}
func Faq(c *gin.Context) {
	var parentChildren = faq.Faq()
	c.JSON(http.StatusOK, gin.H{"data": parentChildren, "hint_text": "Select Category"})
}
