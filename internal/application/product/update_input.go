package product

type UpdateProductInput struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId string  `json:"category_id"`
}

func (p *UpdateProductInput) Validate() (errors []string) {
	if p.ID == "" {
		errors = append(errors, "ID is required")
	}
	if p.Name == "" {
		errors = append(errors, "Name is required")
	}
	if p.Price <= 0 {
		errors = append(errors, "Price is less than or equal to zero")
	}
	return
}
