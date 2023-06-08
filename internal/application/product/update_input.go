package product

import "github.com/vitormoschetta/go/internal/domain/models"

type UpdateProductInput struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

func (p *UpdateProductInput) Validate() (response models.Response) {
	if p.ID == "" {
		response.Errors = append(response.Errors, "ID is required")
	}
	if p.Name == "" {
		response.Errors = append(response.Errors, "Name is required")
	}
	if p.Price <= 0 {
		response.Errors = append(response.Errors, "Price is less than or equal to zero")
	}
	return
}
