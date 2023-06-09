package product

import (
	"context"
	"log"

	applicationCommon "github.com/vitormoschetta/go/internal/application/common"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
	"github.com/vitormoschetta/go/internal/domain/product"
	"github.com/vitormoschetta/go/internal/share/utils"
)

type ProductUseCase struct {
	ProductRepository  product.IProductRepository
	CategoryRepository common.IRepository[category.Category]
}

func NewProductUseCase(pR product.IProductRepository, cR common.IRepository[category.Category]) *ProductUseCase {
	return &ProductUseCase{ProductRepository: pR, CategoryRepository: cR}
}

func (u *ProductUseCase) Create(ctx context.Context, input CreateProductInput) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	output.Errors = input.Validate()
	if len(output.Errors) > 0 {
		output.Code = 400
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	category, err := u.CategoryRepository.FindByID(ctx, input.CategoryId)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if category.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Category not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	product := input.ToProductModel(category)
	err = u.ProductRepository.Save(ctx, product)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	output.Data = product
	return output
}

func (u *ProductUseCase) Update(ctx context.Context, input UpdateProductInput) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	output.Errors = input.Validate()
	if len(output.Errors) > 0 {
		output.Code = 400
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	product, err := u.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if product.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Product not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	category, err := u.CategoryRepository.FindByID(ctx, product.Category.ID)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if category.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Category not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	product.Update(input.Name, input.Price, category)
	err = u.ProductRepository.Update(ctx, product)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	output.Data = product
	return output
}

func (u *ProductUseCase) Delete(ctx context.Context, id string) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	product, err := u.ProductRepository.FindByID(ctx, id)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if product.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Product not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	err = u.ProductRepository.Delete(ctx, product.ID)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	output.Data = product
	return output
}

func (u *ProductUseCase) ApplyPromotion(ctx context.Context, input ApplyPromotionProductInput) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	output = input.Validate()
	if len(output.Errors) > 0 {
		output.Code = 400
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	product, err := u.ProductRepository.FindByID(ctx, input.ProductId)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if product.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Product not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	product.ApplyPromotion(input.Percentage)
	err = u.ProductRepository.Update(ctx, product)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	output.Data = product
	return output
}

func (u *ProductUseCase) ApplyPromotionOnProductsByCategory(ctx context.Context, input ApplyPromotionProductByCategoryInput) applicationCommon.Output {
	output := applicationCommon.NewOutput(ctx)
	output = input.Validate()
	if len(output.Errors) > 0 {
		output.Code = 400
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	category, err := u.CategoryRepository.FindByID(ctx, input.CategoryId)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	if category.ID == "" {
		output.Code = 404
		output.Errors = append(output.Errors, "Category not found")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	err = u.ProductRepository.ApplyPromotionOnProductsByCategory(ctx, input.CategoryId, input.Percentage)
	if err != nil {
		output.Code = 500
		output.Errors = append(output.Errors, "Internal error")
		log.Println(output.BuildLogger(), " - ", utils.GetCallingPackage())
		return output
	}
	return output
}
