package innerstruct

import (
	"encoding/json"
	"reflect"
)

func structToMap(obj interface{}) map[string]interface{} {
	ret := map[string]interface{}{}

	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		tag := field.Tag.Get("field")
		ret[field.Name] = reflect.Indirect(objValue).FieldByName(field.Name).FieldByName(tag)
	}

	return ret
}

func InnerStructToJson(obj interface{}) []byte {
	mappedObj := structToMap(obj)
	ret, _ := json.Marshal(mappedObj)
	return ret
}
