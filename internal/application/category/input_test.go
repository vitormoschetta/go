package category

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateInput(t *testing.T) {
	// When | Arrange
	input := CreateCategoryInput{
		Name: "Category 1",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
}

func TestCreateInput_Fail(t *testing.T) {
	// When | Arrange
	input := CreateCategoryInput{
		Name: "",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
}

func TestUpdateInput(t *testing.T) {
	// When | Arrange
	input := UpdateCategoryInput{
		ID:   "1",
		Name: "Category 1",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
}

func TestUpdateInput_Fail(t *testing.T) {
	// When | Arrange
	input := UpdateCategoryInput{
		ID:   "",
		Name: "",
	}

	// Given | Act
	errs := input.Validate()

	// Then | Assert
	assert.Equal(t, 2, len(errs))
}
