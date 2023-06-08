package mock

import (
	"context"
	"errors"

	"github.com/vitormoschetta/go/internal/domain/category"
)

type CategoryRepositoryFake struct {
	Db        []category.Category
	SaveError bool
}

func NewCategoryRepositoryFake() *CategoryRepositoryFake {
	return &CategoryRepositoryFake{
		Db: []category.Category{},
	}
}

func (r *CategoryRepositoryFake) FindAll(ctx context.Context) (categories []category.Category, err error) {
	return r.Db, nil
}

func (r *CategoryRepositoryFake) FindByID(ctx context.Context, id string) (category category.Category, err error) {
	for _, category := range r.Db {
		if category.ID == id {
			return category, nil
		}
	}
	return category, nil
}

func (r *CategoryRepositoryFake) Save(ctx context.Context, p category.Category) error {
	if r.SaveError {
		return errors.New("Error on save category")
	}
	r.Db = append(r.Db, p)
	return nil
}

func (r *CategoryRepositoryFake) Update(ctx context.Context, p category.Category) error {
	for i, category := range r.Db {
		if category.ID == p.ID {
			r.Db[i] = p
		}
	}
	return nil
}

func (r *CategoryRepositoryFake) Delete(ctx context.Context, id string) error {
	for i, category := range r.Db {
		if category.ID == id {
			r.Db = append(r.Db[:i], r.Db[i+1:]...)
		}
	}
	return nil
}
