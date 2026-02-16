package helper

import (
	"fmt"
	"reflect"
	"strings"
)

type DataFilter struct {
	allowedFields map[string][]string
	blacklist     map[string][]string
}

func NewDataFilter() *DataFilter {
	return &DataFilter{
		allowedFields: make(map[string][]string),
		blacklist:     make(map[string][]string),
	}
}

func (df *DataFilter) AddAllowedFields(typeName string, fields []string) {
	df.allowedFields[typeName] = fields
}

func (df *DataFilter) AddBlacklistedFields(typeName string, fields []string) {
	df.blacklist[typeName] = fields
}

func (df *DataFilter) FilterRequestData(input interface{}) (interface{}, error) {
	if input == nil {
		return nil, nil
	}

	if reflect.TypeOf(input).Kind() == reflect.Slice {
		return df.filterSliceData(input)
	}

	return df.filterSingleData(input)
}

func (df *DataFilter) filterSingleData(input interface{}) (interface{}, error) {
	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	if inputType.Kind() == reflect.Ptr {
		if inputValue.IsNil() {
			return nil, nil
		}
		inputValue = inputValue.Elem()
		inputType = inputType.Elem()
	}

	if inputType.Kind() != reflect.Struct {
		return input, nil
	}

	allowed, hasAllowList := df.allowedFields[inputType.Name()]
	blacklisted, hasBlackList := df.blacklist[inputType.Name()]

	result := reflect.New(inputType).Elem()

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fieldValue := inputValue.Field(i)

		if !field.IsExported() {
			continue
		}

		fieldName := field.Name
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			fieldName = strings.Split(jsonTag, ",")[0]
		}

		if hasBlackList {
			isBlacklisted := false
			for _, blacklistedField := range blacklisted {
				if blacklistedField == field.Name || blacklistedField == fieldName {
					isBlacklisted = true
					break
				}
			}
			if isBlacklisted {
				continue
			}
		}

		if hasAllowList {
			isAllowed := false
			for _, allowedField := range allowed {
				if allowedField == field.Name || allowedField == fieldName {
					isAllowed = true
					break
				}
			}
			if !isAllowed {
				continue
			}
		}

		if result.Field(i).CanSet() {
			result.Field(i).Set(fieldValue)
		}
	}

	return result.Interface(), nil
}

func (df *DataFilter) filterSliceData(input interface{}) (interface{}, error) {
	inputValue := reflect.ValueOf(input)
	if inputValue.Kind() != reflect.Slice {
		return input, nil
	}

	resultSlice := reflect.MakeSlice(inputValue.Type(), 0, inputValue.Len())

	for i := 0; i < inputValue.Len(); i++ {
		item := inputValue.Index(i).Interface()
		filteredItem, err := df.filterSingleData(item)
		if err != nil {
			return nil, err
		}
		resultSlice = reflect.Append(resultSlice, reflect.ValueOf(filteredItem))
	}

	return resultSlice.Interface(), nil
}

func (df *DataFilter) SanitizeResponseData(output interface{}) (interface{}, error) {
	return df.FilterRequestData(output)
}

func (df *DataFilter) ValidateAndFilterMap(input map[string]interface{}, allowedKeys []string) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range input {
		if len(allowedKeys) > 0 {
			isAllowed := false
			for _, allowedKey := range allowedKeys {
				if key == allowedKey {
					isAllowed = true
					break
				}
			}
			if !isAllowed {
				continue
			}
		}

		result[key] = value
	}

	return result
}

func (df *DataFilter) ValidateInputFields(input interface{}, expectedFields map[string]bool) error {
	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	if inputType.Kind() == reflect.Ptr {
		if inputValue.IsNil() {
			return nil
		}
		inputValue = inputValue.Elem()
		inputType = inputType.Elem()
	}

	if inputType.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)

		if !field.IsExported() {
			continue
		}

		fieldName := field.Name
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			fieldName = strings.Split(jsonTag, ",")[0]
		}

		if !expectedFields[fieldName] && !expectedFields[field.Name] {
			return fmt.Errorf("unexpected field '%s' in input data", fieldName)
		}
	}

	return nil
}