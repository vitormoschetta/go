package category

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vitormoschetta/go/mock"
	"github.com/vitormoschetta/go/pkg/middlewares"
	"github.com/vitormoschetta/go/pkg/output"
)

func TestCategoryUseCase_Create_ValidData(t *testing.T) {
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
	assert.Equal(t, output.DomainCodeSuccess, out.GetCode())
	assert.Len(t, out.GetErrors(), 0)
}

func TestCategoryUseCase_Create_InvalidData(t *testing.T) {
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
	assert.Equal(t, output.DomainCodeInvalidInput, out.GetCode())
	assert.NotNil(t, out.GetErrors())
	assert.Len(t, out.GetErrors(), 1)
}

func TestCategoryUseCase_Create_DatabaseError(t *testing.T) {
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
	assert.Equal(t, output.DomainCodeInternalError, out.GetCode())
	assert.NotNil(t, out.GetErrors())
	assert.Equal(t, 1, len(out.GetErrors()))
}
