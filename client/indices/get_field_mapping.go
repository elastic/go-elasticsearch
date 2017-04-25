// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetFieldMapping - the get field mapping API allows you to retrieve mapping definitions for one or more fields. See http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-field-mapping.html for more info.
//
// fields: a comma-separated list of fields.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithType, WithErrorTrace, WithExpandWildcards, WithFilterPath, WithHuman, WithIgnoreUnavailable, WithIncludeDefaults, WithIndex, WithLocal, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) GetFieldMapping(fields []string, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithAllowNoIndices":    struct{}{},
		"WithType":              struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithExpandWildcards":   struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithIncludeDefaults":   struct{}{},
		"WithIndex":             struct{}{},
		"WithLocal":             struct{}{},
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
