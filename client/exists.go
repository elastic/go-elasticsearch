// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Exists - the get API allows to get a typed JSON document from the index based on its id. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-get.html for more info.
//
// documentType: the type of the document (use "_all" to fetch the first document matching the ID across all types).
//
// id: the document ID.
//
// index: the name of the index.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithParent, WithPreference, WithPretty, WithRealtime, WithRefresh, WithRouting, WithSource, WithSourceExclude, WithSourceInclude, WithSourceParam, WithStoredFields, WithVersion, WithVersionType, see the Option type in this package for more info.
func (c *Client) Exists(documentType string, id string, index string, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithHuman":         struct{}{},
		"WithParent":        struct{}{},
		"WithPreference":    struct{}{},
		"WithPretty":        struct{}{},
		"WithRealtime":      struct{}{},
		"WithRefresh":       struct{}{},
		"WithRouting":       struct{}{},
		"WithSource":        struct{}{},
		"WithSourceExclude": struct{}{},
		"WithSourceInclude": struct{}{},
		"WithSourceParam":   struct{}{},
		"WithStoredFields":  struct{}{},
		"WithVersion":       struct{}{},
		"WithVersionType":   struct{}{},
	}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "HEAD",
	}
	return c.client.Do(req)
}
