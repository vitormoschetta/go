package mock

import (
	"context"

	"github.com/vitormoschetta/go/internal/domain/product"
)

type ProductRepositoryFake struct {
	Db []product.Product
}

func NewProductRepositoryFake() product.IProductRepository {
	return &ProductRepositoryFake{
		Db: []product.Product{},
	}
}

func (r *ProductRepositoryFake) FindAll(ctx context.Context) (products []product.Product, err error) {
	return r.Db, nil
}

func (r *ProductRepositoryFake) FindByID(ctx context.Context, id string) (product product.Product, err error) {
	for _, product := range r.Db {
		if product.ID == id {
			return product, nil
		}
	}
	return product, nil
}

func (r *ProductRepositoryFake) Save(ctx context.Context, p product.Product) error {
	r.Db = append(r.Db, p)
	return nil
}

func (r *ProductRepositoryFake) Update(ctx context.Context, p product.Product) error {
	for i, product := range r.Db {
		if product.ID == p.ID {
			r.Db[i] = p
		}
	}
	return nil
}

func (r *ProductRepositoryFake) Delete(ctx context.Context, id string) error {
	for i, product := range r.Db {
		if product.ID == id {
			r.Db = append(r.Db[:i], r.Db[i+1:]...)
		}
	}
	return nil
}

func (r *ProductRepositoryFake) FindByCategory(ctx context.Context, categoryID string) (products []product.Product, err error) {
	for _, product := range r.Db {
		if product.Category.ID == categoryID {
			products = append(products, product)
		}
	}
	return products, nil
}

func (r *ProductRepositoryFake) ApplyPromotionOnProductsByCategory(ctx context.Context, categoryId string, percentage float64) error {
	for i, product := range r.Db {
		if product.Category.ID == categoryId {
			r.Db[i].Price = product.Price - (product.Price * percentage)
		}
	}
	return nil
}
