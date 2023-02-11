package roxy

import (
	"errors"
	"testing"

	"github.com/rotisserie/eris"
)

func TestLogLevel(t *testing.T) {
	t.Parallel()
	t.Run("TestSetLogLevel", testSetLogLevel)
	t.Run("TestGetLogLevel", testGetLogLevel)
	t.Run("TestSetLogLevelNil", testSetLogLevelNil)
	t.Run("TestGetCustomLogLevel", testGetCustomLogLevel)
}

func testSetLogLevel(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	baseLogLevel := InfoLevel

	err = SetErrorLogLevel(err, baseLogLevel)

	eDetailedError, ok := err.(*detailedError)
	if !ok {
		t.Error("Could not cast to DetailedError")
	}
	if eDetailedError.errLogLevel == nil || *eDetailedError.errLogLevel != baseLogLevel {
		t.Errorf("Responses do not match; %v; %v", eDetailedError.errLogLevel, baseLogLevel)
	}
}

func testGetLogLevel(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	baseLogLevel := ErrorLevel

	err = Wrap(err, "Another error 1")
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")

	logLevel := GetErrorLogLevel(err)
	if logLevel != logLevel {
		t.Errorf("Responses do not match; %v; %v", logLevel, baseLogLevel)
	}
}

func testSetLogLevelNil(t *testing.T) {
	t.Parallel()

	var err error
	err = SetErrorLogLevel(err, WarnLevel)

	if err != nil {
		t.Error("Error should be nil")
	}
}

func testGetCustomLogLevel(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	baseLogLevel := WarnLevel

	err = Wrap(err, "Another error 1")
	err = SetErrorLogLevel(err, DebugLevel)
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")
	err = SetErrorLogLevel(err, baseLogLevel)
	err = wrap(err, "Another error 4")
	err = eris.Wrap(err, "Another error 5")

	logLevel := GetErrorLogLevel(err)
	if logLevel != logLevel {
		t.Errorf("Responses do not match; %v; %v", logLevel, baseLogLevel)
	}
}
