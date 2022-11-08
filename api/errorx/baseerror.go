package errorx

import "net/http"

const defaultCode = 42
const defaultHttpStatus = http.StatusOK

type CodeError struct {
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	httpStatus uint
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ApiError interface {
	error
	Data() *CodeErrorResponse
	StatusHttp() uint
}

func NewCodeError(code int, msg string, status uint) error {
	return &CodeError{Code: code, Msg: msg, httpStatus: status}
}

func NewDefaultError(msg string) error {
	return &CodeError{Code: defaultCode, Msg: msg, httpStatus: defaultHttpStatus}
}

func (c *CodeError) Error() string {
	return c.Msg
}

func (c *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: c.Code,
		Msg:  c.Msg,
	}
}

func (c *CodeError) StatusHttp() uint {
	return c.httpStatus
}
