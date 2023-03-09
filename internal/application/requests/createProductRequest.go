package requests

import (
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *CreateProductRequest) Validate() (response models.Response) {
	if p.Name == "" {
		response.Errors = append(response.Errors, "Name is required")
	}
	if p.Price <= 0 {
		response.Errors = append(response.Errors, "Price is less than or equal to zero")
	}
	return
}

func (p *CreateProductRequest) ToProductModel() models.Product {
	return models.NewProduct(p.Name, p.Price)
}
