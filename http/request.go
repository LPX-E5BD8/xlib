package http

import (
	"net/http"
	"io"
)

type Request struct {
	*http.Request
	*Client
}

// NewRequest return a Request extends http.Request
func NewRequest(method, url string, body io.Reader, headers ...map[string]string) (*Request, error) {
	// new http request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// set headers
	for _, header := range headers {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	return &Request{
		Request: req,
		Client:  Session,
	}, nil

}
