package constants

const (
	ERROR_INVALID_FORMAT         string = "INVALID_FORMAT"         // 400
	ERROR_BAD_REQUEST            string = "BAD_REQUEST"            // 400
	ERROR_UNAUTHORIZED           string = "UNAUTHORIZED"           // 401
	ERROR_FORBIDDEN              string = "FORBIDDEN"              // 403
	ERROR_NOT_FOUND              string = "NOT_FOUND"              // 404
	ERROR_CONFLICT               string = "CONFLICT"               // 409
	ERROR_INTERNAL               string = "INTERNAL"               // 500
	ERROR_REDIS_CONNECTION_ERROR string = "REDIS_CONNECTION_ERROR" // 500
	ERROR_REDIS                  string = "REDIS_ERROR"            // 500
	ERROR_NOT_IMPLEMENTED        string = "NOT_IMPLEMENTED"        // 501
)
