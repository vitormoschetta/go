package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormoschetta/go/internal/infra/webserver/controllers"
)

func AddProductRouter(router *gin.Engine, c *controllers.ProductController) {
	router.GET("/products", c.GetAll)
	router.GET("/products/:id", c.Get)
	router.GET("/products/category/:category_id", c.GetByCategory)
	router.POST("/products", c.Post)
	router.PUT("/products/:id", c.Put)
	router.DELETE("/products/:id", c.Delete)
	router.POST("/products/promotion", c.PutPromotion)
	router.POST("/products/promotion_by_category", c.PutPromotionbyCategory)
}
