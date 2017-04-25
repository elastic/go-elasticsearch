// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// Msearch - the multi search API allows to execute several search requests within the same API. See http://www.elastic.co/guide/en/elasticsearch/reference/master/search-multi-search.html for more info.
//
// body: the request definitions (metadata-search request definition pairs), separated by newlines.
//
// options: optional parameters. Supports the following functional options: WithType, WithErrorTrace, WithFilterPath, WithHuman, WithIndex, WithMaxConcurrentSearches, WithPretty, WithSearchType, WithSourceParam, WithTypedKeys, see the Option type in this package for more info.
func (c *Client) Msearch(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithType":                  struct{}{},
		"WithErrorTrace":            struct{}{},
		"WithFilterPath":            struct{}{},
		"WithHuman":                 struct{}{},
		"WithIndex":                 struct{}{},
		"WithMaxConcurrentSearches": struct{}{},
		"WithPretty":                struct{}{},
		"WithSearchType":            struct{}{},
		"WithSourceParam":           struct{}{},
		"WithTypedKeys":             struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.Scheme,
			Host:   c.transport.Host,
		},
		Method: "GET",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return c.transport.Do(req)
}
