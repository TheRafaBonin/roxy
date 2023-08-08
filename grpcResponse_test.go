package roxy

import (
	"errors"
	"testing"

	"github.com/rotisserie/eris"
	"google.golang.org/grpc/codes"
)

func TestGrpcResponse(t *testing.T) {
	t.Parallel()
	t.Run("TestSetDefaultGrpcResponse", testSetDefaultGrpcResponse)
	t.Run("TestGetDefaultGrpcResponse", testGetDefaultGrpcResponse)
	t.Run("TestSetDefaultGrpcResponseNil", testSetDefaultGrpcResponseNil)
	t.Run("TestGetCustomDefaultGrpcResponse", testGetCustomDefaultGrpcResponse)
}

func testSetDefaultGrpcResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultGrpcResponse := GrpcResponse{
		Code: codes.Canceled,
	}

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
	defaultGrpcResponse := GrpcResponse{
		Message: "internal server error",
		Code:    codes.Internal,
	}

	err = Wrap(err, "Another error 1")
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")

	grpcResponse := GetDefaultGrpcResponse(err)
	if grpcResponse != defaultGrpcResponse {
		t.Errorf("Responses do not match; %v; %v", grpcResponse, defaultGrpcResponse)
	}
}

func testSetDefaultGrpcResponseNil(t *testing.T) {
	t.Parallel()

	var err error
	err = SetDefaultGrpcResponse(err, GrpcResponse{
		Code: codes.Canceled,
	})

	if err != nil {
		t.Error("Error should be nil")
	}
}

func testGetCustomDefaultGrpcResponse(t *testing.T) {
	t.Parallel()

	err := errors.New("Root error")
	defaultGrpcResponse := GrpcResponse{
		Code: codes.NotFound,
	}

	err = Wrap(err, "Another error 1")
	err = SetDefaultGrpcResponse(err, GrpcResponse{
		Code: codes.Aborted,
	})
	err = Wrap(err, "Another error 2")
	err = Wrap(err, "Another error 3")
	err = SetDefaultGrpcResponse(err, defaultGrpcResponse)
	err = wrap(err, "Another error 4")
	err = eris.Wrap(err, "Another error 5")

	grpcResponse := GetDefaultGrpcResponse(err)
	if grpcResponse != defaultGrpcResponse {
		t.Errorf("Responses do not match; %v; %v", grpcResponse, defaultGrpcResponse)
	}
}
