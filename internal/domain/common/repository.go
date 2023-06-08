package common

type IRepository[T any] interface {
	Save(T) error
	Update(T) error
	Delete(id string) error
	FindAll() ([]T, error)
	FindByID(id string) (T, error)
}
