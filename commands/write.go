package commands

import (
	"github.com/miguelmartinez624/commonapi/errors"
	"github.com/miguelmartinez624/commonapi/reflectutils"
)

type EntityWriterFunc[T any] func(entity T) (*T, error)

type WriteConfigurationConfiguration[E any] struct {
	WriteFunc EntityWriterFunc[E]
}

// WriteEntityCommand
// set a writer before been able to execute the action
type WriteEntityCommand[I any, E any] struct {
	WriteConfigurationConfiguration[E]
}

// Execute aca se ejecutan los procesadores antes de escribir,
//
//	 [TODO] regla de unanicidad de algo seria definir en el objeto "unique" y buscar por esos campos si existe uno se notifica
//		que ya existe entitdad con valor X, seria jalas los unique field y ahcer las queries de ese nombre, este podria ser combinado
//		al estilo id:"unique:name,user_id" no quiere decir q ese sea el ID sino q el id mas esos campos lo identifican como una entidad
//		ahi se tomaran esos campos y se realizara una "validacion de unisidad antes de guardar"
func (cmd WriteEntityCommand[I, E]) Execute(input I) (*E, *errors.CommandError) {
	// TODO add error handler for create entity
	entity := new(E)
	reflectutils.ReflectValues[E, I](entity, &input)
	// TODO handle error
	enti, _ := cmd.WriteFunc(*entity)
	return enti, nil
}
