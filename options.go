package roxy

type NewErrorOptions func(error) error

func WithMessageAction(defaultMessageAction MessageAction) NewErrorOptions {
	return func(err error) error {
		return SetDefaultMessageAction(err, defaultMessageAction)
	}
}

func WithHTTPResponse(defaultHTTPResponse HTTPResponse) NewErrorOptions {
	return func(err error) error {
		return SetDefaultHTTPResponse(err, defaultHTTPResponse)
	}
}

func WithGrpcResponse(defaultGrpcResponse GrpcResponse) NewErrorOptions {
	return func(err error) error {
		return SetDefaultGrpcResponse(err, defaultGrpcResponse)
	}
}

func WithLogLevel(defaultLogLevel LogLevel) NewErrorOptions {
	return func(err error) error {
		return SetErrorLogLevel(err, defaultLogLevel)
	}
}

func WithPublicError(defaultPublicError error) NewErrorOptions {
	return func(err error) error {
		return SetPublicError(err, defaultPublicError)
	}
}
