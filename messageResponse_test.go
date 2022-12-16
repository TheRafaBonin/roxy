package roxy

import (
	"errors"
	"testing"
)

func TestMessageAction(t *testing.T) {
	t.Parallel()
	t.Run("TestSetDefaultMessageAction", testSetDefaultMessageAction)
	t.Run("TestGetDefaultMessageAction", testGetDefaultMessageAction)
	t.Run("TestGetCustomDefaultMessageAction", testGetCustomDefaultMessageAction)
}

func testSetDefaultMessageAction(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultMessageAction := SuccessMessageAction

	err = SetDefaultMessageAction(err, defaultMessageAction)

	eDetailedError, ok := err.(*detailedError)
	if !ok {
		t.Error("Could not cast to DetailedError")
	}
	if eDetailedError.defaultMessageAction == nil || *eDetailedError.defaultMessageAction != defaultMessageAction {
		t.Errorf("Responses do not match; %v; %v", eDetailedError.defaultMessageAction, defaultMessageAction)
	}
}

func testGetDefaultMessageAction(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultMessageAction := DeadLetterMessageAction

	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")

	httpResponse := GetDefaultMessageAction(err)
	if httpResponse != defaultMessageAction {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultMessageAction)
	}
}

func testGetCustomDefaultMessageAction(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultMessageAction := RequeueMessageAction

	err = Wrap(err, "Another error")
	err = SetDefaultMessageAction(err, defaultMessageAction)
	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")

	httpResponse := GetDefaultMessageAction(err)
	if httpResponse != defaultMessageAction {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultMessageAction)
	}
}
