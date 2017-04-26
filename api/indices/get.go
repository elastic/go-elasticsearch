// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// Get - the get index API allows to retrieve information about one or more indexes. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-get-index.html for more info.
//
// index: a comma-separated list of index names.
//
// options: optional parameters. Supports the following functional options: WithFeature, WithAllowNoIndices, WithExpandWildcards, WithFlatSettings, WithIgnoreUnavailable, WithIncludeDefaults, WithLocal, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Get(index []string, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithFeature":           struct{}{},
		"WithAllowNoIndices":    struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithFlatSettings":      struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithIncludeDefaults":   struct{}{},
		"WithLocal":             struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithPretty":            struct{}{},
		"WithSourceParam":       struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return i.transport.Do(req)
}
