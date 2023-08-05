package common

type EntityReaderFunc[T any] func(entity T, filters map[string]any) (*T, error)

type ReadConfiguration[E any, F Filterable] struct {
	ReadFunc EntityReaderFunc[E]
	Filters  F
}
