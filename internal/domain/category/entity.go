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
