package commands

import (
	"github.com/miguelmartinez624/commonapi/common"
	"github.com/miguelmartinez624/commonapi/errors"
	"github.com/miguelmartinez624/commonapi/reflectutils"
)

type EntityWriterFunc[T any] func(entity T) (*T, error)

type WriteConfigurationConfiguration[E any] struct {

	// function to actually save/persist the data
	WriteFunc EntityWriterFunc[E]
	// Pipeline Points
	// What to do before the write
	BeforeWrite *common.Processor[E]
	//to run after the write
	AfterWrite *common.Processor[E]
}

func (wc *WriteConfigurationConfiguration[E]) AddBeforeWrite(process common.PostProcessorFunc[E]) {
	if wc.BeforeWrite == nil {
		wc.BeforeWrite = &common.Processor[E]{}
	}
	wc.BeforeWrite.PostProcessorList = append(wc.BeforeWrite.PostProcessorList, process)

}

func (wc *WriteConfigurationConfiguration[E]) AddAfterWrite(process ...common.PostProcessorFunc[E]) {
	if wc.AfterWrite == nil {
		wc.AfterWrite = &common.Processor[E]{}
	}
	wc.AfterWrite.PostProcessorList = append(wc.AfterWrite.PostProcessorList, process...)
}

// WriteEntityCommand
// set a writer before been able to execute the action
type WriteEntityCommand[I any, E any] struct {
	WriteConfigurationConfiguration[E]
}

// Execute TODO document flow
func (cmd WriteEntityCommand[I, E]) Execute(input I) (*E, *errors.CommandError) {
	// TODO add error handler for create entity
	entity := new(E)
	reflectutils.ReflectValues[E, I](entity, &input)
	if cmd.BeforeWrite != nil {
		if processErr := cmd.BeforeWrite.Run(entity); processErr != nil {
			return nil, processErr
		}
	}

	// TODO handle error
	enti, _ := cmd.WriteFunc(*entity)
	if cmd.AfterWrite != nil {
		if afterProcessErr := cmd.AfterWrite.Run(enti); afterProcessErr != nil {
			// TODO wrap with a after process code so we can ignore this errors
			return enti, afterProcessErr
		}
	}
	return enti, nil
}
