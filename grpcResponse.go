package roxy

import (
	"reflect"

	"google.golang.org/grpc/codes"
)

// UnhandledErrorGrpcResponse defines a default response for an error that
// has not been attributed any other gRPC response
var UnhandledErrorGrpcResponse = GrpcResponse{
	Message: "internal server error",
	Code:    codes.Internal,
}

// OKGrpcResponse defines the default gRPC OK response for nil errors
var OKGrpcResponse = GrpcResponse{
	Code:    codes.OK,
	Message: "",
}

// SetDefaultGrpcResponse sets the error gRPC response to a defined response
func SetDefaultGrpcResponse(err error, response GrpcResponse) error {
	if err == nil {
		return nil
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		err = new(err)
	}

	eDetailedError := err.(*detailedError)
	(*eDetailedError).defaultGrpcResponse = &response
	return eDetailedError
}

// GetDefaultGrpcResponse gets the error set gRPC response
func GetDefaultGrpcResponse(err error) GrpcResponse {
	if err == nil {
		return OKGrpcResponse
	}

	currentGrpcResponse := UnhandledErrorGrpcResponse

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
		if valid && detailedErr.defaultGrpcResponse != nil {
			return *detailedErr.defaultGrpcResponse
		}
	}

	return currentGrpcResponse
}
