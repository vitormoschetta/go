package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/vitormoschetta/go/docs"
	categoryApplication "github.com/vitormoschetta/go/internal/application/category"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/infra/api/controllers"
	"github.com/vitormoschetta/go/internal/infra/config"
	"github.com/vitormoschetta/go/internal/infra/database"
	"github.com/vitormoschetta/go/internal/infra/database/repositories"
	"github.com/vitormoschetta/go/pkg/middlewares"
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

	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.Use(middlewares.ErrorHandling)
	router.Use(middlewares.LoggingHandling)
	router.Use(middlewares.AcceptJSON)

	router.HandleFunc("/api/v1/categories", categoryController.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/categories", categoryController.Post).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/categories/{id}", categoryController.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/categories/{id}", categoryController.Put).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/categories/{id}", categoryController.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/v1/products", productController.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/products", productController.Post).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/products/{id}", productController.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/products/{id}", productController.Put).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/products/{id}", productController.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/promotion", productController.PutPromotion).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/promotion_by_category", productController.PutPromotionbyCategory).Methods(http.MethodPut)

	port := appConfig.Port
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
