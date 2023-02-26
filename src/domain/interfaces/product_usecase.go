package interfaces

import (
	"github.com/vitormoschetta/go/src/domain/models"
)

type ProductUseCase interface {
	Save(p models.Product) models.Response
	Update(p models.Product) models.Response
	Delete(id string) models.Response
}
