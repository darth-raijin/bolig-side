package utility

import (
	errorDto "github.com/darth-raijin/bolig-side/api/models/dtos/error"

	"github.com/go-playground/validator/v10"
)

func CreateValidationSlice(errors validator.ValidationErrors) []errorDto.DomainError {
	var slice []errorDto.DomainError
	for _, error := range errors {
		slice = append(slice, errorDto.DomainError{
			Message: error.Error(),
		})
	}
	return slice
}
