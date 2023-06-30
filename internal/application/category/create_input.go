package category

import (
	"github.com/vitormoschetta/go/internal/domain/category"
)

type CreateCategoryInput struct {
	Name   string   `json:"name"`
	Errors []string `json:"-"`
}

func (c *CreateCategoryInput) IsInvalid() bool {
	c.validate()
	return len(c.Errors) > 0
}

func (c *CreateCategoryInput) ToEntity() category.Category {
	return category.NewCategory(c.Name)
}

func (c *CreateCategoryInput) validate() {
	if c.Name == "" {
		c.Errors = append(c.Errors, "Name is required")
	}
}
