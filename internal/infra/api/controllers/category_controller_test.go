package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
	applicationCategory "github.com/vitormoschetta/go/internal/application/category"
	domainCategory "github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/infra/api/requests"
	"github.com/vitormoschetta/go/internal/infra/api/responses"
	"github.com/vitormoschetta/go/mock"
	"github.com/vitormoschetta/go/pkg/middlewares"
)

type CategoryControllerTest struct {
	suite.Suite
	CategoryController *CategoryController
}

func (suite *CategoryControllerTest) SetupTest() {
	categoryRepository := mock.NewCategoryRepositoryFake()
	categoryRepository.Seed(BuildCategories())
	categoryUseCase := applicationCategory.NewCategoryUseCase(categoryRepository)
	CategoryController := NewCategoryController(categoryRepository, categoryUseCase)
	suite.CategoryController = CategoryController
}

func (suite *CategoryControllerTest) TestGetAll_Ok() {
	// Arrange
	suite.SetupTest()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/", suite.CategoryController.GetAll).Methods(http.MethodGet)

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	var categories []domainCategory.Category
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &categories)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal Categorys")
	}

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().Len(categories, 2)
}

func (suite *CategoryControllerTest) TestGet_Ok() {
	// Arrange
	suite.SetupTest()
	req, err := http.NewRequest(http.MethodGet, "/123", nil)
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/{id}", suite.CategoryController.Get).Methods(http.MethodGet)

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &response)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal Category")
	}

	var category domainCategory.Category
	errUnmarshal = json.Unmarshal(recorder.Body.Bytes(), &category)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal Category")
	}

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().NotNil(category)
	suite.Assert().Equal("123", category.ID)
}

func (suite *CategoryControllerTest) TestGet_NotFound() {
	// Arrange
	suite.SetupTest()
	req, err := http.NewRequest(http.MethodGet, "/1234", nil)
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/{id}", suite.CategoryController.Get).Methods(http.MethodGet)

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &response)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal Category")
	}

	suite.Assert().Equal(http.StatusNotFound, recorder.Code)
	suite.Assert().NotNil(response)
}

func (suite *CategoryControllerTest) TestPost() {
	// Arrange
	suite.SetupTest()
	request := requests.CreateCategoryRequest{
		Name: "Category 1",
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
	router.HandleFunc("/", suite.CategoryController.Post).Methods(http.MethodPost)

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &response)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal output")
	}

	suite.Assert().Equal(http.StatusCreated, recorder.Code)
	suite.Assert().NotNil(response)
	suite.Assert().NotNil(response.Data)
	suite.Assert().Len(response.Errors, 0)
}

func (suite *CategoryControllerTest) TestPut() {
	// Arrange
	suite.SetupTest()
	request := requests.UpdateCategoryRequest{
		ID:   "123",
		Name: "Category 1",
	}
	jsonData, err := json.Marshal(request)
	if err != nil {
		suite.Fail("Error marshal item")
	}
	req, err := http.NewRequest(http.MethodPut, "/123", bytes.NewBuffer(jsonData))
	if err != nil {
		suite.Fail("Error creating request")
	}
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Use(middlewares.Tracing)
	router.HandleFunc("/{id}", suite.CategoryController.Put).Methods(http.MethodPut)

	// Act
	router.ServeHTTP(recorder, req)

	// Assert
	var response responses.Response
	errUnmarshal := json.Unmarshal(recorder.Body.Bytes(), &response)
	if errUnmarshal != nil {
		suite.Fail("Error unmarshal output")
	}

	suite.Assert().Equal(http.StatusOK, recorder.Code)
	suite.Assert().NotNil(recorder.Body)
}

func TestCategorySuiteStart(t *testing.T) {
	suite.Run(t, new(CategoryControllerTest))
}

func BuildCategories() []domainCategory.Category {
	return []domainCategory.Category{
		{
			ID:   "123",
			Name: "Category 1",
		},
		{
			ID:   "456",
			Name: "Category 2",
		},
	}
}
