package http

import (
	"encoding/base64"
	"io"
	"net/http"
)

type Request struct {
	*http.Request
	*Client

	// Err is the error when structure Request
	Err error
}

//
var BaseHeader = map[string]string{
	"User-Agent": "xlib/http v1.0",
	"Connection": "keep-alive",
	"Accept":     "application/json",
}

// NewRequest returns a Request extends http.Request using global http.client 'Session'
func NewRequest(method, url string, body io.Reader, headers ...map[string]string) *Request {
	return newRequest(method, url, body, Session, headers...)
}

func newRequest(method, url string, body io.Reader, client *Client, headers ...map[string]string) *Request {
	// new http request
	req, err := http.NewRequest(method, url, body)

	// add base header
	headers = append([]map[string]string{BaseHeader}, headers...)

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

// SetBasicAuth for request.
func (req *Request) SetBasicAuth(userName, password string) {
	req.Request.SetBasicAuth(userName, password)
}

// Do will do the response.
func (req *Request) Do() *Response {
	return Session.Do(req)
}

func BasicAuthHeader(userName, password string) map[string]string {
	auth := userName + ":" + password
	return map[string]string{
		"Authorization": base64.StdEncoding.EncodeToString([]byte(auth)),
	}
}
