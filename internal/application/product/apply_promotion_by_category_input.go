package product

type ApplyPromotionProductByCategoryInput struct {
	CategoryId string  `json:"category_id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
	Errors     []string
}

func (p *ApplyPromotionProductByCategoryInput) IsInvalid() bool {
	p.validate()
	return len(p.Errors) > 0
}

func (p *ApplyPromotionProductByCategoryInput) validate() {
	if p.CategoryId == "" {
		p.Errors = append(p.Errors, "Category is required")
	}
	if p.Percentage <= 0 {
		p.Errors = append(p.Errors, "Percentage is less than or equal to zero")
	}
}
