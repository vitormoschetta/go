package useCases

import (
	"log"

	"github.com/vitormoschetta/go/internal/application/requests"
	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CategoryUseCases struct {
	Repository interfaces.IRepository[models.Category]
}

func NewCategoryUseCase(pR interfaces.IRepository[models.Category]) *CategoryUseCases {
	return &CategoryUseCases{Repository: pR}
}

func (u *CategoryUseCases) Save(p requests.CreateCategoryRequest) (response models.Response, statusCode int) {
	response = p.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	category := p.ToCategoryModel()
	response.Data = category
	err := u.Repository.Save(category)
	if err != nil {
		log.Println("Error on save product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 201
}

func (u *CategoryUseCases) Update(p requests.UpdateCategoryRequest) (response models.Response, statusCode int) {
	response = p.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	category, err := u.Repository.FindByID(p.ID)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if category.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	category.Update(p.Name)
	response.Data = category
	err = u.Repository.Update(category)
	if err != nil {
		log.Println("Error on update product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}

func (u *CategoryUseCases) Delete(id string) (response models.Response, statusCode int) {
	category, err := u.Repository.FindByID(id)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if category.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	err = u.Repository.Delete(category.ID)
	if err != nil {
		log.Println("Error on delete product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}
