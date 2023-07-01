package product

type ApplyPromotionProductInput struct {
	ProductId  string
	Percentage float64
}

func NewApplyPromotionProductInput(productId string, percentage float64) *ApplyPromotionProductInput {
	return &ApplyPromotionProductInput{
		ProductId:  productId,
		Percentage: percentage,
	}
}

func (p *ApplyPromotionProductInput) Validate() (errs []string) {
	if p.ProductId == "" {
		errs = append(errs, "ProductId is required")
	}
	if p.Percentage <= 0 {
		errs = append(errs, "Percentage is less than or equal to zero")
	}
	return errs
}
