package webserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vitormoschetta/go/docs"
	categoryApplication "github.com/vitormoschetta/go/internal/application/category"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/infra/config"
	"github.com/vitormoschetta/go/internal/infra/database"
	"github.com/vitormoschetta/go/internal/infra/database/repositories"
	"github.com/vitormoschetta/go/internal/infra/webserver/controllers"
	"github.com/vitormoschetta/go/internal/infra/webserver/routers"
)

func Start() {
	appConfig := config.Load()
	db := database.ConnectDB(appConfig)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryUseCase := categoryApplication.NewCategoryUseCase(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryRepository, categoryUseCase)

	productRepository := repositories.NewProductRepository(db)
	productUseCase := productApplication.NewProductUseCase(productRepository, categoryRepository)
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

	err := router.Run(":" + appConfig.Port)
	if err != nil {
		panic(err)
	}
}
