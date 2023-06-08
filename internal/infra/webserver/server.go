package webserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/vitormoschetta/go/docs"
	categoryApplication "github.com/vitormoschetta/go/internal/application/category"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/infra/config"
	"github.com/vitormoschetta/go/internal/infra/database"
	"github.com/vitormoschetta/go/internal/infra/database/repositories"
	"github.com/vitormoschetta/go/internal/infra/webserver/controllers"
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

	router.HandleFunc("/api/v1/categories", categoryController.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/categories", categoryController.Post).Methods("POST")
	router.HandleFunc("/api/v1/categories/{id}", categoryController.Get).Methods("GET")
	router.HandleFunc("/api/v1/categories/{id}", categoryController.Put).Methods("PUT")
	router.HandleFunc("/api/v1/categories/{id}", categoryController.Delete).Methods("DELETE")

	router.HandleFunc("/api/v1/products", productController.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/products", productController.Post).Methods("POST")
	router.HandleFunc("/api/v1/products/{id}", productController.Get).Methods("GET")
	router.HandleFunc("/api/v1/products/{id}", productController.Put).Methods("PUT")
	router.HandleFunc("/api/v1/products/{id}", productController.Delete).Methods("DELETE")
	router.HandleFunc("/api/v1/promotion", productController.PutPromotion).Methods("PUT")
	router.HandleFunc("/api/v1/promotion_by_category", productController.PutPromotionbyCategory).Methods("PUT")

	// http.HandleFunc("/api/v1/categories", categoryController.GetAll)
	// http.HandleFunc("/api/v1/categories", categoryController.Post)
	// http.HandleFunc("/api/v1/categories/:id", categoryController.Get)
	// http.HandleFunc("/api/v1/categories/:id", categoryController.Put)
	// http.HandleFunc("/api/v1/categories/:id", categoryController.Delete)

	// http.HandleFunc("/api/v1/products", productController.GetAll)
	// http.HandleFunc("/api/v1/products", productController.Post)
	// http.HandleFunc("/api/v1/products/:id", productController.Get)
	// http.HandleFunc("/api/v1/products/:id", productController.Put)
	// http.HandleFunc("/api/v1/products/:id", productController.Delete)
	// http.HandleFunc("/api/v1/promotion", productController.PutPromotion)
	// http.HandleFunc("/api/v1/promotion_by_category", productController.PutPromotionbyCategory)

	port := appConfig.Port
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port", port)
	http.ListenAndServe(":"+port, router)
}
