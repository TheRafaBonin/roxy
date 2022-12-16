package roxy

import (
	"errors"
	"testing"

	"google.golang.org/grpc/codes"
)

func TestGrpcResponse(t *testing.T) {
	t.Parallel()
	t.Run("TestSetDefaultGrpcResponse", testSetDefaultGrpcResponse)
	t.Run("TestGetDefaultGrpcResponse", testGetDefaultGrpcResponse)
	t.Run("TestGetCustomDefaultGrpcResponse", testGetCustomDefaultGrpcResponse)
}

func testSetDefaultGrpcResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultGrpcResponse := codes.OK

	err = SetDefaultGrpcResponse(err, defaultGrpcResponse)

	eDetailedError, ok := err.(*detailedError)
	if !ok {
		t.Error("Could not cast to DetailedError")
	}
	if eDetailedError.defaultGrpcResponse == nil || *eDetailedError.defaultGrpcResponse != defaultGrpcResponse {
		t.Errorf("Responses do not match; %v; %v", eDetailedError.defaultGrpcResponse, defaultGrpcResponse)
	}
}

func testGetDefaultGrpcResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultGrpcResponse := codes.Internal

	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")

	httpResponse := GetDefaultGrpcResponse(err)
	if httpResponse != defaultGrpcResponse {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultGrpcResponse)
	}
}

func testGetCustomDefaultGrpcResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultGrpcResponse := codes.OK

	err = Wrap(err, "Another error")
	err = SetDefaultGrpcResponse(err, defaultGrpcResponse)
	err = Wrap(err, "Another error")
	err = Wrap(err, "Another error")

	httpResponse := GetDefaultGrpcResponse(err)
	if httpResponse != defaultGrpcResponse {
		t.Errorf("Responses do not match; %v; %v", httpResponse, defaultGrpcResponse)
	}
}
