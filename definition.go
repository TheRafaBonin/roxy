package roxy

import "google.golang.org/grpc/codes"

// HTTPResponse defines a default http response
type HTTPResponse struct {
	Message string
	Status  int
}

// GrpcResponse defines a default grpc response
type GrpcResponse struct {
	Message string
	Code    codes.Code
}

// MessageAction defines a default message response
type MessageAction int8

// LogLevel defines the log level
type LogLevel int8

// detailedError a error that wraps a random error
type detailedError struct {
	err                  error
	publicErr            error
	errLogLevel          *LogLevel
	defaultGrpcResponse  *GrpcResponse
	defaultHTTPResponse  *HTTPResponse
	defaultMessageAction *MessageAction
}

func (de detailedError) Error() string {
	return de.err.Error()
}

func (de *detailedError) Unwrap() error {
	return Unwrap(de.err)
}
