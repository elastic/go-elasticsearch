// Package transport implements the transport layer.
package transport

import (
	"net/http"
	"net/url"
)

// DefaultURL is the default host if none is configured
var DefaultURL = &url.URL{
	Scheme: "http",
	Host:   "localhost:9200",
}

const (
	// DefaultPort is the default port if none is configured
	DefaultPort = 9200
)

// Transport defines the transport for the connection to Elasticsearch.
type Transport struct {
	URL    *url.URL
	Client *http.Client
}

// New is the constructor for the Transport
func New() *Transport {
	return &Transport{
		URL:    DefaultURL,
		Client: &http.Client{},
	}
}

// Do sends an HTTP request
func (t *Transport) Do(req *http.Request) (*http.Response, error) {
	return t.Client.Do(req)
}
