// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Mtermvectors - multi termvectors API allows to get multiple termvectors at once. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-multi-termvectors.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithFieldStatistics, WithFields, WithIds, WithOffsets, WithParent, WithPayloads, WithPositions, WithPreference, WithRealtime, WithRouting, WithTermStatistics, WithVersion, WithVersionType, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Mtermvectors(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithIndex":           struct{}{},
		"WithType":            struct{}{},
		"WithFieldStatistics": struct{}{},
		"WithFields":          struct{}{},
		"WithIds":             struct{}{},
		"WithOffsets":         struct{}{},
		"WithParent":          struct{}{},
		"WithPayloads":        struct{}{},
		"WithPositions":       struct{}{},
		"WithPreference":      struct{}{},
		"WithRealtime":        struct{}{},
		"WithRouting":         struct{}{},
		"WithTermStatistics":  struct{}{},
		"WithVersion":         struct{}{},
		"WithVersionType":     struct{}{},
		"WithBody":            struct{}{},
		"WithErrorTrace":      struct{}{},
		"WithFilterPath":      struct{}{},
		"WithHuman":           struct{}{},
		"WithPretty":          struct{}{},
		"WithSourceParam":     struct{}{},
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
