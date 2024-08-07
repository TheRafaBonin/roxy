package roxy

import (
	"errors"
	"testing"

	"github.com/rotisserie/eris"
)

func TestInspection(t *testing.T) {
	t.Parallel()
	t.Run("testIs", testIs)
	t.Run("testIsFromRoxyNew", testIsFromRoxyNew)
	t.Run("testCause", testCause)
}

func testIs(t *testing.T) {
	baseError := errors.New("BaseError")
	wrappedError := Wrap(baseError, "Another Error")
	wrappedError = eris.Wrap(wrappedError, "Another Error 2")
	wrappedError = Wrap(wrappedError, "Another Error 3")
	wrappedError = SetDefaultGrpcResponse(wrappedError, GrpcResponse{})

	valid := Is(wrappedError, baseError)
	if !valid {
		t.Error("wrapped error should be baseError")
	}

	notValid := Is(wrappedError, errors.New("Not BaseError"))
	if notValid {
		t.Error("wrapped error should not be Not BaseError")
	}
}

func testIsFromRoxyNew(t *testing.T) {
	baseError := New("BaseError")
	wrappedError := Wrap(baseError, "Another Error")
	wrappedError = eris.Wrap(wrappedError, "Another Error 2")
	wrappedError = Wrap(wrappedError, "Another Error 3")
	wrappedError = SetDefaultGrpcResponse(wrappedError, GrpcResponse{})

	valid := Is(wrappedError, baseError)
	if !valid {
		t.Error("wrapped error should be baseError")
	}

	notValid := Is(wrappedError, New("Not BaseError"))
	if notValid {
		t.Error("wrapped error should not be Not BaseError")
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
