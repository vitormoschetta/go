package httpServer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitormoschetta/go/config"
	"github.com/vitormoschetta/go/internal/application/useCases"
	"github.com/vitormoschetta/go/internal/infra/database"
	"github.com/vitormoschetta/go/internal/infra/database/repositories"
	"github.com/vitormoschetta/go/internal/infra/httpServer/controllers"
	"github.com/vitormoschetta/go/internal/infra/httpServer/routers"
)

func StartServer() {
	config.Load()
	db := database.ConnectDB()

	repository := repositories.NewProductRepository(db)
	useCase := useCases.NewProductUseCase(repository)
	controller := controllers.NewProductController(repository, useCase)

	router := gin.Default()
	routers.AddProductRouter(router, controller)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	router.Run(":" + config.ApiPort)
}
