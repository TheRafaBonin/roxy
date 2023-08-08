package roxy

import (
	"net/http"
	"reflect"
)

// OkHTTPResponse defines the default HTTP OK response for nil errors
var OkHTTPResponse = HTTPResponse{
	Message: http.StatusText(http.StatusOK),
	Status:  http.StatusOK,
}

// UnhandledHTTPResponse defines a default response for an error that
// has not been attributed any other HTTP response
var UnhandledHTTPResponse = HTTPResponse{
	Message: http.StatusText(http.StatusInternalServerError),
	Status:  http.StatusInternalServerError,
}

// SetDefaultHTTPResponse sets the error HTTP response to a defined response
func SetDefaultHTTPResponse(err error, response HTTPResponse) error {
	if err == nil {
		return nil
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		err = new(err)
	}

	eDetailedError := err.(*detailedError)

	(*eDetailedError).defaultHTTPResponse = &response
	return eDetailedError
}

// GetDefaultHTTPResponse gets the error set HTTP response
func GetDefaultHTTPResponse(err error) HTTPResponse {
	if err == nil {
		return OkHTTPResponse
	}

	currentHTTPResponse := UnhandledHTTPResponse

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
		if valid && detailedErr.defaultHTTPResponse != nil {
			return *detailedErr.defaultHTTPResponse
		}
	}

	return currentHTTPResponse
}
