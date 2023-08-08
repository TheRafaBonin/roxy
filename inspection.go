package roxy

import "github.com/rotisserie/eris"

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if it implements a method
// Is(error) bool such that Is(target) returns true.
func Is(err error, target error) bool {
	return eris.Is(err, target)
}

// As ...
// TODO: Implement

// Cause returns the root cause of the error, which is defined as the first error in the chain. The original
// error is returned if it does not implement `Unwrap() error` and nil is returned if the error is nil.
func Cause(err error, target error) error {
	return eris.Cause(err)
}
