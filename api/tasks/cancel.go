// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package tasks

import (
	"fmt"
	"net/http"
	"net/url"
)

// Cancel - the task management API allows to retrieve information about the tasks currently executing on one or more nodes in the cluster. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/tasks.html for more info.
//
// options: optional parameters. Supports the following functional options: WithActions, WithErrorTrace, WithFilterPath, WithHuman, WithNodeID, WithParentNode, WithParentTask, WithPretty, WithSourceParam, WithTaskID, see the Option type in this package for more info.
func (t *Tasks) Cancel(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithActions":     struct{}{},
		"WithErrorTrace":  struct{}{},
		"WithFilterPath":  struct{}{},
		"WithHuman":       struct{}{},
		"WithNodeID":      struct{}{},
		"WithParentNode":  struct{}{},
		"WithParentTask":  struct{}{},
		"WithPretty":      struct{}{},
		"WithSourceParam": struct{}{},
		"WithTaskID":      struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: t.transport.Scheme,
			Host:   t.transport.Host,
		},
		Method: "POST",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return t.transport.Do(req)
}
