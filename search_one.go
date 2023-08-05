package commonapi

import (
	"github.com/miguelmartinez624/commonapi/commands"
	"github.com/miguelmartinez624/commonapi/common"
)

type QueryEntityDescription[E any, F common.Filterable] struct {
	readConfiguration common.ReadConfiguration[E, F]

	// Components
	processor common.Processor[E]
}

// TODO Have the input data to the processor to be able to use the whole context
func (descriptor *QueryEntityDescription[E, F]) DoAfter(callbacks ...common.PostProcessorFunc[E]) *QueryEntityDescription[E, F] {
	descriptor.processor.PostProcessorList = append(descriptor.processor.PostProcessorList, callbacks...)
	return descriptor
}

func (descriptor *QueryEntityDescription[E, F]) ReadFrom(fn common.EntityReaderFunc[E]) commands.QueryEntityCommand[E, F] {
	descriptor.readConfiguration.ReadFunc = fn

	return commands.QueryEntityCommand[E, F]{
		ReadConfiguration: descriptor.readConfiguration,
		Processor:         descriptor.processor,
	}
}

func (descriptor *QueryEntityDescription[E, F]) Filter(filters F) *QueryEntityDescription[E, F] {
	descriptor.readConfiguration.Filters = filters
	return descriptor
}
