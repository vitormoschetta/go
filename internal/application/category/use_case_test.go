package category

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/tests/mock"
)

func Test_With_Category_Add_With_Valid_Data(t *testing.T) {
	// Arrange
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	request := CreateCategoryInput{
		Name: "Category 1",
	}
	// Act
	response, statusCode := useCase.Create(request)
	// Assert
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, response.Errors)
}

func Test_With_Category_Add_With_Invalid_Name(t *testing.T) {
	// Arrange
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	request := CreateCategoryInput{
		Name: "",
	}
	// Act
	response, statusCode := useCase.Create(request)
	// Assert
	assert.Equal(t, 400, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}

func Test_With_Category_Update_With_Valid_Data(t *testing.T) {
	// Arrange
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	request := CreateCategoryInput{
		Name: "Category 1",
	}
	response, statusCode := useCase.Create(request)
	if statusCode != 201 {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}
	// Act
	request2 := UpdateCategoryInput{
		ID:   response.Data.(category.Category).ID,
		Name: "Category 2",
	}
	response, statusCode = useCase.Update(request2)
	// Assert
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, response.Errors)
	assert.Equal(t, "Category 2", response.Data.(category.Category).Name)
}

func Test_With_Category_Update_With_Invalid_ID(t *testing.T) {
	// Arrange
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	request := CreateCategoryInput{
		Name: "Category 1",
	}
	response, statusCode := useCase.Create(request)
	if statusCode != 201 {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}
	// Act
	request2 := UpdateCategoryInput{
		ID:   uuid.NewString(),
		Name: "Category 2",
	}
	response, statusCode = useCase.Update(request2)
	// Assert
	assert.Equal(t, 404, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}

func Test_With_Category_Update_With_Invalid_Name(t *testing.T) {
	// Arrange
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	request := CreateCategoryInput{
		Name: "Category 1",
	}
	response, statusCode := useCase.Create(request)
	if statusCode != 201 {
		t.Errorf("Expected status code 201, got %v", statusCode)
	}
	// Act
	request2 := UpdateCategoryInput{
		ID:   uuid.NewString(),
		Name: "",
	}
	response, statusCode = useCase.Update(request2)
	// Assert
	assert.Equal(t, 400, statusCode)
	assert.NotNil(t, response.Errors)
	assert.Equal(t, 1, len(response.Errors))
}
