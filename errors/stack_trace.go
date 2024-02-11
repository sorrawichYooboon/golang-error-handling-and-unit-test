package ce

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func LogErrorStackTrace(err error) {
	if sterr, ok := errors.Cause(err).(stackTracer); ok {
		log.Printf("\n[Error Message]: %v \n[Stack trace]:\n", err)
		for n, f := range sterr.StackTrace() {
			fmt.Printf("%d: %s %n:%d\n", n, f, f, f)
		}
	}
}

func LogCauseErrorMessages(err error) {
	fmt.Println(GetCauseErrorMessages(err))
}

func GetErrorStackTraceAndEndpoint(err error, method string, endpoint string, errCode string) string {
	var errMsg string
	if sterr, ok := errors.Cause(err).(stackTracer); ok {
		errMsg = fmt.Sprintf("[Error Code]: %s \n[Error Message]: %v \n[Stack trace]:\n", errCode, err)
		for n, f := range sterr.StackTrace() {
			errMsg += fmt.Sprintf("%d: %s %n:%d\n", n, f, f, f)
		}
	}

	return fmt.Sprintf("\n[%s-Endpoint]: %s \n%s", method, endpoint, errMsg)
}

func GetErrorStackTrace(err error) string {
	var errMsg string
	if sterr, ok := errors.Cause(err).(stackTracer); ok {
		errMsg = fmt.Sprintf("\n Error Message: %v \nStack trace:\n", err)
		for n, f := range sterr.StackTrace() {
			errMsg += fmt.Sprintf("%d: %s %n:%d\n", n, f, f, f)
		}
	}

	return errMsg
}

func GetCauseErrorMessages(err error) string {
	var errMsg string
	for {
		errMsg = err.Error()
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
	}

	return errMsg
}
