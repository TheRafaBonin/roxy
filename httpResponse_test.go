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
	if eDetailedError.defaultHTTPResponse != defaultHTTPResponse {
		t.Errorf("Responses do not match; %v; %v", eDetailedError.defaultHTTPResponse, defaultHTTPResponse)
	}
}

func testGetDefaultHTTPResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultHTTPResponse := HTTPResponse{
		Message: http.StatusText(http.StatusInternalServerError),
		Status:  http.StatusInternalServerError,
	}

	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")

	httpResponse := GetDefaultHTTPResponse(err)
	if httpResponse != defaultHTTPResponse {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultHTTPResponse)
	}
}

func testGetCustomDefaultHTTPResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultHTTPResponse := HTTPResponse{
		Message: http.StatusText(http.StatusOK),
		Status:  http.StatusOK,
	}

	err = Wrap(err, "Another error")
	err = SetDefaultHTTPResponse(err, defaultHTTPResponse)
	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")

	httpResponse := GetDefaultHTTPResponse(err)
	if httpResponse != defaultHTTPResponse {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultHTTPResponse)
	}
}
