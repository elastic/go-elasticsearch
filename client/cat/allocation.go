// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Allocation - see http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-allocation.html for more info.
//
// options: optional parameters. Supports the following functional options: WithNodeID, WithBytes, WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithV, see the Option type in this package for more info.
func (c *Cat) Allocation(options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithNodeID":        struct{}{},
		"WithBytes":         struct{}{},
		"WithFormat":        struct{}{},
		"WithH":             struct{}{},
		"WithHelp":          struct{}{},
		"WithLocal":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithS":             struct{}{},
		"WithV":             struct{}{},
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
