package commonapi

import (
	"github.com/miguelmartinez624/commonapi/commands"
)

type CreateDescription[I any, E any] struct {
	writeFunc commands.EntityWriterFunc[E]
}

func (descriptor *CreateDescription[I, E]) WriteTo(fn commands.EntityWriterFunc[E]) commands.WriteEntityCommand[I, E] {
	return commands.WriteEntityCommand[I, E]{
		WriteConfigurationConfiguration: commands.WriteConfigurationConfiguration[E]{
			WriteFunc: fn,
		},
	}
}
