package roxy

import (
	"reflect"
)

// DefaultMessageResponse values
const (
	DropMessageAction MessageAction = iota
	SuccessMessageAction
	RequeueMessageAction
	DeadLetterMessageAction
)

// SetDefaultMessageAction sets the error message response to a defined response
func SetDefaultMessageAction(err error, response MessageAction) error {
	if err == nil {
		return nil
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		err = new(err)
	}

	eDetailedError := err.(*detailedError)

	(*eDetailedError).defaultMessageAction = &response
	return eDetailedError
}

// GetDefaultMessageAction gets the error set message response
func GetDefaultMessageAction(err error) MessageAction {
	if err == nil {
		return SuccessMessageAction
	}

	currentMessageAction := DeadLetterMessageAction

	var ok bool = true
	var u interface {
		Unwrap() error
	}
	for ok {
		u, ok = err.(interface {
			Unwrap() error
		})
		if ok {
			err = u.Unwrap()
		}

		detailedErr, valid := u.(*detailedError)
		if valid && detailedErr.defaultMessageAction != nil {
			return *detailedErr.defaultMessageAction
		}
	}

	return currentMessageAction
}
