// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Delete - the delete index API allows to delete an existing index. See http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-delete-index.html for more info.
//
// index: a comma-separated list of indices to delete; use _all or * string to delete all indices.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, see the Option type in this package for more info.
func (i *Indices) Delete(index []string, options ...Option) (*http.Response, error) {
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
