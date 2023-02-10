package roxy

import (
	"net/http"
	"reflect"
)

// SetDefaultHTTPResponse ...
func SetDefaultHTTPResponse(err error, response HTTPResponse) error {
	if err == nil {
		return nil
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(detailedError{}) {
		err = new(err)
	}

	eDetailedError := err.(*detailedError)

	(*eDetailedError).defaultHTTPResponse = &response
	return eDetailedError
}

// GetDefaultHTTPResponse ...
func GetDefaultHTTPResponse(err error) HTTPResponse {
	if err == nil {
		return HTTPResponse{
			Message: http.StatusText(http.StatusOK),
			Status:  http.StatusOK,
		}
	}

	currentHTTPResponse := HTTPResponse{
		Message: http.StatusText(http.StatusInternalServerError),
		Status:  http.StatusInternalServerError,
	}

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
