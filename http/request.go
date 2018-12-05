/*
Copyright 2018 liipx(lipengxiang)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package http

import (
	"io"
	"net/http"
)

// Request is a http request client
type Request struct {
	*http.Request
	*Client

	// Err is the error when structure Request
	Err error
}

// BaseHeader xlib/http client header
var BaseHeader = http.Header{
	"User-Agent": []string{"xlib/http " + version},
	"Connection": []string{"keep-alive"},
	"Accept":     []string{"application/json"},
}

// NewRequest returns a Request extends http.Request using global http.client 'Cli'
func NewRequest(method, url string, body io.Reader, headers ...http.Header) *Request {
	return newRequest(method, url, body, Cli, headers...)
}

func newRequest(method, url string, body io.Reader, client *Client, headers ...http.Header) *Request {
	// new http request
	req, err := http.NewRequest(method, url, body)

	// add base header
	req.Header = BaseHeader

	// merge & add custom headers
	headerItem := make(map[string]interface{})
	for _, header := range headers {
		for k, v := range header {
			for idx, item := range v {

				// overwrite base header
				if _, ok := BaseHeader[k]; ok && idx == 0 {
					req.Header.Set(k, item)
				} else {
					// add item
					if _, ok := headerItem[k+item]; !ok {
						req.Header.Add(k, item)
					}
				}

				// register item
				headerItem[k+item] = true
			}
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
	return req.Client.Do(req)
}

func (req *Request) debug() {
	if !Debug {
		return
	}

	log.Debug(req.Method + ": " + req.URL.String())
	log.Debug("Request header:", headerPretty(req.Header))
}
