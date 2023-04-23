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

func AddProductRouter2(router *gin.RouterGroup, c *controllers.ProductController) {
	router.GET("", c.GetAll)
	router.GET("/:id", c.Get)
	router.GET("/category/:category_id", c.GetByCategory)
	router.POST("", c.Post)
	router.PUT("/:id", c.Put)
	router.DELETE("/:id", c.Delete)
	router.POST("/promotion", c.PutPromotion)
	router.POST("/promotion_by_category", c.PutPromotionbyCategory)
}
