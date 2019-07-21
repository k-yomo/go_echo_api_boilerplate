package error_code

import "net/http"

// ErrorCode is a readable custom_error type code
type ErrorCode string

// String returns stringified error code
func (ec ErrorCode) String() string {
	return string(ec)
}

const (
	BadRequest       ErrorCode = "BadRequest"
	InvalidHeader    ErrorCode = "InvalidHeader"
	Unauthenticated  ErrorCode = "Unauthenticated"
	Expired          ErrorCode = "Expired"
	InvalidParams    ErrorCode = "InvalidParams"
	NotFound         ErrorCode = "NotFound"
	MethodNotAllowed ErrorCode = "MethodNotAllowed"
	AlreadyTaken     ErrorCode = "AlreadyTaken"
	InternalError    ErrorCode = "InternalError"
	UnhandledError   ErrorCode = "UnhandledError"
)

var errorCodeStatusMap = map[ErrorCode]int{
	BadRequest:       http.StatusBadRequest,
	InvalidHeader:    http.StatusBadRequest,
	Unauthenticated:  http.StatusUnauthorized,
	Expired:          http.StatusUnauthorized,
	InvalidParams:    http.StatusUnprocessableEntity,
	NotFound:         http.StatusNotFound,
	MethodNotAllowed: http.StatusMethodNotAllowed,
	AlreadyTaken:     http.StatusConflict,
	InternalError:    http.StatusInternalServerError,
}

// GetHTTPStatus returns http status code which matches given custom_error code
func GetHTTPStatusByErrorCode(code ErrorCode) int {
	return errorCodeStatusMap[code]
}
