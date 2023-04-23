package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitormoschetta/go/internal/application/requests"
	"github.com/vitormoschetta/go/internal/application/useCases"
	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CategoryController struct {
	UseCase    *useCases.CategoryUseCases
	Repository interfaces.IRepository[models.Category]
}

func NewCategoryController(repository interfaces.IRepository[models.Category], useCase *useCases.CategoryUseCases) *CategoryController {
	return &CategoryController{
		UseCase:    useCase,
		Repository: repository,
	}
}

// Show Categories godoc
//
// @Summary      Show all categories
// @Description  Get all categories
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Category
// @Router       /categories [get]
func (c *CategoryController) GetAll(ctx *gin.Context) {
	items, err := c.Repository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (c *CategoryController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	item, err := c.Repository.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (c *CategoryController) Post(ctx *gin.Context) {
	var request requests.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Save(request)
	ctx.JSON(statusCode, response)
}

func (c *CategoryController) Put(ctx *gin.Context) {
	var request requests.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Update(request)
	ctx.JSON(statusCode, response)
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	response, statusCode := c.UseCase.Delete(id)
	ctx.JSON(statusCode, response)
}
