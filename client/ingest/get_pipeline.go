// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package ingest

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetPipeline - the ingest plugins extend Elasticsearch by providing additional ingest node capabilities. See https://www.elastic.co/guide/en/elasticsearch/plugins/master/ingest.html for more info.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithID, WithMasterTimeout, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Ingest) GetPipeline(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithID":            struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.Scheme,
			Host:   i.transport.Host,
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
