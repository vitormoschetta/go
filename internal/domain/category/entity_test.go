package category

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategory_Success(t *testing.T) {
	// When | Arrange
	name := "Category 1"
	entity := NewCategory(name)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
	assert.NotNil(t, entity)
	assert.Equal(t, name, entity.Name)
}

func TestCreateCategory_Fail(t *testing.T) {
	// When | Arrange
	name := ""
	entity := NewCategory(name)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 1, len(errs))
	assert.NotNil(t, entity)
	assert.Equal(t, name, entity.Name)
}
