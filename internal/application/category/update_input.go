package category

type UpdateCategoryInput struct {
	ID   string
	Name string
}

func (c *UpdateCategoryInput) Validate() (errs []string) {
	if c.ID == "" {
		errs = append(errs, "ID is required")
	}
	if c.Name == "" {
		errs = append(errs, "Name is required")
	}
	return errs
}
