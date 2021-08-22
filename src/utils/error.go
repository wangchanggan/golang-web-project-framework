package utils

import "fmt"

const (
	ErrorCode = "errorCode"
	ErrorDesc = "errorDesc"
	Content   = "content"
)

const (
	ErrorCodeSuccess          = "Success"
	ErrorCodeParameterError   = "ParameterError"
	ErrorCodeInvalidParameter = "InvalidParameter"
	ErrorCodeMissingParameter = "MissingParameter"
	ErrorCodeInnerError       = "InnerError"
	ErrorCodeNotFound         = "NotFound"
)

const (
	ErrorDescSuccess          = "success"
	ErrorDescInvalidParameter = "parameter %s is invalid"
	ErrorDescMissingParameter = "parameter %s is missing"
	ErrorDescNotFound         = "not found"
)

type Error struct {
	ErrorCode string `json:"errorCode"`
	ErrorDesc string `json:"errorDesc"`
}

func NewError(code, desc string) *Error {
	return &Error{ErrorCode: code, ErrorDesc: desc}
}

func NewErrorWithInvalidParameter(param string) *Error {
	return &Error{ErrorCode: ErrorCodeInvalidParameter, ErrorDesc: fmt.Sprintf(ErrorDescInvalidParameter, param)}
}

func NewErrorWithMissingParameter(param string) *Error {
	return &Error{ErrorCode: ErrorCodeMissingParameter, ErrorDesc: fmt.Sprintf(ErrorDescMissingParameter, param)}
}

func NewErrorWithInnerError(errStr string) *Error {
	return &Error{ErrorCode: ErrorCodeInnerError, ErrorDesc: errStr}
}

func NewErrorWithNotFound() *Error {
	return &Error{ErrorCode: ErrorCodeNotFound, ErrorDesc: ErrorDescNotFound}
}
