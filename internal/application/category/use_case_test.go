package category

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/mock"
	"github.com/vitormoschetta/go/pkg/middlewares"
	"github.com/vitormoschetta/go/pkg/output"
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
	out := useCase.Create(ctx, input)
	// Assert
	assert.Equal(t, output.DomainCodeSuccess, out.Code)
	assert.Len(t, out.Errors, 0)
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
	out := useCase.Create(ctx, input)
	// Assert
	assert.Equal(t, output.DomainCodeInvalidInput, out.Code)
	assert.NotNil(t, out.Errors)
	assert.Len(t, out.Errors, 1)
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
	out := useCase.Create(ctx, input)
	// Assert
	assert.Equal(t, output.DomainCodeInternalError, out.Code)
	assert.NotNil(t, out.Errors)
	assert.Equal(t, 1, len(out.Errors))
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
	out := useCase.Create(ctx, input)
	if out.Code != output.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", output.DomainCodeSuccess, out.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   out.Data.(category.Category).ID,
		Name: "Category 2",
	}
	out = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, output.DomainCodeSuccess, out.Code)
	assert.Len(t, out.Errors, 0)
	assert.Equal(t, "Category 2", out.Data.(category.Category).Name)
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
	out := useCase.Create(ctx, input)
	if out.Code != output.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", output.DomainCodeSuccess, out.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   uuid.NewString(),
		Name: "Category 2",
	}
	out = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, output.DomainCodeNotFound, out.Code)
	assert.NotNil(t, out.Errors)
	assert.Len(t, out.Errors, 1)
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
	out := useCase.Create(ctx, input)
	if out.Code != output.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", output.DomainCodeSuccess, out.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   "",
		Name: "Category 2",
	}
	out = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, output.DomainCodeInvalidInput, out.Code)
	assert.NotNil(t, out.Errors)
	assert.Len(t, out.Errors, 1)
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
	out := useCase.Create(ctx, input)
	if out.Code != output.DomainCodeSuccess {
		t.Errorf("Expected domain code %v, got %v", output.DomainCodeSuccess, out.Code)
	}
	// Act
	input2 := UpdateCategoryInput{
		ID:   uuid.NewString(),
		Name: "",
	}
	out = useCase.Update(ctx, input2)
	// Assert
	assert.Equal(t, output.DomainCodeInvalidInput, out.Code)
	assert.NotNil(t, out.Errors)
	assert.Len(t, out.Errors, 1)
}
