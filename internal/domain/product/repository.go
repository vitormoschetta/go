package product

import (
	"context"

	"github.com/vitormoschetta/go/internal/domain/common"
)

type IProductRepository interface {
	common.IRepository[Product]
	FindByCategory(ctx context.Context, categoryID string) ([]Product, error)
	ApplyPromotionOnProductsByCategory(ctx context.Context, categoryId string, percentage float64) error
}
