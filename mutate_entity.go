package commonapi

import (
	"github.com/miguelmartinez624/commonapi/commands"
	"github.com/miguelmartinez624/commonapi/common"
)

type MutationDescription[E any, F common.Filterable, I any] struct {
	readConfiguration  common.ReadConfiguration[E, F]
	writeConfiguration commands.WriteConfigurationConfiguration[E]
}

func (descriptor *MutationDescription[E, F, I]) FilterBy(filter F) *MutationDescription[E, F, I] {
	descriptor.readConfiguration.Filters = filter
	return descriptor
}

func (descriptor *MutationDescription[E, F, I]) ReadFrom(fn common.EntityReaderFunc[E]) commands.MutateEntityCommand[E, F, I] {
	descriptor.readConfiguration.ReadFunc = fn

	return commands.MutateEntityCommand[E, F, I]{
		// Create Sub Command
		QueryEntityCommand: commands.QueryEntityCommand[E, F]{
			ReadConfiguration: descriptor.readConfiguration,
		},
		// writing configuration
		WriteConfiguration: descriptor.writeConfiguration,
	}
}

func (descriptor *MutationDescription[E, F, I]) WriteTo(fn commands.EntityWriterFunc[E]) *MutationDescription[E, F, I] {
	descriptor.writeConfiguration.WriteFunc = fn
	return descriptor
}
