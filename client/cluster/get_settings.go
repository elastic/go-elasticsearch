// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetSettings - allows to update cluster wide specific settings. See http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-update-settings.html for more info.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithFlatSettings, WithHuman, WithIncludeDefaults, WithMasterTimeout, WithPretty, WithSourceParam, WithTimeout, see the Option type in this package for more info.
func (c *Cluster) GetSettings(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":      struct{}{},
		"WithFilterPath":      struct{}{},
		"WithFlatSettings":    struct{}{},
		"WithHuman":           struct{}{},
		"WithIncludeDefaults": struct{}{},
		"WithMasterTimeout":   struct{}{},
		"WithPretty":          struct{}{},
		"WithSourceParam":     struct{}{},
		"WithTimeout":         struct{}{},
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
