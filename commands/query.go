package commands

import (
	"github.com/miguelmartinez624/commonapi/common"
	"github.com/miguelmartinez624/commonapi/errors"
)

// QueryEntityCommand Command execution
type QueryEntityCommand[E any, F common.Filterable] struct {
	common.ReadConfiguration[E, F]
	common.Processor[E]
}

func (cmd QueryEntityCommand[E, F]) Execute() (*E, *errors.CommandError) {
	entityModel := new(E)
	filterMap := cmd.Filters.GetMap(cmd.Filters)
	result, errQuery := cmd.ReadFunc(*entityModel, filterMap)
	if errQuery != nil {
		// TODO add error from function
		cmdErr := errors.NewQueryError(
			errors.QueryError("QE-000"),
			errors.CommandInput{"filters": cmd.Filters},
		)
		return nil, cmdErr
	}

	// Should do on error?
	postProcessorErr := cmd.Processor.Run(result)
	if postProcessorErr != nil {
		return nil, postProcessorErr
	}
	return result, nil
}
