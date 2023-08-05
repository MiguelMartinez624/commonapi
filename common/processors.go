package common

import (
	"github.com/miguelmartinez624/commonapi/errors"
)

// Processor work with collections payload of data
type Processor[E any] struct {
	PostBulkProcessorList []PostBulkProcessorFunc[E]
	PostProcessorList     []PostProcessorFunc[E]
}

func (p Processor[E]) New() Processor[E] {
	return Processor[E]{
		PostBulkProcessorList: []PostBulkProcessorFunc[E]{},
	}
}

func NewPostProcessor[E any]() Processor[E] {
	return Processor[E]{
		PostBulkProcessorList: []PostBulkProcessorFunc[E]{},
	}
}

func (p *Processor[E]) RunOnMany(data []E) *errors.CommandError {
	for _, p := range p.PostBulkProcessorList {
		errProcess := p(data)
		if errProcess != nil {
			return errProcess
		}
	}

	return nil
}

func (p *Processor[E]) Run(data *E) *errors.CommandError {
	for _, p := range p.PostProcessorList {
		errProcess := p(data)
		if errProcess != nil {
			return errProcess
		}
	}

	return nil
}

// PostBulkProcessorFunc run after the query or the create are made they can modify the object
// and return a different entity, this are executed linearly to avoid race condition
// implement a DoAfterOnEach too to have more granularity, for mapping we have json tags.
type PostBulkProcessorFunc[E any] func(entity []E) *errors.CommandError

// PostProcessorFunc run after the query or the create are made they can modify the object
// and return a different entity, this are executed linearly to avoid race condition
// implement a DoAfterOnEach too to have more granularity, for mapping we have json tags.
type PostProcessorFunc[E any] func(entity *E) *errors.CommandError
