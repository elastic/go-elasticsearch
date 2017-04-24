// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// DeleteTemplate - index templates allow you to define templates that will automatically be applied when new indices are created. See http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html for more info.
//
// name: the name of the template.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, see the Option type in this package for more info.
func (i *Indices) DeleteTemplate(name string, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithMasterTimeout": struct{}{},
		"WithTimeout":       struct{}{},
	}
	for _, option := range options {
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "DELETE",
	}
	return i.client.Do(req)
}
