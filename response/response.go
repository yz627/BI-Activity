package response

import (
	"bi-activity/response/errors"
	"fmt"
	"net/http"
)

type Response struct {
	Status int         `json:"label"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func (r *Response) WithMsg(msg string) *Response {
	r.Msg = msg
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) WithStatus(status int) *Response {
	r.Status = status
	return r
}

func (r *Response) WithError(err string) *Response {
	r.Error = err
	return r
}

func Fail(err errors.SelfError) (int, *Response) {
	return errors.ErrStatus[err], &Response{
		Status: errors.SelfErrStatus[err],
		Error:  err.Err,
	}
}

func Failf(err errors.SelfError, format string, args ...interface{}) (int, *Response) {
	return errors.ErrStatus[err], &Response{
		Status: errors.SelfErrStatus[err],
		Error:  err.Error(),
		Msg:    fmt.Sprintf(format, args...),
	}
}

func Success(data ...interface{}) (int, *Response) {
	res := &Response{
		Status: http.StatusOK,
	}

	if len(data) > 0 {
		res.WithData(data[0])
	}
	return http.StatusOK, res
}
