package roxy

import (
	"reflect"
)

// UnexpectedPublicError ...
var UnexpectedPublicError error = New("Unexpected Error")

// SetPublicError ...
func SetPublicError(err error, publicErr error) error {
	if err == nil {
		return nil
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		err = new(err)
	}

	eDetailedError := err.(*detailedError)
	(*eDetailedError).publicErr = publicErr
	return eDetailedError
}

// GetPublicError ...
func GetPublicError(err error, defaultMessage *string) (currentPublicError error) {
	// NIL case
	if err == nil {
		return currentPublicError
	}

	// Default case
	if defaultMessage == nil {
		currentPublicError = UnexpectedPublicError
	} else {
		currentPublicError = New(*defaultMessage)
	}

	// Named case
	var ok bool = true
	var u interface {
		Unwrap() error
	}
	for ok {
		u, ok = err.(interface {
			Unwrap() error
		})
		if ok {
			err = u.Unwrap()
		}

		detailedErr, valid := u.(*detailedError)
		if valid && detailedErr.publicErr != nil {
			return detailedErr.publicErr
		}
	}

	return currentPublicError
}
