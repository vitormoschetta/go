package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	applicationProduct "github.com/vitormoschetta/go/internal/application/product"
	domainProduct "github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/internal/infra/api/responses"
	"github.com/vitormoschetta/go/mock"
	"github.com/vitormoschetta/go/pkg/middlewares"
)

type ProductControllerTest struct {
	suite.Suite
	productController *ProductController
}

func TestProductSuiteStart(t *testing.T) {
	suite.Run(t, new(ProductControllerTest))
}

func (suite *ProductControllerTest) SetupTest() {
	productRepository := mock.NewProductRepositoryFake()
	categoryRepository := mock.NewCategoryRepositoryFake()
	productUseCase := applicationProduct.NewProductUseCase(productRepository, categoryRepository)
	productController := NewProductController(productRepository, productUseCase)
	suite.productController = productController
}

func (suite *ProductControllerTest) TestGetAll() {
	// Arrange
	suite.SetupTest()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/", suite.productController.GetAll).Methods("GET")

	// Act
	router.ServeHTTP(recorder, req)
	var products []domainProduct.Product
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &products)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal products")
	}

	// Assert
	suite.Equal(http.StatusOK, recorder.Code)
	suite.Equal(0, len(products))
}

func (suite *ProductControllerTest) TestGet() {
	// Arrange
	suite.SetupTest()
	req, err := http.NewRequest("GET", "/123", nil)
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/{id}", suite.productController.Get).Methods("GET")

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	suite.Equal(http.StatusNotFound, recorder.Code)
}

func (suite *ProductControllerTest) TestPost() {
	// Arrange
	suite.SetupTest()
	input := applicationProduct.CreateProductInput{
		Name:       "Product 1",
		Price:      10.0,
		CategoryId: "123",
	}
	jsonData, err := json.Marshal(input)
	if err != nil {
		suite.Fail("Error marshal item")
	}
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonData))
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/", suite.productController.Post).Methods("POST")

	// Act
	router.ServeHTTP(recorder, req)

	var response responses.Response
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &response)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal output")
	}

	// Assert
	suite.Assert().Equal(http.StatusNotFound, recorder.Code)
	suite.Assert().Len(response.Errors, 1)
	suite.Assert().Equal("Category not found", response.Errors[0])
}
