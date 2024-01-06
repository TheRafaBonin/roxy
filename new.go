package roxy

import (
	"github.com/rotisserie/eris"
)

func new(err error) error {
	if err == nil {
		return nil
	}

	return &detailedError{
		err: err,
	}
}

// New ...
func New(msg string, opts ...NewErrorOption) error {
	newError := eris.New(msg)
	for _, opt := range opts {
		newError = opt(newError)
	}

	return new(newError)
}

// Errorf ...
func Errorf(format string, args ...interface{}) error {
	newError := eris.Errorf(format, args...)
	return new(newError)
}
