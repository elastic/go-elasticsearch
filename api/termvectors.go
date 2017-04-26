// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Termvectors - returns information and statistics on terms in the fields of a particular document. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-termvectors.html for more info.
//
// documentType: the type of the document.
//
// index: the index in which the document resides.
//
// body: define parameters and or supply a document to get termvectors for. See documentation.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFieldStatistics, WithFields, WithFilterPath, WithHuman, WithID, WithOffsets, WithParent, WithPayloads, WithPositions, WithPreference, WithPretty, WithRealtime, WithRouting, WithSourceParam, WithTermStatistics, WithVersion, WithVersionType, see the Option type in this package for more info.
func (a *API) Termvectors(documentType string, index string, body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":      struct{}{},
		"WithFieldStatistics": struct{}{},
		"WithFields":          struct{}{},
		"WithFilterPath":      struct{}{},
		"WithHuman":           struct{}{},
		"WithID":              struct{}{},
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
