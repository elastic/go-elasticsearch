// Package transport implements the transport layer.
package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/util"
)

// DefaultURL is the default host if none is configured.
var DefaultURL = &url.URL{
	Scheme: "http",
	Host:   "localhost:9200",
}

const (
	// DefaultPort is the default port if none is configured.
	DefaultPort = 9200
)

// Transport defines the transport for the connection to Elasticsearch.
type Transport struct {
	URL    *url.URL
	client *http.Client
}

// New is the constructor for the Transport.
func New() *Transport {
	return &Transport{
		URL:    DefaultURL,
		client: &http.Client{},
	}
}

// Do sends an HTTP request.
func (t *Transport) Do(req *Request) (*http.Response, error) {
	resp, err := t.client.Do(req.Request)
	if err == nil && resp.StatusCode >= 300 {
		if _, ok := req.IgnoredStatuses[resp.StatusCode]; !ok {
			err = fmt.Errorf("%d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
		}
	}
	return resp, err
}

// DecodeResponseBody is a helper to decode the JSON body of an HTTP response.
func DecodeResponseBody(resp *http.Response) (util.MapStr, error) {
	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return nil, err
	}
	return util.MapStr(body), nil
}

// Request wraps an HTTP request with some options.
type Request struct {
	// Request is the HTTP request
	*http.Request
	// IgnoredStatuses contains the HTTP status codes to ignore. By default, all codes >= 300 result in an error.
	IgnoredStatuses map[int]struct{}
}

// NewRequest instantiates a new requests for a transport.
func (t *Transport) NewRequest(method string) *Request {
	return &Request{
		Request: &http.Request{
			URL:    t.URL,
			Method: method,
		},
		IgnoredStatuses: map[int]struct{}{},
	}
}
