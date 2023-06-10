package product

type UpdateProductInput struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	CategoryId string   `json:"category_id"`
	Errors     []string `json:"-"`
}

func (p *UpdateProductInput) IsInvalid() bool {
	p.validate()
	return len(p.Errors) > 0
}

func (p *UpdateProductInput) validate() {
	if p.ID == "" {
		p.Errors = append(p.Errors, "ID is required")
	}
	if p.Name == "" {
		p.Errors = append(p.Errors, "Name is required")
	}
	if p.Price <= 0 {
		p.Errors = append(p.Errors, "Price is less than or equal to zero")
	}
}
