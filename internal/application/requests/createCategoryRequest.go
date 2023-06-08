package requests

import (
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (c *CreateCategoryRequest) Validate() (response models.Response) {
	if c.Name == "" {
		response.Errors = append(response.Errors, "Name is required")
	}
	return
}

func (c *CreateCategoryRequest) ToCategoryEntity() category.Category {
	return category.NewCategory(c.Name)
}
