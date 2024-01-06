package roxy

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	t.Run("testNew", testNew)
	t.Run("testErrorf", testErrorf)
}

func testNew(t *testing.T) {
	tests := []struct {
		name string
		opts []NewErrorOptions
	}{
		{
			name: "WithMessageAction",
			opts: []NewErrorOptions{
				WithMessageAction(DropMessageAction),
			},
		},
		{
			name: "WithHTTPResponse",
			opts: []NewErrorOptions{
				WithHTTPResponse(HTTPResponse{
					Message: "test",
					Status:  200,
				}),
			},
		},
		{
			name: "WithGrpcResponse",
			opts: []NewErrorOptions{
				WithGrpcResponse(GrpcResponse{
					Message: "test",
					Code:    0,
				}),
			},
		},
		{
			name: "WithLogLevel",
			opts: []NewErrorOptions{
				WithLogLevel(DebugLevel),
			},
		},
		{
			name: "WithPublicError",
			opts: []NewErrorOptions{
				WithPublicError(errors.New("test")),
			},
		},
		{
			name: "WithMessageAction, WithHTTPResponse, WithGrpcResponse, WithLogLevel, WithPublicError",
			opts: []NewErrorOptions{
				WithMessageAction(DropMessageAction),
				WithHTTPResponse(HTTPResponse{
					Message: "test",
					Status:  200,
				}),
				WithGrpcResponse(GrpcResponse{
					Message: "test",
					Code:    0,
				}),
				WithLogLevel(DebugLevel),
				WithPublicError(errors.New("test")),
			},
		},
		{
			name: "BaseError",
			opts: []NewErrorOptions{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseString := tt.name
			err := New(baseString, tt.opts...)

			stringError := err.Error()
			if stringError != baseString {
				t.Errorf("%s is not equal to %s", stringError, baseString)
			}

			_, ok := err.(*detailedError)
			if !ok {
				t.Error("Could not cast to DetailedError")
			}
		})
	}

}

func testErrorf(t *testing.T) {
	baseString := "BaseError"
	err := Errorf("%sError", "Base")

	stringError := err.Error()
	if stringError != baseString {
		t.Errorf("%s is not equal to %s", stringError, baseString)
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(&detailedError{}) {
		t.Error("Could not cast to DetailedError")
	}
}
