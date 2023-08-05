package common

import (
	"reflect"
)

type Filter struct{}

type Filterable interface {
	GetMap(root any) map[string]any
}

// TODO Tags solo serian para aplicar operadores igual a (que seria el por defecto) oo mayorq ue contains start with y asi
// els istema q interpreta y mapeo esto a el gestor de base, de momento memoria
func (q Filter) GetMap(root any) map[string]any {
	immutable := reflect.ValueOf(root)
	typeStruct := reflect.TypeOf(root)
	//v := reflect.ValueOf(&structType).Elem()
	queryParms := map[string]any{}
	for i := 0; i < immutable.NumField(); i++ {

		fieldValue := immutable.Field(i)
		fieldType := typeStruct.Field(i)

		// This case are ignore as are not fields for filter
		// one per user choice and the other one is the Filter composition struct
		if fieldType.Tag.Get("filter") == "ignore" {
			continue
		}
		if fieldType.Type == reflect.TypeOf(Filter{}) {
			continue
		}
		// zero values will be ignored too, should they?
		if fieldValue.IsZero() {
			continue
		}
		//log.Println(fieldValue)
		queryParms[fieldType.Name] = fieldValue

	}
	// need  access top the parent
	//log.Println(queryParms)
	return queryParms
}
