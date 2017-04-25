// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// Refresh allows to explicitly refresh one or more index, making all operations performed since the last refresh available for search. See http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-refresh.html for more info.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithErrorTrace, WithExpandWildcards, WithFilterPath, WithForce, WithHuman, WithIgnoreUnavailable, WithIndex, WithOperationThreading, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Refresh(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllowNoIndices":     struct{}{},
		"WithErrorTrace":         struct{}{},
		"WithExpandWildcards":    struct{}{},
		"WithFilterPath":         struct{}{},
		"WithForce":              struct{}{},
		"WithHuman":              struct{}{},
		"WithIgnoreUnavailable":  struct{}{},
		"WithIndex":              struct{}{},
		"WithOperationThreading": struct{}{},
		"WithPretty":             struct{}{},
		"WithSourceParam":        struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.Scheme,
			Host:   i.transport.Host,
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
