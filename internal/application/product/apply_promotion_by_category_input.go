package product

import "github.com/vitormoschetta/go/internal/application/general"

type ApplyPromotionProductByCategoryInput struct {
	CategoryId string  `json:"category_id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

func (p *ApplyPromotionProductByCategoryInput) Validate() (output general.Output) {
	if p.CategoryId == "" {
		output.Errors = append(output.Errors, "Category is required")
	}
	if p.Percentage <= 0 {
		output.Errors = append(output.Errors, "Percentage is less than or equal to zero")
	}
	return
}
