package category

import "github.com/vitormoschetta/go/internal/application/common"

type UpdateCategoryInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *UpdateCategoryInput) Validate() (output common.Output) {
	if c.ID == "" {
		output.Errors = append(output.Errors, "ID is required")
	}
	if c.Name == "" {
		output.Errors = append(output.Errors, "Name is required")
	}
	return
}
