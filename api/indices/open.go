// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// Open - the open and close index APIs allow to close an index, and later on opening it. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-open-close.html for more info.
//
// index: a comma separated list of indices to open.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithErrorTrace, WithExpandWildcards, WithFilterPath, WithHuman, WithIgnoreUnavailable, WithMasterTimeout, WithPretty, WithSourceParam, WithTimeout, see the Option type in this package for more info.
func (i *Indices) Open(index []string, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllowNoIndices":    struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithMasterTimeout":     struct{}{},
		"WithPretty":            struct{}{},
		"WithSourceParam":       struct{}{},
		"WithTimeout":           struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "POST",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return i.transport.Do(req)
}
