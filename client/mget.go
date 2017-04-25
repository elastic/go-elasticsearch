// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"net/url"
)

// Mget - multi GET API allows to get multiple documents based on an index, type (optional) and id (and possibly routing). See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-multi-get.html for more info.
//
// body: document identifiers; can be either "docs" (containing full document information) or "ids" (when index and type is provided in the URL.
//
// options: optional parameters. Supports the following functional options: WithType, WithErrorTrace, WithFilterPath, WithHuman, WithIndex, WithPreference, WithPretty, WithRealtime, WithRefresh, WithRouting, WithSource, WithSourceExclude, WithSourceInclude, WithSourceParam, WithStoredFields, see the Option type in this package for more info.
func (c *Client) Mget(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithType":          struct{}{},
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithIndex":         struct{}{},
		"WithPreference":    struct{}{},
		"WithPretty":        struct{}{},
		"WithRealtime":      struct{}{},
		"WithRefresh":       struct{}{},
		"WithRouting":       struct{}{},
		"WithSource":        struct{}{},
		"WithSourceExclude": struct{}{},
		"WithSourceInclude": struct{}{},
		"WithSourceParam":   struct{}{},
		"WithStoredFields":  struct{}{},
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
