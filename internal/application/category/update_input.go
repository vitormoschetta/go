package category

import "github.com/vitormoschetta/go/internal/domain/models"

type UpdateCategoryInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *UpdateCategoryInput) Validate() (response models.Response) {
	if c.ID == "" {
		response.Errors = append(response.Errors, "ID is required")
	}
	if c.Name == "" {
		response.Errors = append(response.Errors, "Name is required")
	}
	return
}
