// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Tasks - the task management API allows to retrieve information about the tasks currently executing on one or more nodes in the cluster. See http://www.elastic.co/guide/en/elasticsearch/reference/master/tasks.html for more info.
//
// options: optional parameters. Supports the following functional options: WithActions, WithDetailed, WithErrorTrace, WithFilterPath, WithFormat, WithH, WithHelp, WithHuman, WithNodeID, WithParentNode, WithParentTask, WithPretty, WithS, WithSourceParam, WithV, see the Option type in this package for more info.
func (c *Cat) Tasks(options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithActions":     struct{}{},
		"WithDetailed":    struct{}{},
		"WithErrorTrace":  struct{}{},
		"WithFilterPath":  struct{}{},
		"WithFormat":      struct{}{},
		"WithH":           struct{}{},
		"WithHelp":        struct{}{},
		"WithHuman":       struct{}{},
		"WithNodeID":      struct{}{},
		"WithParentNode":  struct{}{},
		"WithParentTask":  struct{}{},
		"WithPretty":      struct{}{},
		"WithS":           struct{}{},
		"WithSourceParam": struct{}{},
		"WithV":           struct{}{},
	}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "GET",
	}
	return c.client.Do(req)
}
