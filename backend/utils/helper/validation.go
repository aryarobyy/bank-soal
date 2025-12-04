package helper

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
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

func IsValidSubjectTitle(title model.SubjectTitle) bool {
	switch title {
	case model.SubjectKalkulus,
		model.SubjectMatDis,
		model.SubjectAutomata,
		model.SubjectData,
		model.SubjectMetNum:
		return true
	default:
		return false
	}
}

func IsValidEmail(e string) bool {
	if e == "" {
		return false
	}

	for _, r := range e {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '@' || r ==
			'.' || r == '_' || r == '-') {
			return false
		}
	}
	return true
}

func GetPaginationQuery(c *gin.Context, defaultLimit int, defaultOffset int) (int, int, error) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit := defaultLimit
	offset := defaultOffset

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	if limit > 200 {
		return 0, 0, fmt.Errorf("wowww banyak banget bro")
	}

	return limit, offset, nil
}

func DetectLoginType(id string) string {
	var (
		nipRegex = regexp.MustCompile(`^(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])(19|20)\d{2}(0[1-9]|1[0-2])[12]\d{3}$`)
		nimRegex = regexp.MustCompile(`^(G1A0\d{5}|Y1G0\d{5})$`)
	)

	switch {
	case nipRegex.MatchString(id):
		return "nip"
	case nimRegex.MatchString(id):
		return "nim"
	default:
		return "username"
	}
}

func IsValidName(s string) bool {
	if s == "" {
		return false
	}

	for _, ch := range s {
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == ' ') {
			return false
		}
	}
	return true
}

func BindAndConvertToPtr(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func BindAndConvertToBoolPtr(value string) *bool {
	if value == "" {
		return nil
	}
	if value == "true" || value == "1" {
		t := true
		return &t
	}
	if value == "false" || value == "0" {
		f := false
		return &f
	}
	return nil
}

func BindToInt(value string) int {
	if value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return 0
}

func BindToIntPtr(value string) *int {
	if value != "" {
		if v, err := strconv.Atoi(value); err == nil {
			return &v
		}
	}

	return nil
}

func ValidateAndFilterUserData(userData map[string]interface{}) map[string]interface{} {
	sensitiveFields := map[string]bool{
		"password": true,
		"salt":     true,
	}

	filteredData := make(map[string]interface{})
	for key, value := range userData {
		if !sensitiveFields[key] {
			filteredData[key] = value
		}
	}
	return filteredData
}
