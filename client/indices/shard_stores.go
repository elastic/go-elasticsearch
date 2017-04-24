// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// ShardStores - provides store information for shard copies of indices. See http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-shards-stores.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithOperationThreading, WithStatus, see the Option type in this package for more info.
func (i *Indices) ShardStores(options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithIndex":              struct{}{},
		"WithAllowNoIndices":     struct{}{},
		"WithExpandWildcards":    struct{}{},
		"WithIgnoreUnavailable":  struct{}{},
		"WithOperationThreading": struct{}{},
		"WithStatus":             struct{}{},
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
	return i.client.Do(req)
}
