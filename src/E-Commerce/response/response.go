package response

import (
	"ServerDrivenUI/src/E-Commerce/response/chunks/deal_of_the_week"
	"ServerDrivenUI/src/E-Commerce/response/chunks/top_categogies"
	"ServerDrivenUI/src/E-Commerce/response/chunks/top_products"
	"ServerDrivenUI/src/E-Commerce/response/chunks/top_viewpager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context) {
	var parentChildren []interface{}
	parentChildren = top_viewpager.TopViewpager(parentChildren)
	parentChildren = top_categogies.TopCategories(parentChildren)
	parentChildren = top_products.TopProducts(parentChildren)
	parentChildren = deal_of_the_week.DealOfTheWeek(parentChildren)
	c.JSON(http.StatusOK, gin.H{"data": parentChildren})
}
