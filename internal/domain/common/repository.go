package common

import (
	"context"

	"github.com/vitormoschetta/go/pkg/pagination"
)

type IRepository[T any] interface {
	Save(ctx context.Context, item T) error
	Update(ctx context.Context, item T) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, pagination *pagination.Pagination) ([]T, error)
	FindByID(ctx context.Context, id string) (T, error)
}
