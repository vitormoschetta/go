package category

import (
	"github.com/vitormoschetta/go/internal/domain/category"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type CreateCategoryInput struct {
	Name string `json:"name"`
}

func (c *CreateCategoryInput) Validate() (response models.Response) {
	if c.Name == "" {
		response.Errors = append(response.Errors, "Name is required")
	}
	return
}

func (c *CreateCategoryInput) ToCategoryEntity() category.Category {
	return category.NewCategory(c.Name)
}
