package estransport // import "github.com/elastic/go-elasticsearch/estransport"

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Interface defines the interface for HTTP client.
//
type Interface interface {
	Perform(*http.Request) (*http.Response, error)
}

// Config represents the configuration of HTTP client.
//
type Config struct {
	URLs      []*url.URL
	Transport http.RoundTripper

	LoggerOutput io.Writer
	LoggerFormat string
	LoggerFunc   func(*http.Request, *http.Response)
}

// Client represents the HTTP client.
//
type Client struct {
	urls      []*url.URL
	transport http.RoundTripper
	selector  Selector

	loggerOutput io.Writer
	loggerFormat string
	loggerFunc   func(*http.Request, *http.Response)
}

// New creates new HTTP client.
//
// http.DefaultTransport will be used if no transport is passed in the configuration.
//
func New(cfg Config) *Client {
	if cfg.Transport == nil {
		cfg.Transport = http.DefaultTransport
	}

	return &Client{
		urls:      cfg.URLs,
		transport: cfg.Transport,
		selector:  NewRoundRobinSelector(cfg.URLs...),

		loggerOutput: cfg.LoggerOutput,
		loggerFormat: cfg.LoggerFormat,
		loggerFunc:   cfg.LoggerFunc,
	}
}

// Perform executes the request and returns a response or error.
//
func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	u, err := c.getURL()
	if err != nil {
		return nil, fmt.Errorf("cannot get URL: %s", err)
	}

	c.setURL(u, req)
	c.setBasicAuth(u, req)

	s := time.Now().UTC()
	res, err := c.transport.RoundTrip(req)
	d := time.Now().UTC().Sub(s)

	if c.loggerOutput != nil {
		fmt.Fprintf(c.loggerOutput, "%s %s %s [status:%d request:%s]\n",
			time.Now().Format(time.RFC3339),
			req.Method,
			req.URL.String(),
			res.StatusCode,
			d.Truncate(time.Millisecond),
		)
		if req.Body != nil {
			// TODO(karmi): Capture the request body before performing the request
			fmt.Fprintln(c.loggerOutput, "> TODO: Capture and print request body")
		}
		if err != nil {
			fmt.Fprintf(c.loggerOutput, "! ERROR: %v", err)
		} else {
			if res.Body != nil {
				body, err := ioutil.ReadAll(res.Body)
				if err == nil {
					defer func() { res.Body = ioutil.NopCloser(bytes.NewReader(body)) }()
					defer func() { res.Body.Close() }()
					for _, line := range strings.Split(string(body), "\n") {
						if line != "" {
							log.Printf("< %s\n", line)
						}
					}
				}
			}
		}
	}

	// TODO(karmi): Wrap error
	return res, err
}

// URLs returns a list of transport URLs.
//
func (c *Client) URLs() []*url.URL {
	return c.urls
}

func (c *Client) getURL() (*url.URL, error) {
	return c.selector.Select()
}

func (c *Client) setURL(u *url.URL, req *http.Request) *http.Request {
	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host

	if u.Path != "" {
		var b strings.Builder
		b.Grow(len(u.Path) + len(req.URL.Path))
		b.WriteString(u.Path)
		b.WriteString(req.URL.Path)
		req.URL.Path = b.String()
	}

	return req
}

func (c *Client) setBasicAuth(u *url.URL, req *http.Request) *http.Request {
	if u.User != nil {
		password, _ := u.User.Password()
		req.SetBasicAuth(u.User.Username(), password)
	}
	return req
}
