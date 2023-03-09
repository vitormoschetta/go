package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoschetta/go/internal/application/requests"
	"github.com/vitormoschetta/go/internal/application/useCases"
	"github.com/vitormoschetta/go/internal/domain/models"
	"github.com/vitormoschetta/go/tests/mock"
)

func Test_With_Product_Add_With_Valid_Data(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "Product 1",
		Price: 10,
	}
	// Act
	response, statusCode := useCase.Save(request)
	// Assert
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, response.Errors)
}

func Test_With_Product_Add_With_Invalid_Name(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "",
		Price: 10,
	}
	// Act
	response, statusCode := useCase.Save(request)
	// Assert
	assert.Equal(t, 400, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}

func Test_With_Product_Add_With_Invalid_Price(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "Product 1",
		Price: 0,
	}
	// Act
	response, statusCode := useCase.Save(request)
	// Assert
	assert.Equal(t, 400, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}

func Test_With_Product_Add_With_Invalid_Name_And_Price(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "",
		Price: 0,
	}
	// Act
	response, statusCode := useCase.Save(request)
	// Assert
	assert.Equal(t, 400, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 2, len(response.Errors))
}

func Test_With_Product_Update_With_Valid_Data(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "Product 1",
		Price: 10,
	}
	response, statusCode := useCase.Save(request)
	if statusCode != 201 {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}
	// Act
	request2 := requests.UpdateProductRequest{
		ID:    response.Data.(models.Product).ID,
		Name:  "Product 2",
		Price: 20,
	}
	response, statusCode = useCase.Update(request2)
	// Assert
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, response.Errors)
	assert.Equal(t, "Product 2", response.Data.(models.Product).Name)
}

func Test_With_Product_Update_With_Invalid_ID(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "Product 1",
		Price: 10,
	}
	response, statusCode := useCase.Save(request)
	if statusCode != 201 {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}
	// Act
	request2 := requests.UpdateProductRequest{
		ID:    uuid.New().String(),
		Name:  "Product 2",
		Price: 20,
	}
	response, statusCode = useCase.Update(request2)
	// Assert
	assert.Equal(t, 404, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}

func Test_With_Product_Update_With_Invalid_Name(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "Product 1",
		Price: 10,
	}
	response, statusCode := useCase.Save(request)
	if statusCode != 201 {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}
	// Act
	request2 := requests.UpdateProductRequest{
		ID:    response.Data.(models.Product).ID,
		Name:  "",
		Price: 20,
	}
	response, statusCode = useCase.Update(request2)
	// Assert
	assert.Equal(t, 400, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}

func Test_With_Product_Update_With_Invalid_Price(t *testing.T) {
	// Arrange
	repository := mock.NewProductRepositoryFake()
	useCase := useCases.NewProductUseCase(repository)
	request := requests.CreateProductRequest{
		Name:  "Product 1",
		Price: 10,
	}
	response, statusCode := useCase.Save(request)
	if statusCode != 201 {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}
	// Act
	request2 := requests.UpdateProductRequest{
		ID:    response.Data.(models.Product).ID,
		Name:  "Product 2",
		Price: 0,
	}
	response, statusCode = useCase.Update(request2)
	// Assert
	assert.Equal(t, 400, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}
