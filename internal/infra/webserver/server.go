package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	routers.AddCategoryRouter(router, categoryController)
	routers.AddProductRouter(router, productController)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.Run(":" + config.ApiPort)
}
