// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package ingest

import (
	"fmt"
	"net/http"
	"net/url"
)

// Simulate - the ingest plugins extend Elasticsearch by providing additional ingest node capabilities. See https://www.elastic.co/guide/en/elasticsearch/plugins/5.x/ingest.html for more info.
//
// body: the simulate definition.
//
// options: optional parameters. Supports the following functional options: WithID, WithVerbose, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Ingest) Simulate(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithID":          struct{}{},
		"WithVerbose":     struct{}{},
		"WithErrorTrace":  struct{}{},
		"WithFilterPath":  struct{}{},
		"WithHuman":       struct{}{},
		"WithPretty":      struct{}{},
		"WithSourceParam": struct{}{},
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
