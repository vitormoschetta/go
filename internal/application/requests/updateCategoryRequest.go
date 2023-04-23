package requests

import "github.com/vitormoschetta/go/internal/domain/models"

type UpdateCategoryRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *UpdateCategoryRequest) Validate() (response models.Response) {
	if c.ID == "" {
		response.Errors = append(response.Errors, "ID is required")
	}
	if c.Name == "" {
		response.Errors = append(response.Errors, "Name is required")
	}
	return
}
