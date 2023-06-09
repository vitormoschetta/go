package product

import (
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type CreateProductInput struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

func (p *CreateProductInput) Validate() (errors []string) {
	if p.Name == "" {
		errors = append(errors, "Name is required")
	}
	if p.Price <= 0 {
		errors = append(errors, "Price is less than or equal to zero")
	}
	if p.CategoryId == "" {
		errors = append(errors, "Category is required")
	}
	return
}

func (p *CreateProductInput) ToProductModel(category category.Category) product.Product {
	return product.NewProduct(p.Name, p.Price, category)
}
