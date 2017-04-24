// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package client

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Reindex - reindex does not attempt to set up the destination index. See https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-reindex.html for more info.
//
// body: the search definition using the Query DSL and the prototype for the index request.
//
// options: optional parameters. Supports the following functional options: WithRefresh, WithRequestsPerSecond, WithSlices, WithTimeout, WithWaitForActiveShards, WithWaitForCompletion, see the Option type in this package for more info.
func (c *Client) Reindex(body map[string]interface{}, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithRefresh":             struct{}{},
		"WithRequestsPerSecond":   struct{}{},
		"WithSlices":              struct{}{},
		"WithTimeout":             struct{}{},
		"WithWaitForActiveShards": struct{}{},
		"WithWaitForCompletion":   struct{}{},
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
