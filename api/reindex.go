// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Reindex - reindex does not attempt to set up the destination index. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-reindex.html for more info.
//
// body: the search definition using the Query DSL and the prototype for the index request.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithRefresh, WithRequestsPerSecond, WithSlices, WithSourceParam, WithTimeout, WithWaitForActiveShards, WithWaitForCompletion, see the Option type in this package for more info.
func (a *API) Reindex(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":          struct{}{},
		"WithFilterPath":          struct{}{},
		"WithHuman":               struct{}{},
		"WithPretty":              struct{}{},
		"WithRefresh":             struct{}{},
		"WithRequestsPerSecond":   struct{}{},
		"WithSlices":              struct{}{},
		"WithSourceParam":         struct{}{},
		"WithTimeout":             struct{}{},
		"WithWaitForActiveShards": struct{}{},
		"WithWaitForCompletion":   struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "POST",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return a.transport.Do(req)
}
