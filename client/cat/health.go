// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"
)

// Health - see http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-health.html for more info.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithFormat, WithH, WithHelp, WithHuman, WithLocal, WithMasterTimeout, WithPretty, WithS, WithSourceParam, WithTs, WithV, see the Option type in this package for more info.
func (c *Cat) Health(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithFormat":        struct{}{},
		"WithH":             struct{}{},
		"WithHelp":          struct{}{},
		"WithHuman":         struct{}{},
		"WithLocal":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithPretty":        struct{}{},
		"WithS":             struct{}{},
		"WithSourceParam":   struct{}{},
		"WithTs":            struct{}{},
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
