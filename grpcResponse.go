package roxy

import (
	"reflect"

	"google.golang.org/grpc/codes"
)

// SetDefaultGrpcResponse ...
func SetDefaultGrpcResponse(err error, response codes.Code) error {
	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(detailedError{}) {
		err = new(err)
	}

	eDetailedError := err.(*detailedError)
	(*eDetailedError).defaultGrpcResponse = &response
	return eDetailedError
}

// GetDefaultGrpcResponse ...
func GetDefaultGrpcResponse(err error) codes.Code {
	currentGrpcResponse := codes.Internal

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
			currentGrpcResponse = *detailedErr.defaultGrpcResponse
		}
	}

	return currentGrpcResponse
}
