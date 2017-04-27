// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package tasks

import (
	"fmt"
	"net/http"
	"net/url"
)

// List - the task management API allows to retrieve information about the tasks currently executing on one or more nodes in the cluster. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/tasks.html for more info.
//
// options: optional parameters. Supports the following functional options: WithActions, WithDetailed, WithGroupBy, WithNodeID, WithParentNode, WithParentTask, WithWaitForCompletion, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (t *Tasks) List(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithActions":           struct{}{},
		"WithDetailed":          struct{}{},
		"WithGroupBy":           struct{}{},
		"WithNodeID":            struct{}{},
		"WithParentNode":        struct{}{},
		"WithParentTask":        struct{}{},
		"WithWaitForCompletion": struct{}{},
		"WithErrorTrace":        struct{}{},
		"WithFilterPath":        struct{}{},
		"WithHuman":             struct{}{},
		"WithPretty":            struct{}{},
		"WithSourceParam":       struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: t.transport.URL.Scheme,
			Host:   t.transport.URL.Host,
		},
		Method: "GET",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return t.transport.Do(req)
}
