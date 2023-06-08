package category

import (
	"log"

	applicationCommon "github.com/vitormoschetta/go/internal/application/common"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
)

type CategoryUseCases struct {
	Repository common.IRepository[category.Category]
}

func NewCategoryUseCase(pR common.IRepository[category.Category]) *CategoryUseCases {
	return &CategoryUseCases{Repository: pR}
}

func (u *CategoryUseCases) Create(input CreateCategoryInput) (output applicationCommon.Output, statusCode int) {
	output = input.Validate()
	if len(output.Errors) > 0 {
		return output, 400
	}
	entity := input.ToCategoryEntity()
	output.Data = entity
	err := u.Repository.Save(entity)
	if err != nil {
		log.Println("Error on save product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	return output, 201
}

func (u *CategoryUseCases) Update(input UpdateCategoryInput) (output applicationCommon.Output, statusCode int) {
	output = input.Validate()
	if len(output.Errors) > 0 {
		return output, 400
	}
	entity, err := u.Repository.FindByID(input.ID)
	if err != nil {
		log.Println("Error on find product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	if entity.ID == "" {
		output.Errors = append(output.Errors, "Product not found")
		return output, 404
	}
	entity.Update(input.Name)
	output.Data = entity
	err = u.Repository.Update(entity)
	if err != nil {
		log.Println("Error on update product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	return output, 200
}

func (u *CategoryUseCases) Delete(id string) (output applicationCommon.Output, statusCode int) {
	entity, err := u.Repository.FindByID(id)
	if err != nil {
		log.Println("Error on find product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	if entity.ID == "" {
		output.Errors = append(output.Errors, "Product not found")
		return output, 404
	}
	err = u.Repository.Delete(entity.ID)
	if err != nil {
		log.Println("Error on delete product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	return output, 200
}
