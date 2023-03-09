package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormoschetta/go/internal/infra/httpServer/controllers"
)

func AddProductRouter(router *gin.Engine, c *controllers.ProductController) {
	router.GET("/products", c.GetProducts)
	router.GET("/products/:id", c.GetProduct)
	router.POST("/products", c.PostProduct)
	router.PUT("/products/:id", c.PutProduct)
	router.DELETE("/products/:id", c.DeleteProduct)
}
