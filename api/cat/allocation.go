// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"
)

// Allocation - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-allocation.html for more info.
//
// options: optional parameters. Supports the following functional options: WithNodeID, WithBytes, WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithV, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Allocation(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithNodeID":        struct{}{},
		"WithBytes":         struct{}{},
		"WithFormat":        struct{}{},
		"WithH":             struct{}{},
		"WithHelp":          struct{}{},
		"WithLocal":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithS":             struct{}{},
		"WithV":             struct{}{},
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
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
