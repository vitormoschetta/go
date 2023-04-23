package requests

import "github.com/vitormoschetta/go/internal/domain/models"

type ApplyPromotionProductByCategoryRequest struct {
	CategoryId string  `json:"category_id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

func (p *ApplyPromotionProductByCategoryRequest) Validate() (response models.Response) {
	if p.CategoryId == "" {
		response.Errors = append(response.Errors, "Category is required")
	}
	if p.Percentage <= 0 {
		response.Errors = append(response.Errors, "Percentage is less than or equal to zero")
	}
	return
}
