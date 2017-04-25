// Package transport implements the transport layer.
package transport

import "net/http"

// Transport defines the transport for the connection to Elasticsearch.
type Transport struct {
	Scheme string
	Host   string
	client *http.Client
}

// New is the constructor for the Transport
func New() *Transport {
	return &Transport{
		client: &http.Client{},
	}
}

// Do sends an HTTP request
func (t *Transport) Do(req *http.Request) (*http.Response, error) {
	return t.client.Do(req)
}
