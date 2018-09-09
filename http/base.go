package http

import (
	"io"
	"net/http"

	"github.com/op/go-logging"
)

var Debug = false

var log = logging.MustGetLogger("xlib")

func Get(url string, headers ...http.Header) *Response {
	return DoRequest("GET", url, nil, headers...)
}

func Post(url string, params io.Reader, headers ...http.Header) *Response {
	return DoRequest("POST", url, params, headers...)
}

// DoRequest returns a Response & an error if something wrong.
func DoRequest(method, url string, params io.Reader, headers ...http.Header) *Response {
	return NewRequest(method, url, params, headers...).Do()
}
