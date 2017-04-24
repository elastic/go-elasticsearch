// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package tasks

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Get - the task management API allows to retrieve information about the tasks currently executing on one or more nodes in the cluster. See http://www.elastic.co/guide/en/elasticsearch/reference/master/tasks.html for more info.
//
// options: optional parameters. Supports the following functional options: WithTaskID, WithWaitForCompletion, see the Option type in this package for more info.
func (t *Tasks) Get(options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithTaskID":            struct{}{},
		"WithWaitForCompletion": struct{}{},
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
	return t.client.Do(req)
}
