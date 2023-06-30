package product

type ApplyPromotionProductByCategoryInput struct {
	CategoryId string
	Percentage float64
}

func (p *ApplyPromotionProductByCategoryInput) Validate() (errs []string) {
	if p.CategoryId == "" {
		errs = append(errs, "CategoryId is required")
	}
	if p.Percentage <= 0 {
		errs = append(errs, "Percentage is less than or equal to zero")
	}
	return errs
}
