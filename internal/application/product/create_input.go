package product

import (
	"github.com/vitormoschetta/go/internal/application/general"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type CreateProductInput struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

func (p *CreateProductInput) Validate() (output general.Output) {
	if p.Name == "" {
		output.Errors = append(output.Errors, "Name is required")
	}
	if p.Price <= 0 {
		output.Errors = append(output.Errors, "Price is less than or equal to zero")
	}
	if p.CategoryId == "" {
		output.Errors = append(output.Errors, "Category is required")
	}
	return
}

func (p *CreateProductInput) ToProductModel(category category.Category) product.Product {
	return product.NewProduct(p.Name, p.Price, category)
}
