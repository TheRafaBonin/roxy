package roxy

import (
	"errors"
	"testing"

	"github.com/rotisserie/eris"
)

func TestPublicError(t *testing.T) {
	t.Parallel()
	t.Run("TestSetPublicError", testSetPublicError)
	t.Run("TestGetPublicError", testGetPublicError)
	t.Run("TestSetPublicErrorNil", testSetPublicErrorNil)
	t.Run("TestGetCustomPublicError", testGetCustomPublicError)
}

func testSetPublicError(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	publicError := New("Public error")

	err = SetPublicError(err, publicError)

	eDetailedError, ok := err.(*detailedError)
	if !ok {
		t.Error("Could not cast to DetailedError")
	}
	if eDetailedError.publicErr == nil || eDetailedError.publicErr != publicError {
		t.Errorf("Responses do not match; %v; %v", eDetailedError.publicErr, publicError)
	}
}

func testGetPublicError(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultPublicError := UnexpectedPublicError

	err = Wrap(err, "Another error 1")
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")

	publicError := GetPublicError(err, nil)
	if !Is(publicError, defaultPublicError) {
		t.Errorf("Responses do not match; %v; %v", defaultPublicError, publicError)
	}
}

func testSetPublicErrorNil(t *testing.T) {
	t.Parallel()

	var err error
	err = SetPublicError(err, New("New public error"))

	if err != nil {
		t.Error("Error should be nil")
	}
}

func testGetCustomPublicError(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	latestPublicError := New("Latest Error")

	err = Wrap(err, "Another error 1")
	err = SetPublicError(err, New("First Public error"))
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")
	err = SetPublicError(err, latestPublicError)
	err = wrap(err, "Another error 4")
	err = eris.Wrap(err, "Another error 5")

	publicError := GetPublicError(err, nil)
	if !Is(publicError, latestPublicError) {
		t.Errorf("Responses do not match; %v; %v", publicError, publicError)
	}
}
