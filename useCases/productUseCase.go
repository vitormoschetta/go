package useCases

import (
	"github.com/vitormoschetta/go/interfaces"
	"github.com/vitormoschetta/go/mock"
	"github.com/vitormoschetta/go/models"
)

type ProductUseCase struct {
	ProductRepository mock.FakeProductRepository
}

func NewProductUseCase() interfaces.ProductUseCase {
	return &ProductUseCase{}
}

func (u *ProductUseCase) Save(p models.Product) models.Response {
	errors := []string{}
	if p.Name == "" {
		errors = append(errors, "Name is required")
	}
	if p.Price == 0 {
		errors = append(errors, "Price is required")
	}
	if len(errors) > 0 {
		return models.Response{Errors: errors}
	}
	err := u.ProductRepository.Save(p)
	if err != nil {
		return models.Response{Errors: []string{err.Error()}}
	}
	return models.Response{}
}

func (u *ProductUseCase) Update(p models.Product) models.Response {
	errors := []string{}
	if p.Name == "" {
		errors = append(errors, "Name is required")
	}
	if p.Price == 0 {
		errors = append(errors, "Price is required")
	}
	if len(errors) > 0 {
		return models.Response{Errors: errors}
	}
	err := u.ProductRepository.Update(p)
	if err != nil {
		return models.Response{Errors: []string{err.Error()}}
	}
	return models.Response{}
}

func (u *ProductUseCase) Delete(id int64) models.Response {
	err := u.ProductRepository.Delete(id)
	if err != nil {
		return models.Response{Errors: []string{err.Error()}}
	}
	return models.Response{}
}
