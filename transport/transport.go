// Package transport implements the transport layer.
package transport

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// DefaultURL is the default host if none is configured.
var DefaultURL = &url.URL{
	Scheme: "http",
	Host:   "localhost:9200",
}

const (
	// DefaultPort is the default port if none is configured.
	DefaultPort = 9200
)

// Transport defines the transport for the connection to Elasticsearch.
type Transport struct {
	URL    *url.URL
	client *http.Client
}

// New is the constructor for the Transport.
func New() *Transport {
	return &Transport{
		URL:    DefaultURL,
		client: &http.Client{},
	}
}

// Do sends an HTTP request.
func (t *Transport) Do(req *http.Request) (*http.Response, error) {
	return t.client.Do(req)
}

// DecodeResponseBody is a helper to decode the JSON body of an HTTP response.
func DecodeResponseBody(resp *http.Response) (map[string]interface{}, error) {
	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(body); err != nil {
		return nil, err
	}
	return body, nil
}

// JSONField is a helper to extract an object from JSON structure. It understands the dotted notation.
func JSONField(body map[string]interface{}, name string) (interface{}, error) {
	path := strings.Split(name, ".")
	var field interface{}
	field = body
	for _, n := range path {
		field = field.(map[string]interface{})[n]
	}
	return field, nil
}
