package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productApplication "github.com/vitormoschetta/go/internal/application/product"
	"github.com/vitormoschetta/go/internal/domain/product"
)

type ProductController struct {
	UseCase    *productApplication.ProductUseCase
	Repository product.IProductRepository
}

func NewProductController(repository product.IProductRepository, useCase *productApplication.ProductUseCase) *ProductController {
	return &ProductController{
		UseCase:    useCase,
		Repository: repository,
	}
}

// Show Products godoc
//
// @Summary      Show all products
// @Description  Get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Product
// @Router       /products [get]
func (c *ProductController) GetAll(ctx *gin.Context) {
	items, err := c.Repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *ProductController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	item, err := c.Repository.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *ProductController) GetByCategory(ctx *gin.Context) {
	categoryId := ctx.Param("category_id")
	items, err := c.Repository.FindByCategory(categoryId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *ProductController) Post(ctx *gin.Context) {
	var input productApplication.CreateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Save(input)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) Put(ctx *gin.Context) {
	var input productApplication.UpdateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Update(input)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	response, statusCode := c.UseCase.Delete(id)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) PutPromotion(ctx *gin.Context) {
	var input productApplication.ApplyPromotionProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.ApplyPromotion(input)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) PutPromotionbyCategory(ctx *gin.Context) {
	var input productApplication.ApplyPromotionProductByCategoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.ApplyPromotionOnProductsByCategory(input)
	ctx.JSON(statusCode, response)
}
