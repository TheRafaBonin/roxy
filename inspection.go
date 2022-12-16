package roxy

import "github.com/rotisserie/eris"

// Is ...
func Is(err error, target error) bool {
	return eris.Is(err, target)
}

// As ...
// TODO: Implement

// Cause ...
func Cause(err error, target error) error {
	return eris.Cause(err)
}
