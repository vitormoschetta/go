package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInput(t *testing.T) {
	// When | Arrange

	input := CreateProductInput{
		Name:       "Product 1",
		Price:      0.01,
		CategoryId: "1",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
}

func TestCreateInput_Fail(t *testing.T) {
	// When | Arrange

	input := CreateProductInput{
		Name:       "",
		Price:      0.00,
		CategoryId: "",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 3, len(errs))
}

func TestUpdateInput(t *testing.T) {
	// When | Arrange

	input := UpdateProductInput{
		ID:         "1",
		Name:       "Product 1",
		Price:      0.01,
		CategoryId: "1",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
}

func TestUpdateInput_Fail(t *testing.T) {
	// When | Arrange

	input := UpdateProductInput{
		ID:         "",
		Name:       "",
		Price:      0.00,
		CategoryId: "",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 4, len(errs))
}

func TestUpdateInput_Fail2(t *testing.T) {
	// When | Arrange

	input := UpdateProductInput{
		ID:         "1",
		Name:       "",
		Price:      0.00,
		CategoryId: "",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 3, len(errs))
}

func TestUpdateInput_Fail3(t *testing.T) {
	// When | Arrange

	input := UpdateProductInput{
		ID:         "",
		Name:       "Product 1",
		Price:      0.00,
		CategoryId: "",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 3, len(errs))
}
