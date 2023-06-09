package category

type UpdateCategoryInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c *UpdateCategoryInput) Validate() (errors []string) {
	if c.ID == "" {
		errors = append(errors, "ID is required")
	}
	if c.Name == "" {
		errors = append(errors, "Name is required")
	}
	return
}
