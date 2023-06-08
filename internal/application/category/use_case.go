package category

import (
	"log"

	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/general"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CategoryUseCases struct {
	Repository general.IRepository[category.Category]
}

func NewCategoryUseCase(pR general.IRepository[category.Category]) *CategoryUseCases {
	return &CategoryUseCases{Repository: pR}
}

func (u *CategoryUseCases) Create(input CreateCategoryInput) (response models.Response, statusCode int) {
	response = input.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	entity := input.ToCategoryEntity()
	response.Data = entity
	err := u.Repository.Save(entity)
	if err != nil {
		log.Println("Error on save product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 201
}

func (u *CategoryUseCases) Update(input UpdateCategoryInput) (response models.Response, statusCode int) {
	response = input.Validate()
	if len(response.Errors) > 0 {
		return response, 400
	}
	entity, err := u.Repository.FindByID(input.ID)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if entity.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	entity.Update(input.Name)
	response.Data = entity
	err = u.Repository.Update(entity)
	if err != nil {
		log.Println("Error on update product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}

func (u *CategoryUseCases) Delete(id string) (response models.Response, statusCode int) {
	entity, err := u.Repository.FindByID(id)
	if err != nil {
		log.Println("Error on find product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	if entity.ID == "" {
		response.Errors = append(response.Errors, "Product not found")
		return response, 404
	}
	err = u.Repository.Delete(entity.ID)
	if err != nil {
		log.Println("Error on delete product: ", err)
		response.Errors = append(response.Errors, err.Error())
		return response, 500
	}
	return response, 200
}
