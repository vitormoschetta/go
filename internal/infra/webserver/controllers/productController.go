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
	var request requests.CreateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Save(request)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) Put(ctx *gin.Context) {
	var request requests.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Update(request)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	response, statusCode := c.UseCase.Delete(id)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) PutPromotion(ctx *gin.Context) {
	var request requests.ApplyPromotionProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.ApplyPromotion(request)
	ctx.JSON(statusCode, response)
}

func (c *ProductController) PutPromotionbyCategory(ctx *gin.Context) {
	var request requests.ApplyPromotionProductByCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.ApplyPromotionOnProductsByCategory(request)
	ctx.JSON(statusCode, response)
}
