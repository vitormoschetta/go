package mock

import (
	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type ProductRepositoryFake struct {
	Db []models.Product
}

func NewProductRepositoryFake() interfaces.IProductRepository {
	return &ProductRepositoryFake{
		Db: []models.Product{},
	}
}

func (r *ProductRepositoryFake) FindAll() (products []models.Product, err error) {
	return r.Db, nil
}

func (r *ProductRepositoryFake) FindByID(id string) (product models.Product, err error) {
	for _, product := range r.Db {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, nil
}

func (r *ProductRepositoryFake) Save(p models.Product) error {
	r.Db = append(r.Db, p)
	return nil
}

func (r *ProductRepositoryFake) Update(p models.Product) error {
	for i, product := range r.Db {
		if product.ID == p.ID {
			r.Db[i] = p
		}
	}
	return nil
}

func (r *ProductRepositoryFake) Delete(id string) error {
	for i, product := range r.Db {
		if product.ID == id {
			r.Db = append(r.Db[:i], r.Db[i+1:]...)
		}
	}
	return nil
}

func (r *ProductRepositoryFake) FindByCategory(categoryID string) (products []models.Product, err error) {
	for _, product := range r.Db {
		if product.Category.ID == categoryID {
			products = append(products, product)
		}
	}
	return products, nil
}

func (r *ProductRepositoryFake) ApplyPromotionOnProductsByCategory(categoryId string, percentage float64) error {
	for i, product := range r.Db {
		if product.Category.ID == categoryId {
			r.Db[i].Price = product.Price - (product.Price * percentage)
		}
	}
	return nil
}
