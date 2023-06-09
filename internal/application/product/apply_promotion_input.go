package product

type ApplyPromotionProductInput struct {
	ProductId  string  `json:"id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
}

func (p *ApplyPromotionProductInput) Validate() (errors []string) {
	if p.ProductId == "" {
		errors = append(errors, "Product is required")
	}
	if p.Percentage <= 0 {
		errors = append(errors, "Percentage is less than or equal to zero")
	}
	return
}
