package ce

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/constants"
)

type ErrorResponseJsonModel struct {
	Code             string `json:"code"`
	MessageTechnical string `json:"messageTechnical"`
}

func ErrorResponseJson(ctx *gin.Context, err error) string {
	var errorCode string
	switch err.(type) {
	case *RedisError:
		errorCode = constants.ERROR_REDIS
	case *InvalidRequestError:
		errorCode = constants.ERROR_BAD_REQUEST
	case *InvalidFormatError:
		errorCode = constants.ERROR_INVALID_FORMAT
	default:
		errorCode = constants.ERROR_INTERNAL
	}

	response := generateErrorResponse(err, errorCode)
	ctx.JSON(ErrorCodeToHTTPStatus[errorCode], response)
	return errorCode
}

func generateErrorResponse(err error, errorCode string) ErrorResponseJsonModel {
	return ErrorResponseJsonModel{
		Code:             errorCode,
		MessageTechnical: err.Error(),
	}
}

func HandleErrorResponse(c *gin.Context, err error) {
	errorCode := ErrorResponseJson(c, err)
	errLog := GetErrorStackTraceAndEndpoint(err, c.Request.Method, c.Request.URL.Path, errorCode)

	fmt.Println(errLog)
}
