package roxy

import (
	"testing"

	"google.golang.org/grpc/codes"
)

func TestAttributes(t *testing.T) {
	t.Parallel()
	t.Run("testError", testAttributes)
}

func testAttributes(t *testing.T) {
	err := Errorf("New %s", "error")
	err = SetErrorLogLevel(err, WarnLevel)
	err = SetDefaultGrpcResponse(err, codes.Canceled)
	err = SetDefaultMessageAction(err, DropMessageAction)

	logLevel := GetErrorLogLevel(err)
	if logLevel != WarnLevel {
		t.Errorf("%v should be equal to %v", logLevel, WarnLevel)
	}

	grpcResponse := GetDefaultGrpcResponse(err)
	if grpcResponse != codes.Canceled {
		t.Errorf("%v should be equal to %v", grpcResponse, codes.Canceled)
	}

	messageAction := GetDefaultMessageAction(err)
	if messageAction != DropMessageAction {
		t.Errorf("%v should be equal to %v", messageAction, DropMessageAction)
	}
}
