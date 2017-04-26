// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// SearchTemplate - see https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html for more info.
//
// body: the search definition template and its params.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithType, WithErrorTrace, WithExpandWildcards, WithExplain, WithFilterPath, WithHuman, WithIgnoreUnavailable, WithIndex, WithPreference, WithPretty, WithProfile, WithRouting, WithScroll, WithSearchType, WithSourceParam, see the Option type in this package for more info.
func (a *API) SearchTemplate(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllowNoIndices":    struct{}{},
		"WithType":              struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithExplain":           struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithIndex":             struct{}{},
		"WithPreference":        struct{}{},
		"WithPretty":            struct{}{},
		"WithProfile":           struct{}{},
		"WithRouting":           struct{}{},
		"WithScroll":            struct{}{},
		"WithSearchType":        struct{}{},
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
