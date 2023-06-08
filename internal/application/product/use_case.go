package product

import (
	"context"
	"log"

	applicationCommon "github.com/vitormoschetta/go/internal/application/common"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type ProductUseCase struct {
	ProductRepository  product.IProductRepository
	CategoryRepository common.IRepository[category.Category]
}

func NewProductUseCase(pR product.IProductRepository, cR common.IRepository[category.Category]) *ProductUseCase {
	return &ProductUseCase{ProductRepository: pR, CategoryRepository: cR}
}

func (u *ProductUseCase) Create(ctx context.Context, input CreateProductInput) (output applicationCommon.Output, statusCode int) {
	output = input.Validate()
	if len(output.Errors) > 0 {
		return output, 400
	}
	category, err := u.CategoryRepository.FindByID(ctx, input.CategoryId)
	if err != nil {
		log.Println("Error on find category: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	if category.ID == "" {
		output.Errors = append(output.Errors, "Category not found")
		return output, 404
	}
	product := input.ToProductModel(category)
	output.Data = product
	err = u.ProductRepository.Save(ctx, product)
	if err != nil {
		log.Println("Error on save product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	return output, 201
}

func (u *ProductUseCase) Update(ctx context.Context, input UpdateProductInput) (output applicationCommon.Output, statusCode int) {
	output = input.Validate()
	if len(output.Errors) > 0 {
		return output, 400
	}
	product, err := u.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		log.Println("Error on find product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	if product.ID == "" {
		output.Errors = append(output.Errors, "Product not found")
		return output, 404
	}
	category, err := u.CategoryRepository.FindByID(ctx, product.Category.ID)
	if err != nil {
		log.Println("Error on find category: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	if category.ID == "" {
		output.Errors = append(output.Errors, "Category not found")
		return output, 404
	}
	product.Update(input.Name, input.Price, category)
	output.Data = product
	err = u.ProductRepository.Update(ctx, product)
	if err != nil {
		log.Println("Error on update product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	return output, 200
}

func (u *ProductUseCase) Delete(ctx context.Context, id string) (output applicationCommon.Output, statusCode int) {
	product, err := u.ProductRepository.FindByID(ctx, id)
	if err != nil {
		log.Println("Error on find product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	if product.ID == "" {
		output.Errors = append(output.Errors, "Product not found")
		return output, 404
	}
	output.Data = product
	err = u.ProductRepository.Delete(ctx, product.ID)
	if err != nil {
		log.Println("Error on delete product: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	return output, 200
}

func (u *ProductUseCase) ApplyPromotion(ctx context.Context, input ApplyPromotionProductInput) (outpu applicationCommon.Output, statusCode int) {
	outpu = input.Validate()
	if len(outpu.Errors) > 0 {
		return outpu, 400
	}
	product, err := u.ProductRepository.FindByID(ctx, input.ProductId)
	if err != nil {
		log.Println("Error on find product: ", err)
		outpu.Errors = append(outpu.Errors, err.Error())
		return outpu, 500
	}
	if product.ID == "" {
		outpu.Errors = append(outpu.Errors, "Product not found")
		return outpu, 404
	}
	product.ApplyPromotion(input.Percentage)
	outpu.Data = product
	err = u.ProductRepository.Update(ctx, product)
	if err != nil {
		log.Println("Error on apply promotion on product: ", err)
		outpu.Errors = append(outpu.Errors, err.Error())
		return outpu, 500
	}
	return outpu, 200
}

func (u *ProductUseCase) ApplyPromotionOnProductsByCategory(ctx context.Context, input ApplyPromotionProductByCategoryInput) (output applicationCommon.Output, statusCode int) {
	output = input.Validate()
	if len(output.Errors) > 0 {
		return output, 400
	}
	category, err := u.CategoryRepository.FindByID(ctx, input.CategoryId)
	if err != nil {
		log.Println("Error on find category: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	if category.ID == "" {
		output.Errors = append(output.Errors, "Category not found")
		return output, 404
	}
	err = u.ProductRepository.ApplyPromotionOnProductsByCategory(ctx, input.CategoryId, input.Percentage)
	if err != nil {
		log.Println("Error on apply promotion on products: ", err)
		output.Errors = append(output.Errors, err.Error())
		return output, 500
	}
	return output, 200
}
