// Package client is the starting point to establish a connection to Elasticsearch, as well as to invoke any of the
// available APIs.
package client

import (
	"net/http"
)

// New is the constructor for the Client
func New(options ...TransportOption) *Client {
	httpClient := &http.Client{}
	c := &Client{
		client: httpClient,
	}
	c.addClients()
	return c
}
