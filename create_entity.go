package commonapi

import (
	"github.com/miguelmartinez624/commonapi/commands"
	"github.com/miguelmartinez624/commonapi/common"
)

type CreateDescription[I any, E any] struct {
	config commands.WriteConfigurationConfiguration[E]
}

func (descriptor *CreateDescription[I, E]) DoBeforeWrite(process ...common.PostProcessorFunc[E]) *CreateDescription[I, E] {
	for _, pfunc := range process {
		descriptor.config.AddBeforeWrite(pfunc)

	}
	return descriptor
}

func (descriptor *CreateDescription[I, E]) DoAfterWrite(process ...common.PostProcessorFunc[E]) *CreateDescription[I, E] {
	for _, pfunc := range process {
		descriptor.config.AddAfterWrite(pfunc)

	}
	return descriptor
}

func (descriptor *CreateDescription[I, E]) WriteTo(fn commands.EntityWriterFunc[E]) *CreateDescription[I, E] {
	descriptor.config.WriteFunc = fn
	return descriptor
}

func (descriptor *CreateDescription[I, E]) Done() commands.WriteEntityCommand[I, E] {
	return commands.WriteEntityCommand[I, E]{
		WriteConfigurationConfiguration: descriptor.config,
	}
}
