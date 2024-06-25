package helpers

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = nil
var lock = &sync.Mutex{}

func getInvalidFromRequest(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "gt":
		return "the value of " + err.Field() + " must be greater than " + err.Param()
	case "eq":
		return "the value of " + err.Field() + " must be equal to " + err.Param()
	case "gte":
		return "the value of " + err.Field() + " must be greater than or equal " + err.Param()
	case "email":
		return "the email is invalid"
	default:
		return "validation error in " + err.Field()
	}
}

func getInstance() *validator.Validate {
	if validate == nil {
		lock.Lock()
		defer lock.Unlock()
		if validate == nil {
			validate = validator.New(validator.WithRequiredStructEnabled())
		}
	}
	return validate
}

func ValidateStruct[T any](data T) []ErrorRequest {
	var errors []ErrorRequest
	if err := getInstance().Struct(data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorRequest
			element.ErrorMessage = getInvalidFromRequest(err)
			element.Field = err.Field()
			errors = append(errors, element)
		}
	}
	return errors
}
