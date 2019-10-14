// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package estransport

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/internal/version"
)

// Version returns the package version as a string.
//
const Version = version.Client

var (
	userAgent   string
	reGoVersion = regexp.MustCompile(`go(\d+\.\d+\..+)`)

	defaultMaxRetries    = 3
	defaultRetryOnStatus = [...]int{502, 503, 504}
)

func init() {
	userAgent = initUserAgent()
}

// Interface defines the interface for HTTP client.
//
type Interface interface {
	Perform(*http.Request) (*http.Response, error)
}

// Config represents the configuration of HTTP client.
//
type Config struct {
	URLs     []*url.URL
	Username string
	Password string
	APIKey   string

	RetryOnStatus        []int
	DisableRetry         bool
	EnableRetryOnTimeout bool
	MaxRetries           int
	RetryBackoff         func(attempt int) time.Duration

	Transport http.RoundTripper
	Logger    Logger
}

// Client represents the HTTP client.
//
type Client struct {
	urls     []*url.URL
	username string
	password string
	apikey   string

	retryOnStatus        []int
	disableRetry         bool
	enableRetryOnTimeout bool
	maxRetries           int
	retryBackoff         func(attempt int) time.Duration

	transport http.RoundTripper
	selector  Selector
	logger    Logger
}

// New creates new HTTP client.
//
// http.DefaultTransport will be used if no transport is passed in the configuration.
//
func New(cfg Config) *Client {
	if cfg.Transport == nil {
		cfg.Transport = http.DefaultTransport
	}

	if len(cfg.RetryOnStatus) == 0 {
		cfg.RetryOnStatus = defaultRetryOnStatus[:]
	}

	if cfg.MaxRetries == 0 {
		cfg.MaxRetries = defaultMaxRetries
	}

	return &Client{
		urls:     cfg.URLs,
		username: cfg.Username,
		password: cfg.Password,
		apikey:   cfg.APIKey,

		retryOnStatus:        cfg.RetryOnStatus,
		disableRetry:         cfg.DisableRetry,
		enableRetryOnTimeout: cfg.EnableRetryOnTimeout,
		maxRetries:           cfg.MaxRetries,
		retryBackoff:         cfg.RetryBackoff,

		transport: cfg.Transport,
		selector:  NewRoundRobinSelector(cfg.URLs...),
		logger:    cfg.Logger,
	}
}

// Perform executes the request and returns a response or error.
//
func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	var (
		res *http.Response
		err error
	)

	// Update request
	//
	c.setReqUserAgent(req)

	if req.Body != nil && req.Body != http.NoBody && req.GetBody == nil {
		if !c.disableRetry || (c.logger != nil && c.logger.RequestBodyEnabled()) {
			var buf bytes.Buffer
			buf.ReadFrom(req.Body)
			req.GetBody = func() (io.ReadCloser, error) {
				r := buf
				return ioutil.NopCloser(&r), nil
			}
			if req.Body, err = req.GetBody(); err != nil {
				return nil, fmt.Errorf("cannot get request body: %s", err)
			}
		}
	}

	for i := 1; i <= c.maxRetries; i++ {
		var (
			nodeURL     *url.URL
			shouldRetry bool
		)

		// Get URL from the Selector
		//
		nodeURL, err = c.getURL()
		if err != nil {
			// TODO(karmi): Log error
			return nil, fmt.Errorf("cannot get URL: %s", err)
		}

		// Update request
		//
		c.setReqURL(nodeURL, req)
		c.setReqAuth(nodeURL, req)

		if !c.disableRetry && i > 1 && req.Body != nil && req.Body != http.NoBody {
			body, err := req.GetBody()
			if err != nil {
				return nil, fmt.Errorf("cannot get request body: %s", err)
			}
			req.Body = body
		}

		// Set up time measures and execute the request
		//
		start := time.Now().UTC()
		res, err = c.transport.RoundTrip(req)
		dur := time.Since(start)

		// Log request and response
		//
		if c.logger != nil {
			if c.logger.RequestBodyEnabled() && req.Body != nil && req.Body != http.NoBody {
				req.Body, _ = req.GetBody()
			}
			c.logRoundTrip(req, res, err, start, dur)
		}

		// Retry only on network errors, but don't retry on timeout errors, unless configured
		//
		if err != nil {
			if err, ok := err.(net.Error); ok {
				if (!err.Timeout() || c.enableRetryOnTimeout) && !c.disableRetry {
					shouldRetry = true
				}
			}
		}

		// Retry on configured response statuses
		//
		if res != nil && !c.disableRetry {
			for _, code := range c.retryOnStatus {
				if res.StatusCode == code {
					shouldRetry = true
				}
			}
		}

		// Break if retry should not be performed
		//
		if !shouldRetry {
			break
		}

		// Delay the retry if a backoff function is configured
		//
		if c.retryBackoff != nil {
			time.Sleep(c.retryBackoff(i))
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

func (c *Client) setReqURL(u *url.URL, req *http.Request) *http.Request {
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

func (c *Client) setReqAuth(u *url.URL, req *http.Request) *http.Request {
	if _, ok := req.Header["Authorization"]; !ok {
		if u.User != nil {
			password, _ := u.User.Password()
			req.SetBasicAuth(u.User.Username(), password)
			return req
		}

		if c.apikey != "" {
			var b bytes.Buffer
			b.Grow(len("APIKey ") + len(c.apikey))
			b.WriteString("APIKey ")
			b.WriteString(c.apikey)
			req.Header.Set("Authorization", b.String())
			return req
		}

		if c.username != "" && c.password != "" {
			req.SetBasicAuth(c.username, c.password)
			return req
		}
	}

	return req
}

func (c *Client) setReqUserAgent(req *http.Request) *http.Request {
	req.Header.Set("User-Agent", userAgent)
	return req
}

func (c *Client) logRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) {
	var dupRes http.Response
	if res != nil {
		dupRes = *res
	}
	if c.logger.ResponseBodyEnabled() {
		if res != nil && res.Body != nil && res.Body != http.NoBody {
			b1, b2, _ := duplicateBody(res.Body)
			dupRes.Body = b1
			res.Body = b2
		}
	}
	c.logger.LogRoundTrip(req, &dupRes, err, start, dur) // errcheck exclude
}

func initUserAgent() string {
	var b strings.Builder

	b.WriteString("go-elasticsearch")
	b.WriteRune('/')
	b.WriteString(Version)
	b.WriteRune(' ')
	b.WriteRune('(')
	b.WriteString(runtime.GOOS)
	b.WriteRune(' ')
	b.WriteString(runtime.GOARCH)
	b.WriteString("; ")
	b.WriteString("Go ")
	if v := reGoVersion.ReplaceAllString(runtime.Version(), "$1"); v != "" {
		b.WriteString(v)
	} else {
		b.WriteString(runtime.Version())
	}
	b.WriteRune(')')

	return b.String()
}
