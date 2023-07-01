package product

type UpdateProductInput struct {
	ID         string
	Name       string
	Price      float64
	CategoryId string
}

func NewUpdateProductInput(id string, name string, price float64, categoryId string) *UpdateProductInput {
	return &UpdateProductInput{
		ID:         id,
		Name:       name,
		Price:      price,
		CategoryId: categoryId,
	}
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
