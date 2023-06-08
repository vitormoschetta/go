package product

import (
	"github.com/vitormoschetta/go/internal/domain/general"
)

type IProductRepository interface {
	general.IRepository[Product]
	FindByCategory(categoryID string) ([]Product, error)
	ApplyPromotionOnProductsByCategory(categoryId string, percentage float64) error
}
