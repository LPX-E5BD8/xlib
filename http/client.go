package http

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
		resp.Err = ErrorEmptyReq
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

func (c *Client) Get(url string, headers ...http.Header) *Response {
	return newRequest("GET", url, nil, c, headers...).Do()
}

func (c *Client) Post(url string, params io.Reader, headers ...http.Header) *Response {
	return newRequest("POST", url, params, c, headers...).Do()
}

func (c *Client) Put(url string, params io.Reader, headers ...http.Header) *Response {
	return newRequest("PUT", url, params, c, headers...).Do()
}

func (c *Client) Delete(url string, params io.Reader, headers ...http.Header) *Response {
	return newRequest("DELETE", url, params, c, headers...).Do()
}
