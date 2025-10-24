package helper

import (
	"fmt"
	"reflect"
)

func GetFieldValue(data interface{}, jsonTag string) interface{} {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag == jsonTag {
			return val.Field(i).Interface()
		}
	}
	return nil
}

func ValidateFieldLengths(data interface{}, rules map[string]int) error {
	for field, max := range rules {
		val := GetFieldValue(data, field)
		if str, ok := val.(string); ok && len(str) > max {
			return fmt.Errorf("%s too long (max %d)", field, max)
		}
	}
	return nil
}
