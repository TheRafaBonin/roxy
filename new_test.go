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
		opts []NewErrorOption
	}{
		{
			name: "WithMessageAction",
			opts: []NewErrorOption{
				WithMessageAction(DropMessageAction),
			},
		},
		{
			name: "WithHTTPResponse",
			opts: []NewErrorOption{
				WithHTTPResponse(HTTPResponse{
					Message: "test",
					Status:  200,
				}),
			},
		},
		{
			name: "WithGrpcResponse",
			opts: []NewErrorOption{
				WithGrpcResponse(GrpcResponse{
					Message: "test",
					Code:    0,
				}),
			},
		},
		{
			name: "WithLogLevel",
			opts: []NewErrorOption{
				WithLogLevel(DebugLevel),
			},
		},
		{
			name: "WithPublicError",
			opts: []NewErrorOption{
				WithPublicError(errors.New("test")),
			},
		},
		{
			name: "WithMessageAction, WithHTTPResponse, WithGrpcResponse, WithLogLevel, WithPublicError",
			opts: []NewErrorOption{
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
			opts: []NewErrorOption{},
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
