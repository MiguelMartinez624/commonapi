package commonapi

import (
	"github.com/miguelmartinez624/commonapi/commands"
	"github.com/miguelmartinez624/commonapi/common"
	"github.com/miguelmartinez624/commonapi/errors"
)

type DeleteDescription[E any, F common.Filterable] struct {
	common.ReadConfiguration[E, F]
	writeFunc EntityRemoveFunc[E]
}

func (descriptor *DeleteDescription[E, F]) FilterBy(filter F) *DeleteDescription[E, F] {
	descriptor.ReadConfiguration.Filters = filter
	return descriptor
}

func (descriptor *DeleteDescription[E, F]) ReadFrom(fn common.EntityReaderFunc[E]) deleteCommand[E, F] {
	descriptor.ReadConfiguration.ReadFunc = fn

	return deleteCommand[E, F]{
		queryEntityCommand: commands.QueryEntityCommand[E, F]{
			ReadConfiguration: descriptor.ReadConfiguration,
		},
		// Need to set the named field to initialize the other fields as zero value.
		writeFunc: descriptor.writeFunc,
	}
}

func (descriptor *DeleteDescription[E, F]) RemoveFrom(fn EntityRemoveFunc[E]) *DeleteDescription[E, F] {
	descriptor.writeFunc = fn
	return descriptor
}

// Delete comand inpunts and logic execution

// This can be compose of some of the core commands like the "searchCommand" and the "writeCommand"
type deleteCommand[E any, F common.Filterable] struct {
	writeFunc EntityRemoveFunc[E]
	// Add config to this commands and reuse configurations for Read and Write so cam be pass down
	//to the downstream structs
	queryEntityCommand commands.QueryEntityCommand[E, F]
}

func (m deleteCommand[E, F]) Execute() (*E, *errors.CommandError) {
	// Configure the query
	filters := m.queryEntityCommand.Filters

	entity, errQuery := m.queryEntityCommand.Execute()
	if errQuery != nil {
		return nil, errQuery
	}

	// MEMORY : there is a copy here be careful to benchmark
	//TODO should this filter be here be here?
	deleteEntity, errDelete := m.writeFunc(*entity, filters.GetMap(filters))
	if errDelete != nil {
		return nil, nil
	}

	return deleteEntity, nil
}
