# Error handling and Unit testing in Golang

Test is in producer folder

## Setup

### Install Gin

> Fiber is an Express inspired web framework written in Go

```bash
go get -u github.com/gin-gonic/gin
```

### Install Testify

> Testify is a toolkit with common assertions and mocks that plays nicely with the standard library.

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
