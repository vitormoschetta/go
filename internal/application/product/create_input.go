package product

import (
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type CreateProductInput struct {
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	CategoryId string   `json:"category_id"`
	Errors     []string `json:"-"`
}

func (p *CreateProductInput) IsInvalid() bool {
	p.validate()
	return len(p.Errors) > 0
}

func (p *CreateProductInput) validate() {
	if p.Name == "" {
		p.Errors = append(p.Errors, "Name is required")
	}
	if p.Price <= 0 {
		p.Errors = append(p.Errors, "Price is less than or equal to zero")
	}
	if p.CategoryId == "" {
		p.Errors = append(p.Errors, "Category is required")
	}
}

func (p *CreateProductInput) ToEntity(category category.Category) product.Product {
	return product.NewProduct(p.Name, p.Price, category)
}
