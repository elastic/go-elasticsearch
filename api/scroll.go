// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Scroll - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-request-scroll.html for more info.
//
// options: optional parameters. Supports the following functional options: WithScrollID, WithScroll, WithScrollIDParam, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Scroll(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithScrollID":      struct{}{},
		"WithScroll":        struct{}{},
		"WithScrollIDParam": struct{}{},
		"WithBody":          struct{}{},
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
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
