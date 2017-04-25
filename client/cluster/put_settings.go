// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cluster

import (
	"fmt"
	"net/http"
	"net/url"
)

// PutSettings - allows to update cluster wide specific settings. See http://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-update-settings.html for more info.
//
// body: the settings to be updated. Can be either "transient" or "persistent" (survives cluster restart).
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithFlatSettings, WithHuman, WithMasterTimeout, WithPretty, WithSourceParam, WithTimeout, see the Option type in this package for more info.
func (c *Cluster) PutSettings(body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithFlatSettings":  struct{}{},
		"WithHuman":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
		"WithTimeout":       struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.Scheme,
			Host:   c.transport.Host,
		},
		Method: "PUT",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return c.transport.Do(req)
}
