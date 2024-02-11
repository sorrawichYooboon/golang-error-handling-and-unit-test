# Error handling and Unit testing in Golang

## Setup

```bash
go mod tidy
```

> then run go main to start the server

```bash
go run main.go
```

### Install Gin

> Fiber is an Express inspired web framework written in Go

```bash
go get -u github.com/gin-gonic/gin
```

### Install Testify

> Testify is a toolkit with common assertions and mocks that plays nicely with the standard library.

> not use `"github.com/go-playground/assert/v2"` because it not compare deep equal

```bash
go get -u github.com/stretchr/testify/mock
```

### Install Mockery

> Mockery is a mock code autogenerator for Golang.

```bash
go get github.com/vektra/mockery/v2/.../
```

### Go coverage highlighting (File should place beside of testing file)

> Go coverage highlighting is a plugin for highlighting lines in the Go source code that are not covered by tests.

place in json settings file

```json
"go.coverageOptions": "showBothCoveredAndUncoveredCode",
  "go.coverageDecorator": {
    "type": "highlight",
    "coveredHighlightColor": "rgba(59, 255, 125, 0.03)",
    "uncoveredHighlightColor": "rgba(255, 100, 81, 0.03)",
    "coveredBorderColor": "rgba(59, 255, 125, 0.03)",
    "uncoveredBorderColor": "rgba(255, 100, 81, 0.03)",
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
  },
  "editor.unicodeHighlight.nonBasicASCII": false,
```

## Mockery

> **--dir**: The directory in which to search for interfaces. (required)<br> **--name**: The name of the interface for which to generate a mock. (required)<br> **--output**: The directory to which the mock file will be written. (required)<br> **--outpkg**: The name of the package in which the mock should be generated. (required)<br> **--filename**: The name of the file to which the mock will be written. (optional)

> <ins>NOTE</ins> (backslash) at the end of each line indicates that the command continues on the next line.

<ins>mock with a specified, clear path</ins>

- **when to use**: `When you want to mock a function that is not in the same package.`

> #### services
>
> ```bash
> mockery \
> --dir=services/cacheService \
> --name=ICacheService \
> --output=services/mock \
> --outpkg=mockService \
> --filename=ICacheService.go
> ```

> #### repositories
>
> ```bash
> mockery \
> --dir=repositories/privilegeRepository \
> --name=IPrivilegeRepository \
> --output=repositories/mock \
> --outpkg=mockRepo \
> --filename=IPrivilegeRepository.go
> ```

<ins>mock all interfaces in a package</ins>

> #### services
>
> ```bash
> mockery \
> --dir=services \
> --output=services/mock \
> --outpkg=mockService \
> --all
> ```

> #### repositories
>
> ```bash
> mockery \
> --dir=repositories \
> --output=repositories/mock \
> --outpkg=mockRepo \
> --all
> ```

## How to run test

### Unit Test

> Unit tests are tests that test a single unit of code.

```bash
go test ./...
```

> to see test coverage

```bash
go test -cover ./...
```

---

# Error Handling in Go Application

This repository demonstrates a common approach to error handling in a Go application using the `github.com/pkg/errors` package. The application follows best practices for handling and logging errors, providing clear error messages and stack traces for debugging purposes.

## Error Wrapping

The `ce` package provides functions for wrapping errors with additional context and handling specific error types. This allows for more informative error messages and consistent error handling across the application.

### Error Types

The following error types are defined in the `ce` package:

- `RedisError`: Represents errors related to Redis operations.
- `InvalidRequestError`: Indicates an invalid request error, such as missing or malformed request parameters.
- `InvalidFormatError`: Denotes errors caused by invalid data formats.
- `InternalError`: Represents generic internal server errors.

## Error Responses

The `ce` package also includes functions for generating and handling error responses. This ensures that error responses returned by the API are consistent and informative.

### Generating Error Responses

The `ErrorResponseJson` function generates JSON error responses based on the error type. It maps each error type to an appropriate HTTP status code and error message.

### Handling Error Responses

The `HandleErrorResponse` function is used to handle error responses in API handlers. It logs the error message and stack trace for debugging purposes.

## Logging

The `ce` package includes functions for logging error messages and stack traces to aid in debugging.

### Logging Error Stack Traces

The `LogErrorStackTrace` function logs the stack trace of an error, if available.

### Logging Cause Error Messages

The `LogCauseErrorMessages` function logs the error message and its causes.

### Logging Error Stack Traces and Endpoints

The `GetErrorStackTraceAndEndpoint` function logs the stack trace of an error along with the endpoint and HTTP method where the error occurred.

## Usage

To use the error handling functionality in your application:

1.  Import the `ce` package.
2.  Wrap errors using the appropriate error type functions provided by the `ce` package.
3.  Handle error responses in API handlers using the `HandleErrorResponse` function.
