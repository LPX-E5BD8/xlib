package http

import (
	"net/http"
	"io"
)

type Request struct {
	*http.Request
	*Client

	// Err is the error when structure Request
	Err error
}

// NewRequest returns a Request extends http.Request using global http.client 'Session'
func NewRequest(method, url string, body io.Reader, headers ...map[string]string) *Request {
	return newRequest(method, url, body, Session, headers...)
}

func newRequest(method, url string, body io.Reader, client *Client, headers ...map[string]string) *Request {
	// new http request
	req, err := http.NewRequest(method, url, body)

	// set headers
	for _, header := range headers {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	return &Request{
		Request: req,
		Client:  client,
		Err:     err,
	}
}

// Do will do the response.
func (resp *Request) Do() *Response {
	return Session.Do(resp)
}
