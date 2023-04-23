package webserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vitormoschetta/go/config"
	"github.com/vitormoschetta/go/internal/application/useCases"
	"github.com/vitormoschetta/go/internal/infra/database"
	"github.com/vitormoschetta/go/internal/infra/database/repositories"
	"github.com/vitormoschetta/go/internal/infra/webserver/controllers"
	"github.com/vitormoschetta/go/internal/infra/webserver/routers"
)

func StartServer() {
	config.Load()
	db := database.ConnectDB()

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryUseCase := useCases.NewCategoryUseCase(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryRepository, categoryUseCase)

	productRepository := repositories.NewProductRepository(db)
	productUseCase := useCases.NewProductUseCase(productRepository, categoryRepository)
	productController := controllers.NewProductController(productRepository, productUseCase)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		categories := v1.Group("/categories")
		{
			routers.AddCategoryRouter2(categories, categoryController)
		}
		products := v1.Group("/products")
		{
			routers.AddProductRouter2(products, productController)
		}
		health := v1.Group("/health")
		{
			health.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"status": "ok"})
			})
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + config.ApiPort)
}
