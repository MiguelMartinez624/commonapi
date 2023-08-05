package commonapi

import "github.com/miguelmartinez624/commonapi/errors"

type Executable[E any] interface {
	Execute() (*E, *errors.CommandError)
}

type ExecutableSearch[E any] interface {
	Execute() ([]E, *errors.CommandError)
}

type EntityBulkReaderFunc[T any] func(entity T, filters map[string]any) ([]T, error)

// What is the entity for you may ask is for searching the type
type EntityRemoveFunc[T any] func(entity T, filters map[string]any) (*T, error)
