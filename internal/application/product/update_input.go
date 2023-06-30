package product

type UpdateProductInput struct {
	ID         string
	Name       string
	Price      float64
	CategoryId string
}

func (p *UpdateProductInput) Validate() (errs []string) {
	if p.ID == "" {
		errs = append(errs, "ID is required")
	}
	if p.Name == "" {
		errs = append(errs, "Name is required")
	}
	if p.Price <= 0 {
		errs = append(errs, "Price is less than or equal to zero")
	}
	return errs
}
