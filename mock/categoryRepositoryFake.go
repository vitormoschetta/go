package mock

import (
	"errors"

	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
)

var SaveError bool

type CategoryRepositoryFake struct {
	Db []category.Category
}

func NewCategoryRepositoryFake() common.IRepository[category.Category] {
	return &CategoryRepositoryFake{
		Db: []category.Category{},
	}
}

func (r *CategoryRepositoryFake) FindAll() (categories []category.Category, err error) {
	return r.Db, nil
}

func (r *CategoryRepositoryFake) FindByID(id string) (category category.Category, err error) {
	for _, category := range r.Db {
		if category.ID == id {
			return category, nil
		}
	}
	return category, nil
}

func (r *CategoryRepositoryFake) Save(p category.Category) error {
	if SaveError {
		return errors.New("Error on save category")
	}
	r.Db = append(r.Db, p)
	return nil
}

func (r *CategoryRepositoryFake) Update(p category.Category) error {
	for i, category := range r.Db {
		if category.ID == p.ID {
			r.Db[i] = p
		}
	}
	return nil
}

func (r *CategoryRepositoryFake) Delete(id string) error {
	for i, category := range r.Db {
		if category.ID == id {
			r.Db = append(r.Db[:i], r.Db[i+1:]...)
		}
	}
	return nil
}
