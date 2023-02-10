package roxy

import (
	"google.golang.org/grpc/codes"
)

// HTTPResponse defines a default http response
type HTTPResponse struct {
	Message string
	Status  int
}

// MessageAction defines a default message response
type MessageAction int

// DefaultMessageResponse values
const (
	DropMessageAction MessageAction = iota
	SuccessMessageAction
	RequeueMessageAction
	DeadLetterMessageAction
)

// detailedError a error that wraps a random error
type detailedError struct {
	err                  error
	defaultGrpcResponse  *codes.Code
	defaultHTTPResponse  *HTTPResponse
	defaultMessageAction *MessageAction
}

func (de detailedError) Error() string {
	return de.err.Error()
}

func (de *detailedError) Unwrap() error {
	return Unwrap(de.err)
}
