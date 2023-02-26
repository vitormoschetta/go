package models

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) Update(name string, price float64) {
	p.Name = name
	p.Price = price
}
