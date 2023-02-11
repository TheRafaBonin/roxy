package roxy

import (
	"errors"
	"testing"
)

func TestDefinition(t *testing.T) {
	t.Parallel()
	t.Run("testError", testAttributes)
	t.Run("testUnwrap", testUnwrap)
}

func testError(t *testing.T) {
	errorMessage := "Base error"
	eDetailedError := detailedError{
		err: errors.New(errorMessage),
	}

	if eDetailedError.Error() != errorMessage {
		t.Errorf("%s should be equal to %s", eDetailedError.Error(), errorMessage)
	}
}

func testUnwrap(t *testing.T) {
	errorMessage := "Base error"
	err := errors.New(errorMessage)
	eDetailedError := detailedError{
		err: err,
	}

	if eDetailedError.Unwrap() != nil {
		t.Errorf("%v should be equal to %v", eDetailedError.Unwrap(), nil)
	}
}
