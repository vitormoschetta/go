package category

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) Category {
	return Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

func (c *Category) Update(name string) {
	c.Name = name
}

func (c *Category) Validate() (errs []string) {
	if c.ID == "" {
		errs = append(errs, "ID is required")
	}
	if c.Name == "" {
		errs = append(errs, "Name is required")
	}
	return errs
}
