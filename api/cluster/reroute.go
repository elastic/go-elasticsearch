// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"
	"net/url"
)

// Reroute - the reroute command allows to explicitly execute a cluster reroute allocation command including specific commands. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-reroute.html for more info.
//
// body: the definition of "commands" to perform ("move", "cancel", "allocate").
//
// options: optional parameters. Supports the following functional options: WithDryRun, WithErrorTrace, WithExplain, WithFilterPath, WithHuman, WithMasterTimeout, WithMetric, WithPretty, WithRetryFailed, WithSourceParam, WithTimeout, see the Option type in this package for more info.
func (c *Cluster) Reroute(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithDryRun":        struct{}{},
		"WithErrorTrace":    struct{}{},
		"WithExplain":       struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithMetric":        struct{}{},
		"WithPretty":        struct{}{},
		"WithRetryFailed":   struct{}{},
		"WithSourceParam":   struct{}{},
		"WithTimeout":       struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "POST",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return c.transport.Do(req)
}
