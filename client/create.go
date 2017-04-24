// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Create - the index API adds or updates a typed JSON document in a specific index, making it searchable. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html for more info.
//
// id: document ID.
//
// index: the name of the index.
//
// documentType: the type of the document.
//
// body: the document.
//
// options: optional parameters. Supports the following functional options: WithParent, WithPipeline, WithRefresh, WithRouting, WithTimeout, WithTimestamp, WithTTL, WithVersion, WithVersionType, WithWaitForActiveShards, see the Option type in this package for more info.
func (c *Client) Create(id string, index string, documentType string, body map[string]interface{}, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithParent":              struct{}{},
		"WithPipeline":            struct{}{},
		"WithRefresh":             struct{}{},
		"WithRouting":             struct{}{},
		"WithTimeout":             struct{}{},
		"WithTimestamp":           struct{}{},
		"WithTTL":                 struct{}{},
		"WithVersion":             struct{}{},
		"WithVersionType":         struct{}{},
		"WithWaitForActiveShards": struct{}{},
	}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "PUT",
	}
	return c.client.Do(req)
}
