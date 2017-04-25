// Package transport implements the transport layer.
package transport

import "net/http"

const (
	// DefaultScheme is the default scheme used for HTTP connection if none is specified for a host
	DefaultScheme = "http"

	// DefaultHost is the default host if none is configured
	DefaultHost = "localhost:9200"
)

// Transport defines the transport for the connection to Elasticsearch.
type Transport struct {
	Scheme string
	Host   string
	client *http.Client
}

// New is the constructor for the Transport
func New() *Transport {
	return &Transport{
		Scheme: DefaultScheme,
		Host:   DefaultHost,
		client: &http.Client{},
	}
}

// Do sends an HTTP request
func (t *Transport) Do(req *http.Request) (*http.Response, error) {
	return t.client.Do(req)
}
