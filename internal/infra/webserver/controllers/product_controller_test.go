package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	"github.com/vitormoschetta/go/internal/application/common"
	applicationProduct "github.com/vitormoschetta/go/internal/application/product"
	domainProduct "github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/internal/share/middlewares"
	"github.com/vitormoschetta/go/mock"
)

type ProductControllerTest struct {
	suite.Suite
	productController *ProductController
}

func TestSuiteStart(t *testing.T) {
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
	router.HandleFunc("/", suite.productController.GetAll).Methods("GET")

	// Act
	router.ServeHTTP(recorder, req)
	var products []domainProduct.Product
	errUnmarshal := json.Unmarshal([]byte(recorder.Body.String()), &products)
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
	router.Use(middlewares.TracingMiddleware)
	router.HandleFunc("/", suite.productController.Post).Methods("POST")

	// Act
	router.ServeHTTP(recorder, req)

	var output common.Output
	errUnmarshal := json.Unmarshal([]byte(recorder.Body.String()), &output)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal output")
	}

	// Assert
	suite.Assert().Equal(http.StatusNotFound, recorder.Code)
	suite.Assert().Len(output.Errors, 1)
	suite.Assert().Equal("Category not found", output.Errors[0])
}
