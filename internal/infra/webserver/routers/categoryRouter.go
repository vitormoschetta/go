package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormoschetta/go/internal/infra/webserver/controllers"
)

func AddCategoryRouter(router *gin.Engine, c *controllers.CategoryController) {
	router.GET("/categories", c.GetAll)
	router.GET("/categories/:id", c.Get)
	router.POST("/categories", c.Post)
	router.PUT("/categories/:id", c.Put)
	router.DELETE("/categories/:id", c.Delete)
}

func AddCategoryRouter2(categories *gin.RouterGroup, c *controllers.CategoryController) {
	categories.GET("", c.GetAll)
	categories.GET("/:id", c.Get)
	categories.POST("", c.Post)
	categories.PUT("/:id", c.Put)
	categories.DELETE("/:id", c.Delete)
}
