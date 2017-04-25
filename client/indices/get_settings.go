// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetSettings - see http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-settings.html for more info.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithErrorTrace, WithExpandWildcards, WithFilterPath, WithFlatSettings, WithHuman, WithIgnoreUnavailable, WithIncludeDefaults, WithIndex, WithLocal, WithName, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) GetSettings(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllowNoIndices":    struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithFilterPath":        struct{}{},
		"WithFlatSettings":      struct{}{},
		"WithHuman":             struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithIncludeDefaults":   struct{}{},
		"WithIndex":             struct{}{},
		"WithLocal":             struct{}{},
		"WithName":              struct{}{},
		"WithPretty":            struct{}{},
		"WithSourceParam":       struct{}{},
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
