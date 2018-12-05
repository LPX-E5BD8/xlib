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

	"github.com/op/go-logging"
)

var (
	// Debug flag
	Debug = false
	log   = logging.MustGetLogger("xlib")
)

// Get will request the url with 'GET' method and return a Response
func Get(url string, headers ...http.Header) *Response {
	return DoRequest("GET", url, nil, headers...)
}

// Post will request the url with 'POST' method and return a Response
func Post(url string, params io.Reader, headers ...http.Header) *Response {
	return DoRequest("POST", url, params, headers...)
}

// DoRequest returns a Response & an error if something wrong.
func DoRequest(method, url string, params io.Reader, headers ...http.Header) *Response {
	return NewRequest(method, url, params, headers...).Do()
}
