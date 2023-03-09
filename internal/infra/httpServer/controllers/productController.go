package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitormoschetta/go/internal/application/requests"
	"github.com/vitormoschetta/go/internal/application/useCases"
	"github.com/vitormoschetta/go/internal/domain/interfaces"
)

type ProductController struct {
	UseCase    *useCases.ProductUseCase
	Repository interfaces.IProductRepository
}

func NewProductController(repository interfaces.IProductRepository, useCase *useCases.ProductUseCase) *ProductController {
	return &ProductController{
		UseCase:    useCase,
		Repository: repository,
	}
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	products, err := c.Repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := c.Repository.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) PostProduct(ctx *gin.Context) {
	var request requests.CreateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Save(request)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) PutProduct(ctx *gin.Context) {
	var request requests.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Update(request)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	response, statusCode := c.UseCase.Delete(id)
	ctx.JSON(statusCode, response)
}
