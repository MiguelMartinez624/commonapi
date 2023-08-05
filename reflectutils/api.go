package reflectutils

import "reflect"

// TODO move to reflection utilities
// cache the types
var cache = map[reflect.Type][]string{}

func ReflectValues[E any, C any](target *E, change *C) {

	changesType := reflect.TypeOf(change).Elem()
	changesValue := reflect.ValueOf(change).Elem()
	targetValue := reflect.ValueOf(target).Elem()
	// trying with cache
	if fieldsList, ok := cache[changesType]; ok {
		for _, fieldName := range fieldsList {
			if f := targetValue.FieldByName(fieldName); f.IsValid() {
				f.Set(changesValue.FieldByName(fieldName))
			}
		}
		return
	}
	fieldsCount := changesType.NumField()
	fieldsToCache := make([]string, fieldsCount)
	for i := 0; i < fieldsCount; i++ {
		changeField := changesType.Field(i)
		// TODO check type
		// TODO Add ingore empty
		// TODO make this tags names constants
		if f := targetValue.FieldByName(changeField.Tag.Get("changes")); f.IsValid() {
			// Return better error when the field its not from the same type
			f.Set(changesValue.FieldByName(changeField.Name))
			fieldsToCache[i] = changeField.Name
		}
	}
	cache[changesType] = fieldsToCache

}
