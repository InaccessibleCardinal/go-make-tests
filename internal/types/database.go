package types

type DBIface[T any] interface {
	GetAll() ([]T, error)
	GetById(id int) (*T, error)
	Upsert(T)
}