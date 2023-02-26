package useCases

import (
	"github.com/vitormoschetta/go/src/interfaces"
	"github.com/vitormoschetta/go/src/models"
)

type ProductUseCase struct {
	ProductRepository interfaces.ProductRepository
}

// Função que retorna uma instância de ProductUseCase
// Injetamos a dependência de ProductRepository (interface)
func NewProductUseCase(pR interfaces.ProductRepository) interfaces.ProductUseCase {
	return &ProductUseCase{ProductRepository: pR}
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
	product := u.ProductRepository.FindByID(p.ID)
	if product == nil {
		return models.Response{Errors: []string{"Product not found"}}
	}
	product.Update(p.Name, p.Price)
	err := u.ProductRepository.Update(p)
	if err != nil {
		return models.Response{Errors: []string{err.Error()}}
	}
	return models.Response{}
}

func (u *ProductUseCase) Delete(id string) models.Response {
	err := u.ProductRepository.Delete(id)
	if err != nil {
		return models.Response{Errors: []string{err.Error()}}
	}
	return models.Response{}
}
