package roxy

import (
	"errors"
	"net/http"
	"testing"
)

func TestHTTPResponse(t *testing.T) {
	t.Parallel()
	t.Run("TestSetDefaultHTTPResponse", testSetDefaultHTTPResponse)
	t.Run("TestGetDefaultHTTPResponse", testGetDefaultHTTPResponse)
	t.Run("TestGetCustomDefaultHTTPResponse", testGetCustomDefaultHTTPResponse)
}

func testSetDefaultHTTPResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultHTTPResponse := HTTPResponse{
		Message: http.StatusText(http.StatusOK),
		Status:  http.StatusOK,
	}

	err = SetDefaultHTTPResponse(err, defaultHTTPResponse)

	eDetailedError, ok := err.(*detailedError)
	if !ok {
		t.Error("Could not cast to DetailedError")
	}
	if (*eDetailedError.defaultHTTPResponse != defaultHTTPResponse && *eDetailedError.defaultHTTPResponse != HTTPResponse{}) {
		t.Errorf("Responses do not match; %v; %v", *eDetailedError.defaultHTTPResponse, defaultHTTPResponse)
	}
}

func testGetDefaultHTTPResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultHTTPResponse := HTTPResponse{
		Message: http.StatusText(http.StatusInternalServerError),
		Status:  http.StatusInternalServerError,
	}

	err = Wrap(err, "Another error 1")
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")

	httpResponse := GetDefaultHTTPResponse(err)
	if httpResponse != defaultHTTPResponse {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultHTTPResponse)
	}
}

func testGetCustomDefaultHTTPResponse(t *testing.T) {
	t.Parallel()

	defaultHTTPResponse := HTTPResponse{
		Message: http.StatusText(http.StatusBadRequest),
		Status:  http.StatusBadRequest,
	}

	err := errors.New("Root error")
	err = SetDefaultHTTPResponse(err, HTTPResponse{
		Message: http.StatusText(http.StatusNotFound),
		Status:  http.StatusNotFound,
	})

	err = Wrap(err, "Another error 1")
	err = Wrap(err, "Another error 2")
	err = SetDefaultHTTPResponse(err, defaultHTTPResponse)
	err = Wrap(err, "Another error 3")

	httpResponse := GetDefaultHTTPResponse(err)
	if httpResponse != defaultHTTPResponse {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultHTTPResponse)
	}
}
