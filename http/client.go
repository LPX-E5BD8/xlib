package http
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

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// Cli is a global http.Client
var Cli = &Client{&http.Client{}}

func init() {
	// This function always return nil error.
	Cli.Jar, _ = cookiejar.New(nil)
}

// ClientOptions describe the options of a Client
type ClientOptions struct {
	Timeout       time.Duration
	RoundTripper  http.RoundTripper
	Jar           http.CookieJar
	CheckRedirect func(req *http.Request, via []*http.Request) error
}

// Client extends http.Client
type Client struct {
	*http.Client
}

// NewClient return a Client pointer
func NewClient(options *ClientOptions) (*Client) {
	client := &Client{
		Client: &http.Client{
			Timeout:       options.Timeout,
			Transport:     options.RoundTripper,
			Jar:           options.Jar,
			CheckRedirect: options.CheckRedirect,
		},
	}
	return client
}

// NewRequest returns a Request extends http.Request with Client
func (c *Client) NewRequest(method, url string, body io.Reader, headers ...http.Header) *Request {
	return newRequest(method, url, body, c, headers...)
}

// SetCookieJar set cookies for client
func (c *Client) SetCookieJar(jar cookiejar.Jar) (*Client) {
	c.Jar = &jar
	return c
}

// SetCookie set the cookie into response
func (c *Client) SetCookie(key, value string) (error) {
	return nil
}

// GetCookie get the cookie from client
func (c *Client) GetCookie(key, value string) (string, error) {
	return "", nil
}

// Do will do the request with this client
func (c *Client) Do(req *Request) *Response {
	resp := &Response{}

	if req == nil {
		resp.Err = errorEmptyReq
		return resp
	}

	if req.Err != nil {
		resp.Err = req.Err
		return resp
	}

	var err error
	resp.Response, err = c.Client.Do(req.Request)
	if err != nil {
		resp.Err = err
	}

	// if debug
	req.debug()
	resp.debug()

	return resp
}

// Get http request GET
func (c *Client) Get(url string, headers ...http.Header) *Response {
	return newRequest("GET", url, nil, c, headers...).Do()
}

// Post http request Post
func (c *Client) Post(url string, params io.Reader, headers ...http.Header) *Response {
	return newRequest("POST", url, params, c, headers...).Do()
}

// Put http request Put
func (c *Client) Put(url string, params io.Reader, headers ...http.Header) *Response {
	return newRequest("PUT", url, params, c, headers...).Do()
}

// Delete http request Delete
func (c *Client) Delete(url string, params io.Reader, headers ...http.Header) *Response {
	return newRequest("DELETE", url, params, c, headers...).Do()
}
