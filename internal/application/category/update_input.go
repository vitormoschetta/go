package category

type UpdateCategoryInput struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Errors []string `json:"-"`
}

func (c *UpdateCategoryInput) IsInvalid() bool {
	c.validate()
	return len(c.Errors) > 0
}

func (c *UpdateCategoryInput) validate() {
	if c.ID == "" {
		c.Errors = append(c.Errors, "ID is required")
	}
	if c.Name == "" {
		c.Errors = append(c.Errors, "Name is required")
	}
}
