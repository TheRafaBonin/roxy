package roxy

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
	t.Parallel()
	t.Run("testWrap", testWrap)
	t.Run("testWrapf", testWrapf)
	t.Run("testUnwarp", testWrapUnwrap)
}

func testWrap(t *testing.T) {
	baseString := "BaseError"
	anotherErrorString := "Another Error"
	baseError := errors.New(baseString)
	err := Wrap(baseError, anotherErrorString)

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		t.Error("Could not cast to DetailedError")
	}

	stringError := err.Error()
	compareString := fmt.Sprintf("%s: %s", anotherErrorString, baseString)
	if stringError != compareString {
		t.Errorf("%s is not equal to %s", stringError, compareString)
	}
}

func testWrapf(t *testing.T) {
	baseString := "BaseError"
	anotherErrorString := "Another Error"
	baseError := errors.New(baseString)

	splitString := strings.Split(anotherErrorString, " ")
	err := Wrapf(baseError, "%s %s", splitString[0], splitString[1])

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		t.Error("Could not cast to DetailedError")
	}

	stringError := err.Error()
	compareString := fmt.Sprintf("%s: %s", anotherErrorString, baseString)
	if stringError != compareString {
		t.Errorf("%s is not equal to %s", stringError, compareString)
	}
}

func testWrapUnwrap(t *testing.T) {
	baseString := "BaseError"
	anotherErrorString := "Another Error"
	baseError := errors.New(baseString)
	err1 := Wrap(baseError, anotherErrorString)
	err2 := Wrap(err1, anotherErrorString)

	err := Unwrap(err2)

	if err.Error() != err1.Error() {
		t.Errorf("%s is not equal to %s", err.Error(), baseError.Error())
	}
}
