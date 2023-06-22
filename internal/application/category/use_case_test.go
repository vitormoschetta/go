package category

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoschetta/go/internal/application/common"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/share/middlewares"
	"github.com/vitormoschetta/go/mock"
)

func Test_With_Category_Add_With_Valid_Data(t *testing.T) {
	// Arrange
	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.CorrelationKey, uuid.New().String())
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	input := CreateCategoryInput{
		Name: "Category 1",
	}
	// Act
	output := useCase.Create(ctx, input)
	// Assert
	assert.Equal(t, common.DomainCodeSuccess, output.Code)
	assert.Len(t, output.Errors, 0)
}

func Test_With_Category_Add_With_Invalid_Name(t *testing.T) {
	// Arrange
	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.CorrelationKey, uuid.New().String())
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	input := CreateCategoryInput{
		Name: "",
	}
	// Act
	output := useCase.Create(ctx, input)
	// Assert
	assert.Equal(t, common.DomainCodeInvalidInput, output.Code)
	assert.NotNil(t, output.Errors)
	assert.Len(t, output.Errors, 1)
}

func Test_With_Category_Add_With_Database_Error(t *testing.T) {
	// Arrange
	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.CorrelationKey, uuid.New().String())
	repository := mock.NewCategoryRepositoryFake()
	repository.SaveError = true
	useCase := NewCategoryUseCase(repository)
	input := CreateCategoryInput{
		Name: "Category 1",
	}
	// Act
	output := useCase.Create(ctx, input)
	// Assert
	assert.Equal(t, common.DomainCodeInternalError, output.Code)
	assert.NotNil(t, output.Errors)
	assert.Equal(t, 1, len(output.Errors))
}

func Test_With_Category_Update_With_Valid_Data(t *testing.T) {
	// Arrange
	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.CorrelationKey, uuid.New().String())
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	input := CreateCategoryInput{
		Name: "Category 1",
	}
	output := useCase.Create(ctx, input)
	if output.Code != common.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", common.DomainCodeSuccess, output.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   output.Data.(category.Category).ID,
		Name: "Category 2",
	}
	output = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, common.DomainCodeSuccess, output.Code)
	assert.Len(t, output.Errors, 0)
	assert.Equal(t, "Category 2", output.Data.(category.Category).Name)
}

func Test_With_Category_Update_With_Invalid_ID(t *testing.T) {
	// Arrange
	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.CorrelationKey, uuid.New().String())
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	input := CreateCategoryInput{
		Name: "Category 1",
	}
	output := useCase.Create(ctx, input)
	if output.Code != common.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", common.DomainCodeSuccess, output.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   uuid.NewString(),
		Name: "Category 2",
	}
	output = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, common.DomainCodeNotFound, output.Code)
	assert.NotNil(t, output.Errors)
	assert.Len(t, output.Errors, 1)
}

func Test_With_Category_Update_With_ID_Empty(t *testing.T) {
	// Arrange
	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.CorrelationKey, uuid.New().String())
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	input := CreateCategoryInput{
		Name: "Category 1",
	}
	output := useCase.Create(ctx, input)
	if output.Code != common.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", common.DomainCodeSuccess, output.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   "",
		Name: "Category 2",
	}
	output = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, common.DomainCodeInvalidInput, output.Code)
	assert.NotNil(t, output.Errors)
	assert.Len(t, output.Errors, 1)
}

func Test_With_Category_Update_With_Invalid_Name(t *testing.T) {
	// Arrange
	ctx := context.Background()
	ctx = context.WithValue(ctx, middlewares.CorrelationKey, uuid.New().String())
	repository := mock.NewCategoryRepositoryFake()
	useCase := NewCategoryUseCase(repository)
	input := CreateCategoryInput{
		Name: "Category 1",
	}
	output := useCase.Create(ctx, input)
	if output.Code != common.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", common.DomainCodeSuccess, output.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   uuid.NewString(),
		Name: "",
	}
	output = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, common.DomainCodeInvalidInput, output.Code)
	assert.NotNil(t, output.Errors)
	assert.Len(t, output.Errors, 1)
}
