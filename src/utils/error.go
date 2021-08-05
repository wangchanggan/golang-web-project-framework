package utils

import "fmt"

const (
	ErrorCode            = "errorCode"
	ErrorDesc            = "errorDesc"
	Content              = "content"
)

const (
	ErrorCodeSuccess = "Success"
	ErrorCodeParameterError = "ParameterError"
	ErrorCodeInvalidParameter = "InvalidParameter"
)

const (
	ErrorDescSuccess = "Success"
	ErrorDescInvalidParameter = "parameter %s is invalid"
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

func (error *Error) Error() string {
	return error.ErrorDesc
}

