package category

type UpdateCategoryInput struct {
	ID   string
	Name string
}

func NewUpdateCategoryInput(id, name string) *UpdateCategoryInput {
	return &UpdateCategoryInput{
		ID:   id,
		Name: name,
	}
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
