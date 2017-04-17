package http

import "net/http"

// Client is the HTTP client.
type Client struct {
	*http.Client
}

// NewClient is the constructor for the Client
func NewClient() *Client {
	return &Client{&http.Client{
		Transport: http.DefaultTransport,
	}}
}
