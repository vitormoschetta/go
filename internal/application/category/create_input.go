package category

import (
	"github.com/vitormoschetta/go/internal/domain/category"
)

type CreateCategoryInput struct {
	Name string `json:"name"`
}

func (c *CreateCategoryInput) Validate() (errors []string) {
	if c.Name == "" {
		errors = append(errors, "Name is required")
	}
	return
}

func (c *CreateCategoryInput) ToCategoryEntity() category.Category {
	return category.NewCategory(c.Name)
}
