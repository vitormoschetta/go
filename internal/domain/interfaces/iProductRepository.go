package interfaces

import "github.com/vitormoschetta/go/internal/domain/models"

type IProductRepository interface {
	IRepository[models.Product]
	FindByCategory(categoryID string) ([]models.Product, error)
	ApplyPromotionOnProductsByCategory(categoryId string, percentage float64) error
}
