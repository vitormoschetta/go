package category

import (
	"github.com/vitormoschetta/go/internal/domain/category"
)

type CreateCategoryInput struct {
	Name string
}

func (c *CreateCategoryInput) Validate() (errs []string) {
	if c.Name == "" {
		errs = append(errs, "Name is required")
	}
	return errs
}

func (c *CreateCategoryInput) ToEntity() category.Category {
	return category.NewCategory(c.Name)
}
