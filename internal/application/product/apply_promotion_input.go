package product

import "github.com/vitormoschetta/go/internal/domain/models"

type ApplyPromotionProductInput struct {
	ProductId  string  `json:"id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

func (p *ApplyPromotionProductInput) Validate() (response models.Response) {
	if p.ProductId == "" {
		response.Errors = append(response.Errors, "Product is required")
	}
	if p.Percentage <= 0 {
		response.Errors = append(response.Errors, "Percentage is less than or equal to zero")
	}
	return
}
