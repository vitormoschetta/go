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
	"github.com/vitormoschetta/go/internal/infra/api/requests"
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
	productRepository.Seed(BuildProducts())
	categoryRepository := mock.NewCategoryRepositoryFake()
	productUseCase := applicationProduct.NewProductUseCase(productRepository, categoryRepository)
	productController := NewProductController(productRepository, productUseCase)
	suite.productController = productController
}

func (suite *ProductControllerTest) TestGetAll_Ok() {
	// Arrange
	suite.SetupTest()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/", suite.productController.GetAll).Methods(http.MethodGet)

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &response)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal Product")
	}

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().NotNil(response)
	suite.Assert().NotNil(response.Data)
	suite.Assert().Len(response.Data, 2)
}

func (suite *ProductControllerTest) TestGet_Ok() {
	// Arrange
	suite.SetupTest()
	req, err := http.NewRequest(http.MethodGet, "/123", nil)
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/{id}", suite.productController.Get).Methods(http.MethodGet)

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &response)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal Product")
	}

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().NotNil(response)
	suite.Assert().NotNil(response.Data)
}

func (suite *ProductControllerTest) TestPost_Ok() {
	// Arrange
	suite.SetupTest()
	request := requests.CreateProductRequest{
		Name:       "Product 1",
		Price:      10.0,
		CategoryID: "123",
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		suite.Fail("Error marshal item")
	}
	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/", suite.productController.Post).Methods(http.MethodPost)

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
}

func BuildProducts() []domainProduct.Product {
	return []domainProduct.Product{
		{
			ID:       "123",
			Name:     "Product 1",
			Price:    10.0,
			Category: BuildCategories()[0],
		},
		{
			ID:       "456",
			Name:     "Product 2",
			Price:    20.0,
			Category: BuildCategories()[1],
		},
	}
}
