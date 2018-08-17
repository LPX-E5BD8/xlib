package http

import (
	"net/http"
	"time"
	"net/http/cookiejar"
)

// Session is a global http.Client
var Session = &Client{&http.Client{}}

func init() {
	Session.Jar, _ = cookiejar.New(nil)
}

// SessionOptions describe the options of a Client
type SessionOptions struct {
	Timeout       time.Duration
	RoundTripper  *http.RoundTripper
	Jar           http.CookieJar
	CheckRedirect func(req *http.Request, via []*http.Request) error
}

// Client extends http.Client
type Client struct {
	*http.Client
}

// NewSession return a Client pointer
func NewSession(options SessionOptions) (*Client) {
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

func (c *Client) SetCookieJar(jar cookiejar.Jar) (*Client) {
	c.Jar = &jar
	return c
}

// SetCookie set the cookie into response
func (c *Client) SetCookie(key, value string) (error) {
	return nil
}

// SetCookie set the cookie into response
func (c *Client) SetCookie(key, value string) (error) {
	return nil
}

// GetCookie get the cookie from client
func (c *Client) GetCookie(key, value string) (string, error) {
	return "", nil
}

func (c *Client) Do(req *Request) (*Response, error) {
	resp, err := c.Client.Do(req.Request)
	return &Response{resp}, err
}
