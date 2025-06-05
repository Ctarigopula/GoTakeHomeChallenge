package coordinators

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	gorm2 "gorm.io/gorm"
)

type validationError struct {
	message string
}

func (e validationError) Error() string {
	return e.message
}

// IsValidationError returned a boolean value indicating whether or not the
// error is an validation error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}

	_, isValidationError := errors.Cause(err).(*validationError)
	return isValidationError
}

// IsRecordNotFoundError returns a boolean value indicating whether or not the
// error is a record not found error.
func IsRecordNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return gorm.IsRecordNotFoundError(errors.Cause(err)) || errors.Is(err, gorm2.ErrRecordNotFound)
}
