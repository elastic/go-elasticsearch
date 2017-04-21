package client

import (
	"net/http"
)

// New is the constructor for the Client
func New() *Client {
	httpClient := &http.Client{}
	c := &Client{
		client: httpClient,
	}
	c.addClients()
	return c
}
