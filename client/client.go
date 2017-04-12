package api

import (
	"github.com/elastic/elasticsearch-go/api"
	"github.com/elastic/elasticsearch-go/client/http"
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

	// httpClient is the underlying transport client
	httpClient *http.Client
}

// New is the constructor for the Client
func New() *Client {
	c := &Client{httpClient: http.NewClient()}
	c.API = api.New(c.httpClient)
	return c
}
