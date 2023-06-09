package product

type ApplyPromotionProductByCategoryInput struct {
	CategoryId string  `json:"category_id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

func (p *ApplyPromotionProductByCategoryInput) Validate() (errors []string) {
	if p.CategoryId == "" {
		errors = append(errors, "Category is required")
	}
	if p.Percentage <= 0 {
		errors = append(errors, "Percentage is less than or equal to zero")
	}
	return
}
