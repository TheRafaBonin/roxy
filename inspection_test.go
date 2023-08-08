package roxy

import (
	"errors"
	"testing"
)

func TestInspection(t *testing.T) {
	t.Parallel()
	t.Run("testIs", testIs)
	t.Run("testCause", testCause)
}

func testIs(t *testing.T) {
	baseError := errors.New("BaseError")
	wrappedError := Wrap(baseError, "Another Error")

	valid := Is(wrappedError, baseError)
	if !valid {
		t.Error("wrapped error should be baseError")
	}
}

func testCause(t *testing.T) {
	baseError := errors.New("BaseError")
	wrappedError := Wrap(baseError, "Another Error")

	err := Cause(wrappedError)
	if err != baseError {
		t.Error("wrapped error should be caused by baseError")
	}
}
