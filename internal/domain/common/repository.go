package common

import "context"

type IRepository[T any] interface {
	Save(ctx context.Context, item T) error
	Update(ctx context.Context, item T) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]T, error)
	FindByID(ctx context.Context, id string) (T, error)
}
