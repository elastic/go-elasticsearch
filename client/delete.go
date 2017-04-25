// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// Delete allows to delete a typed JSON document from a specific index based on its id. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-delete.html for more info.
//
// documentType: the type of the document.
//
// id: the document ID.
//
// index: the name of the index.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithParent, WithPretty, WithRefresh, WithRouting, WithSourceParam, WithTimeout, WithVersion, WithVersionType, WithWaitForActiveShards, see the Option type in this package for more info.
func (c *Client) Delete(documentType string, id string, index string, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":          struct{}{},
		"WithFilterPath":          struct{}{},
		"WithHuman":               struct{}{},
		"WithParent":              struct{}{},
		"WithPretty":              struct{}{},
		"WithRefresh":             struct{}{},
		"WithRouting":             struct{}{},
		"WithSourceParam":         struct{}{},
		"WithTimeout":             struct{}{},
		"WithVersion":             struct{}{},
		"WithVersionType":         struct{}{},
		"WithWaitForActiveShards": struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.Scheme,
			Host:   c.transport.Host,
		},
		Method: "DELETE",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return c.transport.Do(req)
}
