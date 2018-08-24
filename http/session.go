package http

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// Session is a global http.Client
var Session = &Client{&http.Client{}}

func init() {
	Session.Jar, _ = cookiejar.New(nil)
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
func NewClient(options ClientOptions) (*Client) {
	session := &Client{
		Client: &http.Client{
			Timeout:       options.Timeout,
			Transport:     options.RoundTripper,
			Jar:           options.Jar,
			CheckRedirect: options.CheckRedirect,
		},
	}
	return session
}

// NewRequest returns a Request extends http.Request with Client
func (c *Client) NewRequest(method, url string, body io.Reader, headers ...map[string]string) *Request {
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
	if req.Err == nil {
		return nil
	}

	resp, err := c.Client.Do(req.Request)
	return &Response{
		Response: resp,
		Err:      err,
	}
}

func (c *Client) Get()    {}
func (c *Client) Post()   {}
func (c *Client) Put()    {}
func (c *Client) Delete() {}
