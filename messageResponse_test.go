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

	err = Wrap(err, "Another error 1")
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")

	httpResponse := GetDefaultMessageAction(err)
	if httpResponse != defaultMessageAction {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultMessageAction)
	}
}

func testGetCustomDefaultMessageAction(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultMessageAction := RequeueMessageAction
	err = SetDefaultMessageAction(err, DropMessageAction)

	err = Wrap(err, "Another error 1")
	err = Wrap(err, "Another error 2")
	err = SetDefaultMessageAction(err, defaultMessageAction)
	err = Wrap(err, "Another error 3")

	httpResponse := GetDefaultMessageAction(err)
	if httpResponse != defaultMessageAction {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultMessageAction)
	}
}
