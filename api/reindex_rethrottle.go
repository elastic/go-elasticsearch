// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// ReindexRethrottle - reindex does not attempt to set up the destination index. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-reindex.html for more info.
//
// requestsPerSecond: the throttle to set on this request in floating sub-requests per second. -1 means set no throttle.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, WithTaskID, see the Option type in this package for more info.
func (a *API) ReindexRethrottle(requestsPerSecond int, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":  struct{}{},
		"WithFilterPath":  struct{}{},
		"WithHuman":       struct{}{},
		"WithPretty":      struct{}{},
		"WithSourceParam": struct{}{},
		"WithTaskID":      struct{}{},
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
