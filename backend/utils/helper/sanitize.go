package helper

import (
	"reflect"
	"strings"
	"time"
)

func SanitizeSensitiveFields(data interface{}, sensitiveFields []string) interface{} {
	if data == nil {
		return nil
	}

	sensitiveMap := make(map[string]bool)
	for _, field := range sensitiveFields {
		sensitiveMap[strings.ToLower(field)] = true
	}

	return sanitizeValue(reflect.ValueOf(data), sensitiveMap)
}

func sanitizeValue(value reflect.Value, sensitiveMap map[string]bool) interface{} {
	if value.Type() == reflect.TypeOf(time.Time{}) {
		return value.Interface()
	}

	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			return nil
		}
		return sanitizeValue(value.Elem(), sensitiveMap)
	case reflect.Struct:
		return sanitizeStruct(value, sensitiveMap)
	case reflect.Slice:
		return sanitizeSlice(value, sensitiveMap)
	case reflect.Map:
		return sanitizeMap(value, sensitiveMap)
	default:
		return value.Interface()
	}
}

func sanitizeStruct(v reflect.Value, sensitiveMap map[string]bool) interface{} {
	t := v.Type()
	result := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if !field.IsExported() {
			continue
		}

		fieldName := field.Name
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			fieldName = strings.Split(jsonTag, ",")[0]
		}

		if sensitiveMap[strings.ToLower(fieldName)] || sensitiveMap[strings.ToLower(field.Name)] {
			continue
		}

		sanitizedValue := sanitizeValue(fieldValue, sensitiveMap)
		result[fieldName] = sanitizedValue
	}

	return result
}

func sanitizeSlice(v reflect.Value, sensitiveMap map[string]bool) interface{} {
	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = sanitizeValue(v.Index(i), sensitiveMap)
	}
	return result
}

func sanitizeMap(v reflect.Value, sensitiveMap map[string]bool) interface{} {
	result := make(map[string]interface{})
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)

		keyStr := key.String()
		if key.Kind() == reflect.String {
			keyStr = key.String()
		} else {
			keyStr = key.Interface().(string)
		}

		if sensitiveMap[strings.ToLower(keyStr)] {
			continue
		}

		result[keyStr] = sanitizeValue(val, sensitiveMap)
	}
	return result
}

func SanitizeUserResponse(user interface{}) interface{} {
	//Data sensitif
	sensitiveFields := []string{
		"password", "Password", //bisa tambah refreshtoken
		"salt", "Salt",
		"verification_token", "VerificationToken",
	}
	return SanitizeSensitiveFields(user, sensitiveFields)
}

func SanitizeQuestionResponse(question interface{}) interface{} {
	sensitiveFields := []string{}
	return SanitizeSensitiveFields(question, sensitiveFields)
}
