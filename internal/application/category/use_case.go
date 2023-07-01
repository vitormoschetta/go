package category

import (
	"context"
	"log"

	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
	"github.com/vitormoschetta/go/pkg/output"
	"github.com/vitormoschetta/go/pkg/utils"
)

type CategoryUseCases struct {
	Repository common.IRepository[category.Category]
}

func NewCategoryUseCase(pR common.IRepository[category.Category]) *CategoryUseCases {
	return &CategoryUseCases{Repository: pR}
}

func (u *CategoryUseCases) Create(ctx context.Context, input CreateCategoryInput) output.Output {
	out := output.NewOutput(ctx)

	errs := input.Validate()
	if len(errs) > 0 {
		out.SetErrors(output.DomainCodeInvalidInput, errs)
		return out
	}

	entity := input.ToEntity()
	errs = entity.Validate()
	if len(errs) > 0 {
		out.SetError(output.DomainCodeInvalidEntity, "Internal error")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}

	err := u.Repository.Save(ctx, entity)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}

	out.SetOk(entity)
	return out
}

func (u *CategoryUseCases) Update(ctx context.Context, input UpdateCategoryInput) output.Output {
	out := output.NewOutput(ctx)

	errs := input.Validate()
	if len(errs) > 0 {
		out.SetErrors(output.DomainCodeInvalidInput, errs)
		return out
	}

	entity, err := u.Repository.FindByID(ctx, input.ID)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}
	if entity.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Category not found")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}

	entity.Update(input.Name)
	err = u.Repository.Update(ctx, entity)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}

	out.SetOk(entity)
	return out
}

func (u *CategoryUseCases) Delete(ctx context.Context, id string) output.Output {
	out := output.NewOutput(ctx)

	entity, err := u.Repository.FindByID(ctx, id)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}
	if entity.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Category not found")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}

	err = u.Repository.Delete(ctx, entity.ID)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(utils.GetCallingPackage()))
		return out
	}

	out.SetOk(nil)
	return out
}
