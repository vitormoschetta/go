package useCases

import (
	"log"

	"github.com/vitormoschetta/go/internal/application/requests"
	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type ProductUseCase struct {
	ProductRepository interfaces.IProductRepository
}

func NewProductUseCase(pR interfaces.IProductRepository) *ProductUseCase {
	return &ProductUseCase{ProductRepository: pR}
}

func (u *ProductUseCase) Save(p requests.CreateProductRequest) (response models.Response, statusCode int) {
	response = p.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	product := p.ToProductModel()
	response.Data = product
	err := u.ProductRepository.Save(product)
	if err != nil {
		log.Println("Error on save product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 201
}

func (u *ProductUseCase) Update(p requests.UpdateProductRequest) (response models.Response, statusCode int) {
	response = p.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	product, err := u.ProductRepository.FindByID(p.ID)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if product.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	product.Update(p.Name, p.Price)
	response.Data = product
	err = u.ProductRepository.Update(product)
	if err != nil {
		log.Println("Error on update product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}

func (u *ProductUseCase) Delete(id string) (response models.Response, statusCode int) {
	product, err := u.ProductRepository.FindByID(id)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if product.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	response.Data = product
	err = u.ProductRepository.Delete(product.ID)
	if err != nil {
		log.Println("Error on delete product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}
