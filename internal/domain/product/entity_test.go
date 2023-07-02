package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitormoschetta/go/internal/domain/category"
)

func TestCreateProduct_Success(t *testing.T) {
	// When | Arrange
	category := category.NewCategory("Category 1")

	name := "Product 1"
	price := 0.01
	entity := NewProduct(name, price, category)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
	assert.NotNil(t, entity)
	assert.Equal(t, name, entity.Name)
	assert.Equal(t, price, entity.Price)
	assert.Equal(t, category, entity.Category)
}

func TestCreateProduct_Fail(t *testing.T) {
	// When | Arrange
	category := category.NewCategory("Category 1")

	name := ""
	price := 0.00
	entity := NewProduct(name, price, category)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 2, len(errs))
	assert.NotNil(t, entity)
	assert.Equal(t, name, entity.Name)
	assert.Equal(t, price, entity.Price)
	assert.Equal(t, category, entity.Category)
}

func TestUpdateProduct_Success(t *testing.T) {
	// When | Arrange
	category := category.NewCategory("Category 1")

	name := "Product 1"
	price := 0.01
	entity := NewProduct(name, price, category)

	newName := "Product 2"
	newPrice := 0.02
	category.Update("Category 2")
	entity.Update(newName, newPrice, category)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 0, len(errs))
	assert.NotNil(t, entity)
	assert.Equal(t, newName, entity.Name)
	assert.Equal(t, newPrice, entity.Price)
	assert.Equal(t, category, entity.Category)
}

func TestUpdateProduct_Fail(t *testing.T) {
	// When | Arrange
	category := category.NewCategory("Category 1")

	name := "Product 1"
	price := 0.01
	entity := NewProduct(name, price, category)

	newName := ""
	newPrice := 0.00
	category.Update("Category 2")
	entity.Update(newName, newPrice, category)

	// Given | Act
	errs := entity.Validate()

	// Then | Assert
	assert.Equal(t, 2, len(errs))
	assert.NotNil(t, entity)
	assert.Equal(t, newName, entity.Name)
	assert.Equal(t, newPrice, entity.Price)
	assert.Equal(t, category, entity.Category)
}
