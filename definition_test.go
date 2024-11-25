package roxy

import (
	"errors"
	"testing"
)

type customError struct {
	err error
}

func (ce customError) Error() string {
	return ce.err.Error()
}
func (ce customError) Unwrap() error {
	return ce.err
}

func TestDefinition(t *testing.T) {
	t.Parallel()
	t.Run("testError", testError)
	t.Run("testUnwrap", testUnwrap)
	t.Run("testErrorIs", testErrorIs)
	t.Run("testErrorAs", testErrorAs)
}

func testError(t *testing.T) {
	errorMessage := "Base error"
	eDetailedError := &detailedError{
		err: errors.New(errorMessage),
	}

	if eDetailedError.Error() != errorMessage {
		t.Errorf("%s should be equal to %s", eDetailedError.Error(), errorMessage)
	}
}

func testUnwrap(t *testing.T) {
	errorMessage := "Base error"
	err := errors.New(errorMessage)
	eDetailedError := &detailedError{
		err: err,
	}

	if eDetailedError.Unwrap() != err {
		t.Errorf("%v should be equal to %v", eDetailedError.Unwrap(), err)
	}
}

func testErrorIs(t *testing.T) {
	errorMessage := "Base error"
	err := errors.New(errorMessage)
	eDetailedError := &detailedError{
		err: err,
	}

	if !errors.Is(eDetailedError, err) {
		t.Errorf("errors.Is(%v, %v) should be true", eDetailedError, err)
	}

	err2 := errors.New("another error")
	if errors.Is(eDetailedError, err2) {
		t.Errorf("errors.Is(%v, %v) should be false", eDetailedError, err2)
	}

	err3 := errors.New(errorMessage)
	cErr3 := customError{err: err3}
	detailedErr3 := &detailedError{err: cErr3}
	if !errors.Is(detailedErr3, err3) {
		t.Errorf("errors.Is(%v, %v) should be true", detailedErr3, err3)
	}
	if !errors.Is(detailedErr3, cErr3) {
		t.Errorf("errors.Is(%v, %v) should be true", detailedErr3, cErr3)
	}

	err4 := errors.New("another error")
	cErr4 := customError{err: err4}
	detailedErr4 := &detailedError{err: cErr4}
	if errors.Is(detailedErr4, err3) {
		t.Errorf("errors.Is(%v, %v) should be false", detailedErr4, err3)
	}
	if errors.Is(detailedErr4, cErr3) {
		t.Errorf("errors.Is(%v, %v) should be false", detailedErr4, cErr3)
	}
}

func testErrorAs(t *testing.T) {
	errorMessage := "Base error"
	err := errors.New(errorMessage)
	cErr := customError{err: err}
	eDetailedError := &detailedError{
		err: cErr,
	}

	var cErr2 customError
	if !errors.As(eDetailedError, &cErr2) {
		t.Errorf("errors.As(%v, %v) should be true", eDetailedError, cErr2)
	}
}
