package interfaces

import (
	"github.com/vitormoschetta/go/models"
)

type ProductRepository interface {
	Save(p models.Product) error
	Update(p models.Product) error
	Delete(id string) error
	FindAll() []models.Product
	FindByID(id string) *models.Product
}
