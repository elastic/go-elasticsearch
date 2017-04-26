// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Mtermvectors - multi termvectors API allows to get multiple termvectors at once. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-multi-termvectors.html for more info.
//
// body: define ids, documents, parameters or a list of parameters per document here. You must at least provide a list of document ids. See documentation.
//
// options: optional parameters. Supports the following functional options: WithType, WithErrorTrace, WithFieldStatistics, WithFields, WithFilterPath, WithHuman, WithIds, WithIndex, WithOffsets, WithParent, WithPayloads, WithPositions, WithPreference, WithPretty, WithRealtime, WithRouting, WithSourceParam, WithTermStatistics, WithVersion, WithVersionType, see the Option type in this package for more info.
func (a *API) Mtermvectors(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithType":            struct{}{},
		"WithErrorTrace":      struct{}{},
		"WithFieldStatistics": struct{}{},
		"WithFields":          struct{}{},
		"WithFilterPath":      struct{}{},
		"WithHuman":           struct{}{},
		"WithIds":             struct{}{},
		"WithIndex":           struct{}{},
		"WithOffsets":         struct{}{},
		"WithParent":          struct{}{},
		"WithPayloads":        struct{}{},
		"WithPositions":       struct{}{},
		"WithPreference":      struct{}{},
		"WithPretty":          struct{}{},
		"WithRealtime":        struct{}{},
		"WithRouting":         struct{}{},
		"WithSourceParam":     struct{}{},
		"WithTermStatistics":  struct{}{},
		"WithVersion":         struct{}{},
		"WithVersionType":     struct{}{},
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
