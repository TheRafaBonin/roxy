package roxy

import (
	"reflect"
)

// LogLevel values
const (
	// DebugLevel defines debug log level.
	DebugLevel LogLevel = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel LogLevel = -1
)

// SetErrorLogLevel ...
func SetErrorLogLevel(err error, level LogLevel) error {
	if err == nil {
		return nil
	}

	errType := reflect.TypeOf(err)
	if errType != reflect.TypeOf(detailedError{}) {
		err = new(err)
	}

	eDetailedError := err.(*detailedError)
	(*eDetailedError).errLogLevel = &level
	return eDetailedError
}

// GetErrorLogLevel ...
func GetErrorLogLevel(err error) LogLevel {
	if err == nil {
		return InfoLevel
	}

	currentLogLevel := ErrorLevel

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
		if valid && detailedErr.errLogLevel != nil {
			return *detailedErr.errLogLevel
		}
	}

	return currentLogLevel
}
