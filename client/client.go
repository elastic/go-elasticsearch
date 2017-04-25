// Package client is the starting point to establish a connection to Elasticsearch, as well as to invoke any of the
// available APIs.
//
// Here's a quick example:
//	c := client.New(client.WithHosts([]string{"https://elasticseach:9200"}))
//	body := map[string]interface{}{
//		"query": map[string]interface{}{
//			"term": map[string]interface{}{
//				"user": "kimchy",
//			},
//		},
//	}
//	resp, err := c.Search(body)
//	// ...
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
