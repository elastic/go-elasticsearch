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
// See the api package for all available methods and options.
package client

import (
	"github.com/elastic/goelasticsearch/api"
	"github.com/elastic/goelasticsearch/transport"
)

// Client is the top-level client.
type Client struct {
	// API is the root of the API. Since it's embedded, it allows to call APIs directly on the client.
	//	client := client.New()
	//	resp, error := client.Search(...)
	//	...
	*api.API

	transport *transport.Transport
}

// New is the constructor for the Client.
func New(options ...Option) *Client {
	t := transport.New()
	c := &Client{
		API:       api.New(t),
		transport: t,
	}
	for _, option := range options {
		option(c)
	}
	return c
}
