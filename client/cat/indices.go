// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"
)

// Indices - see http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-indices.html for more info.
//
// options: optional parameters. Supports the following functional options: WithBytes, WithErrorTrace, WithFilterPath, WithFormat, WithH, WithHealth, WithHelp, WithHuman, WithIndex, WithLocal, WithMasterTimeout, WithPretty, WithPri, WithS, WithSourceParam, WithV, see the Option type in this package for more info.
func (c *Cat) Indices(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithBytes":         struct{}{},
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithFormat":        struct{}{},
		"WithH":             struct{}{},
		"WithHealth":        struct{}{},
		"WithHelp":          struct{}{},
		"WithHuman":         struct{}{},
		"WithIndex":         struct{}{},
		"WithLocal":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithPretty":        struct{}{},
		"WithPri":           struct{}{},
		"WithS":             struct{}{},
		"WithSourceParam":   struct{}{},
		"WithV":             struct{}{},
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
