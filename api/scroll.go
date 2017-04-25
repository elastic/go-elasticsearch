// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Scroll - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-request-scroll.html for more info.
//
// body: the scroll ID if not passed by URL or query parameter.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithScroll, WithScrollID, WithScrollIDParam, WithSourceParam, see the Option type in this package for more info.
func (a *API) Scroll(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithPretty":        struct{}{},
		"WithScroll":        struct{}{},
		"WithScrollID":      struct{}{},
		"WithScrollIDParam": struct{}{},
		"WithSourceParam":   struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.Scheme,
			Host:   a.transport.Host,
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
