package http

import "net/http"

type Response struct {
	*http.Response

	// Err is the error doing the request.
	Err error
}

func (resp *Response) JSONResult() {}
func (resp *Response) XMLResult() {}