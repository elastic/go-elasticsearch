// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// ExistsAlias - APIs in Elasticsearch accept an index name when working against a specific index, and several indices when applicable. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-aliases.html for more info.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithErrorTrace, WithExpandWildcards, WithFilterPath, WithHuman, WithIgnoreUnavailable, WithIndex, WithLocal, WithName, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) ExistsAlias(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllowNoIndices":    struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithIndex":             struct{}{},
		"WithLocal":             struct{}{},
		"WithName":              struct{}{},
		"WithPretty":            struct{}{},
		"WithSourceParam":       struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "HEAD",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return i.transport.Do(req)
}
