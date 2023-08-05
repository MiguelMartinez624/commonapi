package commands

import (
	"github.com/miguelmartinez624/commonapi/common"
	"github.com/miguelmartinez624/commonapi/errors"
	"github.com/miguelmartinez624/commonapi/reflectutils"
)

// MutateEntityCommand its composed by a search and a write function the input for should be fields
// that exist on the target entity the pipeline itself its a process like
//
// [Search] -> [Mutate] -> [PostProcessors] -> [Persist]
//
// This can be compose of some of the core commands like the "searchCommand" and the "writeCommand"
type MutateEntityCommand[E any, F common.Filterable, I any] struct {
	// Add config to this commands and reuse configurations for Read and Write so cam be pass down
	//to the downstream structs
	QueryEntityCommand QueryEntityCommand[E, F]
	// This will receibe the same entity and write the same entity as input
	//this will not make the mutation
	WriteConfiguration WriteConfigurationConfiguration[E]
}

func (m MutateEntityCommand[E, F, I]) Execute(changes I) (*E, *errors.CommandError) {
	// Configure the query
	entity, errQuery := m.QueryEntityCommand.Execute()
	if errQuery != nil {
		return nil, errQuery
	}

	reflectutils.ReflectValues[E, I](entity, &changes)
	persisted, errWrite := m.WriteConfiguration.WriteFunc(*entity)
	if errWrite != nil {
		return nil, nil
	}

	return persisted, nil
}
