package roxy

type NewErrorOption func(error) error

func WithMessageAction(defaultMessageAction MessageAction) NewErrorOption {
	return func(err error) error {
		return SetDefaultMessageAction(err, defaultMessageAction)
	}
}

func WithHTTPResponse(defaultHTTPResponse HTTPResponse) NewErrorOption {
	return func(err error) error {
		return SetDefaultHTTPResponse(err, defaultHTTPResponse)
	}
}

func WithGrpcResponse(defaultGrpcResponse GrpcResponse) NewErrorOption {
	return func(err error) error {
		return SetDefaultGrpcResponse(err, defaultGrpcResponse)
	}
}

func WithLogLevel(defaultLogLevel LogLevel) NewErrorOption {
	return func(err error) error {
		return SetErrorLogLevel(err, defaultLogLevel)
	}
}

func WithPublicError(defaultPublicError error) NewErrorOption {
	return func(err error) error {
		return SetPublicError(err, defaultPublicError)
	}
}
