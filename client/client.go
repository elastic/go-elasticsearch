package api

import (
	"net/http"

	"github.com/elastic/elasticsearch-go/api"
)

// Client is the elasticsearch client. It implements the transport for it and embeds all the APIs.
// You can call any API on the client by simply instatiating the client.
// For example:
// c := client.New(...)
// c.Indices.Create(...)
// c.Index(...)
type Client struct {
	// API embeds the API objects
	*api.API

	// client is the underlying transport client
	client *http.Client
}

// New is the constructor for the Client
func New() *Client {
	httpClient := &http.Client{}
	return &Client{
		API:    api.New(httpClient),
		client: httpClient,
	}
}
