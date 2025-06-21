package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID string
func GenerateUUID() string {
	return uuid.New().String()
}

// StringPtr returns a pointer to the string value
func StringPtr(s string) *string {
	return &s
}

// TimePtr returns a pointer to the time value
func TimePtr(t time.Time) *time.Time {
	return &t
}

// Contains checks if a slice contains a value
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Filter filters a slice based on a predicate function
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map transforms a slice using a mapping function
func Map[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

// StructToMap converts a struct to a map
func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return result
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.IsExported() {
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" && jsonTag != "-" {
				fieldName := strings.Split(jsonTag, ",")[0]
				result[fieldName] = v.Field(i).Interface()
			} else {
				result[field.Name] = v.Field(i).Interface()
			}
		}
	}
	return result
}

// UpdateStructFromMap updates struct fields from a map
func UpdateStructFromMap(obj interface{}, updates map[string]interface{}) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("obj must be a pointer to struct")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		jsonTag := field.Tag.Get("json")
		var fieldName string
		if jsonTag != "" && jsonTag != "-" {
			fieldName = strings.Split(jsonTag, ",")[0]
		} else {
			fieldName = field.Name
		}

		if value, exists := updates[fieldName]; exists && value != nil {
			fieldValue := v.Field(i)
			if fieldValue.CanSet() {
				newValue := reflect.ValueOf(value)
				if newValue.Type().ConvertibleTo(fieldValue.Type()) {
					fieldValue.Set(newValue.Convert(fieldValue.Type()))
				}
			}
		}
	}

	return nil
}

// SliceToString converts a slice to a comma-separated string
func SliceToString[T any](slice []T) string {
	if len(slice) == 0 {
		return ""
	}

	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(strSlice, ",")
}

// StringToSlice converts a comma-separated string to a slice of strings
func StringToSlice(str string) []string {
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ",")
}