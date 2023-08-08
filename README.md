# roxy
Roxy is a project named after a character from "Mushoku Tensei," designed as a wrapper for another project named Eris, which coincidentally shares its name with another character from the same series.

## Description
Roxy aims to simplify error handling in your codebase by providing a mechanism to implement default behaviors for different types of errors. This approach allows you to define multiple responses for the same error while reducing the complexity of your error-handling code.

## Features
Roxy introduces the following functions to enhance your error handling workflow:

`SetDefaultGrpcResponse`, `GetDefaultGrpcResponse`: Set and retrieve gRPC responses for specific errors. <br/>
`SetDefaultHTTPResponse`, `GetDefaultHTTPResponse`: Set and retrieve HTTP responses for specific errors. <br/>
`SetDefaultMessageAction`, `GetDefaultMessageAction`: Define and retrieve message actions to be executed for certain errors. <br/>
`SetErrorLogLevel`, `GetErrorLogLevel`: Configure and obtain the desired logging level for specific errors (e.g., info, warn, error). <br/>

## How to use
Suppose you are building a movie app and want to handle different error scenarios gracefully. Here's how you can leverage Roxy to achieve this:

``1. Default Case: you want to return a "404 Movie not found" error when querying for a movie, and none is found. Log it as an info level.`` <br/>
``2. You want to return a "500" error when performing an operation on a movie, and none is found during querying.``

---

<br>

Define a custom error type, "MovieNotFound," to be used as the default behavior:
``` golang
// errors/movie.go

var MovieNotFound = roxy.New("Movie not found")
MovieNotFound = roxy.SetDefaultHTTPResponse(MovieNotFound, roxy.HTTPResponse{
    Message: "Movie not found",
    Status:  http.StatusNotFound,
  })
MovieNotFound = roxy.SetErrorLogLevel(MovieNotFound, roxy.InfoLevel)
```

<br>

Overwrite your default response for the necessary cases:
``` golang
// feature/someCode.go

// ... previous code
_, err = query.GetMovie(ctx, id)
if roxy.Is(err, errors.MovieNotFound) {
  err = roxy.SetDefaultHTTPResponse(MovieNotFound, UnhandledHTTPResponse)
  err = roxy.SetErrorLogLevel(err, roxy.ErrorLevel)
}
if err != nil{
  return err
}
// ... following code
```

<br>

Implement a function that logs based on the logger used by your system:
``` golang
// errors/logError.go

// LogError logs error based on roxy's log level
func LogError(ctx context.Context, err error) {
  logger := log.Ctx(ctx)
  logLevel := roxy.GetErrorLogLevel(err)

  if err == nil {
    return
  }

  switch logLevel {
  case roxy.Disabled:
    return
  case roxy.TraceLevel:
    logger.Trace().Err(err).Msg(err.Error())
  // ... Other log level cases ...
  default:
    logger.Error().Err(err).Msg(err.Error())
  }
}
```

<br>

Implement an error handler for the HTTP interface of your system:
``` golang
// controller/errorHandler.go

func handleError(ctx context.Context, err error, rw http.ResponseWriter) {
  httpError := roxy.GetDefaultHTTPResponse(err)
  errorMessage := httpError.Message
  errorCode := httpError.Status

  errors.LogError(ctx, err)
  http.Error(rw, errorMessage, errorCode)
}
```

## Conclusion
Roxy provides an elegant solution to streamline error handling in your codebase by allowing you to define and manage default behaviors for various error types. By following the example provided above, you can create a more resilient and maintainable application that gracefully handles different error scenarios.
