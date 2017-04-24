// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Upgrade - see http://www.elastic.co/guide/en/elasticsearch/reference/master/indices-upgrade.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithOnlyAncientSegments, WithWaitForCompletion, see the Option type in this package for more info.
func (i *Indices) Upgrade(options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithIndex":               struct{}{},
		"WithAllowNoIndices":      struct{}{},
		"WithExpandWildcards":     struct{}{},
		"WithIgnoreUnavailable":   struct{}{},
		"WithOnlyAncientSegments": struct{}{},
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
	return i.client.Do(req)
}
