package product

import (
	"context"
	"log"

	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
	"github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/pkg/output"
	"github.com/vitormoschetta/go/pkg/utils"
)

type ProductUseCase struct {
	ProductRepository  product.IProductRepository
	CategoryRepository common.IRepository[category.Category]
}

func NewProductUseCase(pR product.IProductRepository, cR common.IRepository[category.Category]) *ProductUseCase {
	return &ProductUseCase{ProductRepository: pR, CategoryRepository: cR}
}

func (u *ProductUseCase) Create(ctx context.Context, input CreateProductInput) output.Output {
	out := output.NewOutput(ctx)
	if input.IsInvalid() {
		out.SetErrors(output.DomainCodeInvalidInput, input.Errors)
		return out
	}
	category, err := u.CategoryRepository.FindByID(ctx, input.CategoryId)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	if category.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Category not found")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	product := input.ToEntity(category)
	err = u.ProductRepository.Save(ctx, product)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	out.SetOk(product)
	return out
}

func (u *ProductUseCase) Update(ctx context.Context, input UpdateProductInput) output.Output {
	out := output.NewOutput(ctx)
	if input.IsInvalid() {
		out.SetErrors(output.DomainCodeInvalidInput, input.Errors)
		return out
	}
	product, err := u.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	if product.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Product not found")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	category, err := u.CategoryRepository.FindByID(ctx, product.Category.ID)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	if category.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Category not found")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	product.Update(input.Name, input.Price, category)
	err = u.ProductRepository.Update(ctx, product)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	out.SetOk(product)
	return out
}

func (u *ProductUseCase) Delete(ctx context.Context, id string) output.Output {
	out := output.NewOutput(ctx)
	product, err := u.ProductRepository.FindByID(ctx, id)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	if product.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Product not found")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	err = u.ProductRepository.Delete(ctx, product.ID)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	out.SetOk(nil)
	return out
}

func (u *ProductUseCase) ApplyPromotion(ctx context.Context, input ApplyPromotionProductInput) output.Output {
	out := output.NewOutput(ctx)
	if input.IsInvalid() {
		out.SetErrors(output.DomainCodeInvalidInput, input.Errors)
		return out
	}
	product, err := u.ProductRepository.FindByID(ctx, input.ProductId)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	if product.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Product not found")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	product.ApplyPromotion(input.Percentage)
	err = u.ProductRepository.Update(ctx, product)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	out.SetOk(product)
	return out
}

func (u *ProductUseCase) ApplyPromotionOnProductsByCategory(ctx context.Context, input ApplyPromotionProductByCategoryInput) output.Output {
	out := output.NewOutput(ctx)
	if input.IsInvalid() {
		out.SetErrors(output.DomainCodeInvalidInput, input.Errors)
		return out
	}
	category, err := u.CategoryRepository.FindByID(ctx, input.CategoryId)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	if category.ID == "" {
		out.SetError(output.DomainCodeNotFound, "Category not found")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	err = u.ProductRepository.ApplyPromotionOnProductsByCategory(ctx, input.CategoryId, input.Percentage)
	if err != nil {
		out.SetError(output.DomainCodeInternalError, "Internal error")
		log.Println(out.BuildLogger(), " - ", utils.GetCallingPackage())
		return out
	}
	out.SetOk(nil)
	return out
}
