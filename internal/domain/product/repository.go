package product

import (
	"github.com/vitormoschetta/go/internal/domain/common"
)

type IProductRepository interface {
	common.IRepository[Product]
	FindByCategory(categoryID string) ([]Product, error)
	ApplyPromotionOnProductsByCategory(categoryId string, percentage float64) error
}
