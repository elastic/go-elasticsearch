// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetTemplate - index templates allow you to define templates that will automatically be applied when new indices are created. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-templates.html for more info.
//
// name: the comma separated names of the index templates.
//
// options: optional parameters. Supports the following functional options: WithName, WithFlatSettings, WithLocal, WithMasterTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) GetTemplate(name []string, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithName":          struct{}{},
		"WithFlatSettings":  struct{}{},
		"WithLocal":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithPretty":        struct{}{},
		"WithSourceParam":   struct{}{},
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
