package interfaces

import (
	"github.com/vitormoschetta/go/internal/domain/models"
)

type IProductUseCase interface {
	Save(p models.Product) (models.Response, int)
	Update(p models.Product) (models.Response, int)
	Delete(id string) (models.Response, int)
}
