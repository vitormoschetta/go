package interfaces

import (
	"github.com/vitormoschetta/go/models"
)

type ProductRepository interface {
	Save(p models.Product) error
	Update(p models.Product) error
	Delete(id int64) error
	FindAll() []models.Product
	FindByID(id int64) *models.Product
}
