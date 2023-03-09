package mock

import (
	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type ProductRepositoryFake struct {
	Products []models.Product
}

func NewProductRepositoryFake() interfaces.IProductRepository {
	return &ProductRepositoryFake{}
}

func (r *ProductRepositoryFake) Save(p models.Product) error {
	r.Products = append(r.Products, p)
	return nil
}

func (r *ProductRepositoryFake) Update(p models.Product) error {
	for i, product := range r.Products {
		if product.ID == p.ID {
			r.Products[i] = p
		}
	}
	return nil
}

func (r *ProductRepositoryFake) Delete(id string) error {
	for i, product := range r.Products {
		if product.ID == id {
			r.Products = append(r.Products[:i], r.Products[i+1:]...)
		}
	}
	return nil
}

func (r *ProductRepositoryFake) FindAll() ([]models.Product, error) {
	return r.Products, nil
}

func (r *ProductRepositoryFake) FindByID(id string) (models.Product, error) {
	for _, product := range r.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, nil
}
