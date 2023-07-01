package product

import (
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type CreateProductInput struct {
	Name       string
	Price      float64
	CategoryId string
}

func NewCreateProductInput(name string, price float64, categoryId string) *CreateProductInput {
	return &CreateProductInput{
		Name:       name,
		Price:      price,
		CategoryId: categoryId,
	}
}

func (p *CreateProductInput) Validate() (errs []string) {
	if p.Name == "" {
		errs = append(errs, "Name is required")
	}
	if p.Price <= 0 {
		errs = append(errs, "Price is less than or equal to zero")
	}
	if p.CategoryId == "" {
		errs = append(errs, "CategoryId is required")
	}
	return errs
}

func (p *CreateProductInput) ToEntity(category category.Category) product.Product {
	return product.NewProduct(p.Name, p.Price, category)
}
