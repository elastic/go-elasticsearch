// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// DeleteAlias - apis in Elasticsearch accept an index name when working against a specific index, and several indices when applicable. See http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html for more info.
//
// index: a comma-separated list of index names (supports wildcards); use _all for all indices.
//
// name: a comma-separated list of aliases to delete (supports wildcards); use _all to delete all aliases for the specified indices.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, see the Option type in this package for more info.
func (i *Indices) DeleteAlias(index []string, name []string, options ...Option) (*http.Response, error) {
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
