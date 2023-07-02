package mock

import (
	"context"
	"errors"
	"log"

	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/pkg/pagination"
	"github.com/vitormoschetta/go/pkg/utils"
)

type CategoryRepositoryFake struct {
	storage   []category.Category
	SaveError bool
}

func NewCategoryRepositoryFake() *CategoryRepositoryFake {
	return &CategoryRepositoryFake{
		storage: []category.Category{},
	}
}

func (r *CategoryRepositoryFake) Seed(items []category.Category) {
	r.storage = items
}

func (r *CategoryRepositoryFake) FindAll(ctx context.Context, pagination *pagination.Pagination) (categories []category.Category, err error) {
	pagination.Total = len(r.storage)
	pagination.BuildLastPage()
	for _, category := range r.storage {
		categories = append(categories, category)
	}
	return
}

func (r *CategoryRepositoryFake) FindByID(ctx context.Context, id string) (category category.Category, err error) {
	for _, category := range r.storage {
		if category.ID == id {
			return category, nil
		}
	}
	return category, nil
}

func (r *CategoryRepositoryFake) Save(ctx context.Context, p category.Category) error {
	if r.SaveError {
		err := errors.New("Error on save category")
		log.Print(utils.BuildLoggerWithErr(ctx, err) + " - " + utils.GetCallingPackage())
		return err
	}
	r.storage = append(r.storage, p)
	return nil
}

func (r *CategoryRepositoryFake) Update(ctx context.Context, p category.Category) error {
	for i, category := range r.storage {
		if category.ID == p.ID {
			r.storage[i] = p
		}
	}
	return nil
}

func (r *CategoryRepositoryFake) Delete(ctx context.Context, id string) error {
	for i, category := range r.storage {
		if category.ID == id {
			r.storage = append(r.storage[:i], r.storage[i+1:]...)
		}
	}
	return nil
}
