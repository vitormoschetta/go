package product

import "github.com/vitormoschetta/go/internal/application/general"

type UpdateProductInput struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

func (p *UpdateProductInput) Validate() (output general.Output) {
	if p.ID == "" {
		output.Errors = append(output.Errors, "ID is required")
	}
	if p.Name == "" {
		output.Errors = append(output.Errors, "Name is required")
	}
	if p.Price <= 0 {
		output.Errors = append(output.Errors, "Price is less than or equal to zero")
	}
	return
}
