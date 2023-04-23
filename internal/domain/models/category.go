package models

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategory(name string) Category {
	return Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

func NewCategoryWithID(id, name string) *Category {
	return &Category{
		ID:   id,
		Name: name,
	}
}

func (c *Category) Update(name string) {
	c.Name = name
}
