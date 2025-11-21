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
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(e)
}

func GetPaginationQuery(c *gin.Context, defaultLimit, defaultOffset int) (int, int, error) {
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

	return limit, offset, nil
}

func DetectLoginType(id string) string {
	var (
		nidnRegex = regexp.MustCompile(`^0\d{9}$`)
		nimRegex  = regexp.MustCompile(`^(G1A0\d{5}|Y1G0\d{5})$`)
	)

	switch {
	case nidnRegex.MatchString(id):
		return "nidn"
	case nimRegex.MatchString(id):
		return "nim"
	default:
		return "username"
	}
}

func ContainsNumber(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
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
