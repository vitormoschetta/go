package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	categoryApplication "github.com/vitormoschetta/go/internal/application/category"
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/common"
)

type CategoryController struct {
	UseCase    *categoryApplication.CategoryUseCases
	Repository common.IRepository[category.Category]
}

func NewCategoryController(repository common.IRepository[category.Category], useCase *categoryApplication.CategoryUseCases) *CategoryController {
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
	var input categoryApplication.CreateCategoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Create(input)
	ctx.JSON(statusCode, response)
}

func (c *CategoryController) Put(ctx *gin.Context) {
	var input categoryApplication.UpdateCategoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, statusCode := c.UseCase.Update(input)
	ctx.JSON(statusCode, response)
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	response, statusCode := c.UseCase.Delete(id)
	ctx.JSON(statusCode, response)
}
