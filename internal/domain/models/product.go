package models

import "github.com/google/uuid"

type Product struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

func NewProduct(name string, price float64, category Category) Product {
	return Product{
		ID:       uuid.New().String(),
		Name:     name,
		Price:    price,
		Category: category,
	}
}

func NewProductWithID(id, name string, price float64, category Category) *Product {
	return &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Category: category,
	}
}

func (p *Product) Update(name string, price float64, category Category) {
	p.Name = name
	p.Price = price
	p.Category = category
}

func (p *Product) ApplyPromotion(discount float64) {
	p.Price = p.Price - (p.Price * discount)
}
