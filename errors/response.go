package ce

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sorrawichYooboon/golang-error-handling-and-unit-test/constants"
)

type ErrorResponseJsonModel struct {
	Code             string `json:"code"`
	MessageTechnical string `json:"messageTechnical"`
}

func ErrorResponseJson(ctx *gin.Context, err error) (errorCode string) {
	var response ErrorResponseJsonModel

	switch err.(type) {
	case *RedisError:
		errorCode = constants.ERROR_REDIS
		response = generateErrorResponse(err, ctx, errorCode)
		ctx.JSON(http.StatusInternalServerError, response)
	case *InvalidRequestError:
		errorCode = constants.ERROR_BAD_REQUEST
		response = generateErrorResponse(err, ctx, errorCode)
		ctx.JSON(http.StatusBadRequest, response)
	case *InvalidFormatError:
		errorCode = constants.ERROR_INVALID_FORMAT
		response = generateErrorResponse(err, ctx, errorCode)
		ctx.JSON(http.StatusBadRequest, response)
	default:
		errorCode = constants.ERROR_INTERNAL
		response = generateErrorResponse(err, ctx, errorCode)
		ctx.JSON(http.StatusInternalServerError, response)
	}

	return errorCode
}

func generateErrorResponse(err error, ctx *gin.Context, errorCode string) ErrorResponseJsonModel {
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
