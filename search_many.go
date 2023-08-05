package commonapi

import (
	"github.com/miguelmartinez624/commonapi/common"
	"github.com/miguelmartinez624/commonapi/errors"
)

type BulEntityQueryDescription[E any, F common.Filterable] struct {
	readFunc EntityBulkReaderFunc[E]
	entity   E
	filters  F

	// Components
	processor common.Processor[E]
}

// TODO Have the input data here
func (descriptor *BulEntityQueryDescription[E, F]) DoAfter(callbacks ...common.PostBulkProcessorFunc[E]) *BulEntityQueryDescription[E, F] {
	descriptor.processor.PostBulkProcessorList = append(descriptor.processor.PostBulkProcessorList, callbacks...)
	return descriptor
}

func (descriptor *BulEntityQueryDescription[E, F]) ReadFrom(fn EntityBulkReaderFunc[E]) bulkEntitiesQuery[E, F] {
	return bulkEntitiesQuery[E, F]{
		readFunc: fn,
		// TODO move to ExecutionCOnfig struct tgo pass all down with one statement
		entity:    descriptor.entity,
		filters:   descriptor.filters,
		Processor: descriptor.processor,
	}
}

func (descriptor *BulEntityQueryDescription[E, F]) Filter(filters F) *BulEntityQueryDescription[E, F] {
	descriptor.filters = filters
	return descriptor
}

// queryEntityCommand Command execution
type bulkEntitiesQuery[E any, F common.Filterable] struct {
	readFunc EntityBulkReaderFunc[E]
	entity   E
	filters  F

	// Components
	common.Processor[E]
}

func (cmd bulkEntitiesQuery[E, F]) Execute() ([]E, *errors.CommandError) {
	entities, _ := cmd.readFunc(cmd.entity, cmd.filters.GetMap(cmd.filters))

	postProcessorErr := cmd.Processor.RunOnMany(entities)
	if postProcessorErr != nil {
		return nil, postProcessorErr
	}
	return entities, nil
}
