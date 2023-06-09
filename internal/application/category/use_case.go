package category

import (
	"context"
	"log"

	applicationCommon "github.com/vitormoschetta/go/internal/application/common"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
	"github.com/vitormoschetta/go/internal/share/utils"
)

type CategoryUseCases struct {
	Repository common.IRepository[category.Category]
}

func NewCategoryUseCase(pR common.IRepository[category.Category]) *CategoryUseCases {
	return &CategoryUseCases{Repository: pR}
}

func (u *CategoryUseCases) Create(ctx context.Context, input CreateCategoryInput) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	output.Errors = input.Validate()
	if len(output.Errors) > 0 {
		output.Code = 400
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	entity := input.ToCategoryEntity()
	err := u.Repository.Save(ctx, entity)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	output.Code = 201
	output.Data = entity
	return output
}

func (u *CategoryUseCases) Update(ctx context.Context, input UpdateCategoryInput) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	output.Errors = input.Validate()
	if len(output.Errors) > 0 {
		output.Code = 400
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	entity, err := u.Repository.FindByID(ctx, input.ID)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if entity.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Category not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	entity.Update(input.Name)
	err = u.Repository.Update(ctx, entity)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	output.Data = entity
	return output
}

func (u *CategoryUseCases) Delete(ctx context.Context, id string) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	entity, err := u.Repository.FindByID(ctx, id)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if entity.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Category not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	err = u.Repository.Delete(ctx, entity.ID)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	output.Data = entity
	return output
}
