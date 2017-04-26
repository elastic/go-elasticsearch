// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"
	"net/url"
)

// Health - the cluster health API allows to get a very simple status on the health of the cluster. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-health.html for more info.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithIndex, WithLevel, WithLocal, WithMasterTimeout, WithPretty, WithSourceParam, WithTimeout, WithWaitForActiveShards, WithWaitForEvents, WithWaitForNoRelocatingShards, WithWaitForNodes, WithWaitForStatus, see the Option type in this package for more info.
func (c *Cluster) Health(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":                struct{}{},
		"WithFilterPath":                struct{}{},
		"WithHuman":                     struct{}{},
		"WithIndex":                     struct{}{},
		"WithLevel":                     struct{}{},
		"WithLocal":                     struct{}{},
		"WithMasterTimeout":             struct{}{},
		"WithPretty":                    struct{}{},
		"WithSourceParam":               struct{}{},
		"WithTimeout":                   struct{}{},
		"WithWaitForActiveShards":       struct{}{},
		"WithWaitForEvents":             struct{}{},
		"WithWaitForNoRelocatingShards": struct{}{},
		"WithWaitForNodes":              struct{}{},
		"WithWaitForStatus":             struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "GET",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return c.transport.Do(req)
}
