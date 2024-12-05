package reponse

import "bi-activity/reponse/errors"

type Response struct {
	Status int         `json:"status"`
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

func Fail(err string) (int, *Response) {
	return errors.Err[err], &Response{
		Status: errors.ErrSelf[err],
		Error:  err,
	}
}

func Success() (int, *Response) {
	return 200, &Response{
		Status: 200,
	}
}
