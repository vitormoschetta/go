package interfaces

import (
	"github.com/vitormoschetta/go/internal/domain/models"
)

type IProductRepository interface {
	Save(p models.Product) error
	Update(p models.Product) error
	Delete(id string) error
	FindAll() ([]models.Product, error)
	FindByID(id string) (models.Product, error)
}
