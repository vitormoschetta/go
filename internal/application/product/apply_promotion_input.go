package product

type ApplyPromotionProductInput struct {
	ProductId  string  `json:"id" binding:"required"`
	Percentage float64 `json:"percentage" binding:"required"`
	Errors     []string
}

func (p *ApplyPromotionProductInput) IsInvalid() bool {
	p.validate()
	return len(p.Errors) > 0
}

func (p *ApplyPromotionProductInput) validate() {
	if p.ProductId == "" {
		p.Errors = append(p.Errors, "Product is required")
	}
	if p.Percentage <= 0 {
		p.Errors = append(p.Errors, "Percentage is less than or equal to zero")
	}
	return
}
