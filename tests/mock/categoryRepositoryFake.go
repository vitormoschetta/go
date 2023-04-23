package mock

import (
	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CategoryRepositoryFake struct {
	Db []models.Category
}

func NewCategoryRepositoryFake() interfaces.IRepository[models.Category] {
	return &CategoryRepositoryFake{
		Db: []models.Category{},
	}
}

func (r *CategoryRepositoryFake) FindAll() (categories []models.Category, err error) {
	return r.Db, nil
}

func (r *CategoryRepositoryFake) FindByID(id string) (category models.Category, err error) {
	for _, category := range r.Db {
		if category.ID == id {
			return category, nil
		}
	}
	return models.Category{}, nil
}

func (r *CategoryRepositoryFake) Save(p models.Category) error {
	r.Db = append(r.Db, p)
	return nil
}

func (r *CategoryRepositoryFake) Update(p models.Category) error {
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
