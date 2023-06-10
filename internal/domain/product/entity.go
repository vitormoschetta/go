package product

import (
	"github.com/google/uuid"
	"github.com/vitormoschetta/go/internal/domain/category"
)

type Product struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Price    float64           `json:"price"`
	Category category.Category `json:"category"`
}

func NewProduct(name string, price float64, category category.Category) Product {
	return Product{
		ID:       uuid.New().String(),
		Name:     name,
		Price:    price,
		Category: category,
	}
}

func (p *Product) Update(name string, price float64, category category.Category) {
	p.Name = name
	p.Price = price
	p.Category = category
}

func (p *Product) ApplyPromotion(discount float64) {
	p.Price = p.Price - (p.Price * discount)
}
