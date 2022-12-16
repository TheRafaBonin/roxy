package roxy

import (
	"fmt"

	"github.com/rotisserie/eris"
)

func wrap(err error, msg string) error {
	eDetailedError, ok := err.(*detailedError)
	if !ok {
		return new(eris.Wrap(err, msg))
	}

	wrappedErr := eris.Wrap(eDetailedError.err, msg)
	cpError := detailedError(*eDetailedError)
	cpError.err = wrappedErr

	return &cpError
}

// Wrap ...
func Wrap(err error, msg string) error {
	return wrap(err, fmt.Sprint(msg))
}

// Wrapf ...
func Wrapf(err error, format string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(format, args...))
}

func unwrap(err error) error {
	u, ok := err.(interface {
		Unwrap() error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}

// Unwrap ...
func Unwrap(err error) error {
	eDetailedError, ok := err.(*detailedError)
	if !ok {
		return unwrap(err)
	}

	return unwrap(eDetailedError.err)
}
