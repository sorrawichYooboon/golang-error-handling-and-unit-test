package ce

import (
	"github.com/pkg/errors"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/constants"
)

var ErrorCodeToHTTPStatus = map[string]int{
	constants.ERROR_INVALID_FORMAT:         400,
	constants.ERROR_BAD_REQUEST:            400,
	constants.ERROR_UNAUTHORIZED:           401,
	constants.ERROR_FORBIDDEN:              403,
	constants.ERROR_NOT_FOUND:              404,
	constants.ERROR_CONFLICT:               409,
	constants.ERROR_INTERNAL:               500,
	constants.ERROR_REDIS_CONNECTION_ERROR: 500,
	constants.ERROR_REDIS:                  500,
	constants.ERROR_NOT_IMPLEMENTED:        501,
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type RedisError struct {
	error
}

func ErrorRedis(err error) error {
	return &RedisError{err}
}

func (e *RedisError) Error() string {
	return e.error.Error()
}

func (e *RedisError) StackTrace() errors.StackTrace {
	if sterr, ok := errors.Cause(e.error).(stackTracer); ok {
		return sterr.StackTrace()
	}
	return nil
}

type InvalidFormatError struct {
	error
}

func ErrorInvalidFormat(err error) error {
	return &InvalidFormatError{err}
}

func (e *InvalidFormatError) Error() string {
	return e.error.Error()
}

func (e *InvalidFormatError) StackTrace() errors.StackTrace {
	if sterr, ok := errors.Cause(e.error).(stackTracer); ok {
		return sterr.StackTrace()
	}
	return nil
}

type InternalError struct {
	error
}

func ErrorInternal(err error) error {
	return &InternalError{err}
}

func (e *InternalError) Error() string {
	return e.error.Error()
}

func (e *InternalError) StackTrace() errors.StackTrace {
	if sterr, ok := errors.Cause(e.error).(stackTracer); ok {
		return sterr.StackTrace()
	}
	return nil
}

type InvalidRequestError struct {
	error
}

func ErrorInvalidRequest(err error) error {
	return &InvalidRequestError{err}
}

func (e *InvalidRequestError) Error() string {
	return e.error.Error()
}

func (e *InvalidRequestError) StackTrace() errors.StackTrace {
	if sterr, ok := errors.Cause(e.error).(stackTracer); ok {
		return sterr.StackTrace()
	}
	return nil
}

func ErrorWrapper(err error, errMsg string) error {
	switch err.(type) {
	case *RedisError:
		return ErrorRedis(errors.Wrap(err, errMsg))
	case *InvalidRequestError:
		return ErrorInvalidRequest(errors.Wrap(err, errMsg))
	case *InvalidFormatError:
		return ErrorInvalidFormat(errors.Wrap(err, errMsg))
	default:
		return ErrorInternal(errors.Wrap(err, errMsg))
	}
}
