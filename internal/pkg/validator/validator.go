package validator

import (
	"strings"

	v10 "github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *v10.Validate
}

func New() *Validator {
	return &Validator{
		validate: v10.New(),
	}
}

func (v *Validator) ValidateStruct(s interface{}) map[string]string {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	result := map[string]string{}
	for _, e := range err.(v10.ValidationErrors) {
		field := strings.ToLower(e.Field())
		switch e.Tag() {
		case "required":
			result[field] = "field is required"
		case "oneof":
			result[field] = "invalid value"
		case "gt":
			result[field] = "must be greater than zero"
		case "min":
			result[field] = "minimum value is not met"
		default:
			result[field] = "invalid value"
		}
	}

	return result
}
