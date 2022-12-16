package roxy

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("testNew", testNew)
	t.Run("testErrorf", testErrorf)
}

func testNew(t *testing.T) {
	baseString := "BaseError"
	err := New(baseString)

	stringError := err.Error()
	if stringError != baseString {
		t.Errorf("%s is not equal to %s", stringError, baseString)
	}

	_, ok := err.(*detailedError)
	if !ok {
		t.Error("Could not cast to DetailedError")
	}
}

func testErrorf(t *testing.T) {
	baseString := "BaseError"
	err := Errorf("%sError", "Base")

	stringError := err.Error()
	if stringError != baseString {
		t.Errorf("%s is not equal to %s", stringError, baseString)
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		t.Error("Could not cast to DetailedError")
	}
}
