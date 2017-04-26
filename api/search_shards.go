// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// SearchShards - the search shards api returns the indices and shards that a search request would be executed against. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-shards.html for more info.
//
// index: a comma-separated list of index names to search; use "_all" or empty string to perform the operation on all indices.
//
// documentType: a comma-separated list of document types to search; leave empty to perform the operation on all types.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithLocal, WithPreference, WithRouting, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) SearchShards(index []string, documentType []string, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithIndex":             struct{}{},
		"WithType":              struct{}{},
		"WithAllowNoIndices":    struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithLocal":             struct{}{},
		"WithPreference":        struct{}{},
		"WithRouting":           struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithPretty":            struct{}{},
		"WithSourceParam":       struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return a.transport.Do(req)
}
