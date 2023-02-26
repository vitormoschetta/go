package mock

import (
	"github.com/vitormoschetta/go/src/domain/interfaces"
	"github.com/vitormoschetta/go/src/domain/models"
)

type FakeProductRepository struct {
	Products []models.Product
}

func NewFakeProductRepository() interfaces.ProductRepository {
	return &FakeProductRepository{}
}

func (r *FakeProductRepository) Save(p models.Product) error {
	r.Products = append(r.Products, p)
	return nil
}

func (r *FakeProductRepository) Update(p models.Product) error {
	for i, product := range r.Products {
		if product.ID == p.ID {
			r.Products[i] = p
		}
	}
	return nil
}

func (r *FakeProductRepository) Delete(id string) error {
	for i, product := range r.Products {
		if product.ID == id {
			r.Products = append(r.Products[:i], r.Products[i+1:]...)
		}
	}
	return nil
}

func (r *FakeProductRepository) FindAll() []models.Product {
	return r.Products
}

func (r *FakeProductRepository) FindByID(id string) *models.Product {
	for _, product := range r.Products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
