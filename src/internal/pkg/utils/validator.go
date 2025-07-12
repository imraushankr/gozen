package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func GetValidationErrors(err error) map[string]string {
	errs := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errs[e.Field()] = getValidationMessage(e)
		}
	}

	return errs
}

func getValidationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Must be at least " + e.Param() + " characters"
	case "max":
		return "Must be at most " + e.Param() + " characters"
	default:
		return "Invalid value"
	}
}