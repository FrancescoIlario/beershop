package beershop

import "fmt"

// ValidationResult the result of a validation operation
type ValidationResult interface {
	// Errors returns the errors
	Errors() map[string]string

	// IsValid checks whether the ValidationResult is valid
	IsValid() bool
}

type validationResult struct {
	errors map[string]string
}

func (v *validationResult) IsValid() bool {
	return v == nil || len(v.Errors()) == 0
}

func (v *validationResult) Errors() map[string]string {
	return v.errors
}

// ErrValidationFailed this error is returned when validation fails
var ErrValidationFailed = fmt.Errorf("validation failed")
