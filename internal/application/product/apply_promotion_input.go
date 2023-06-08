package product

import "github.com/vitormoschetta/go/internal/application/general"

type ApplyPromotionProductInput struct {
	ProductId  string  `json:"id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

func (p *ApplyPromotionProductInput) Validate() (output general.Output) {
	if p.ProductId == "" {
		output.Errors = append(output.Errors, "Product is required")
	}
	if p.Percentage <= 0 {
		output.Errors = append(output.Errors, "Percentage is less than or equal to zero")
	}
	return
}
