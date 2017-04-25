// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Bulk makes it possible to perform many index/delete operations in a single API call. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html for more info.
//
// body: the operation definition and data (action-data pairs), separated by newlines.
//
// options: optional parameters. Supports the following functional options: WithType, WithTypeParam, WithErrorTrace, WithFields, WithFilterPath, WithHuman, WithIndex, WithPipeline, WithPretty, WithRefresh, WithRouting, WithSource, WithSourceExclude, WithSourceInclude, WithSourceParam, WithTimeout, WithWaitForActiveShards, see the Option type in this package for more info.
func (c *Client) Bulk(body map[string]interface{}, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithType":                struct{}{},
		"WithTypeParam":           struct{}{},
		"WithErrorTrace":          struct{}{},
		"WithFields":              struct{}{},
		"WithFilterPath":          struct{}{},
		"WithHuman":               struct{}{},
		"WithIndex":               struct{}{},
		"WithPipeline":            struct{}{},
		"WithPretty":              struct{}{},
		"WithRefresh":             struct{}{},
		"WithRouting":             struct{}{},
		"WithSource":              struct{}{},
		"WithSourceExclude":       struct{}{},
		"WithSourceInclude":       struct{}{},
		"WithSourceParam":         struct{}{},
		"WithTimeout":             struct{}{},
		"WithWaitForActiveShards": struct{}{},
	}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "POST",
	}
	return c.client.Do(req)
}
