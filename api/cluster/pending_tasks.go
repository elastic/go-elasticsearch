// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"
	"net/url"
)

// PendingTasks - the pending cluster tasks API returns a list of any cluster-level changes (e.g. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-pending.html for more info.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithLocal, WithMasterTimeout, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cluster) PendingTasks(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithLocal":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.Scheme,
			Host:   c.transport.Host,
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
