// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Fielddata - see http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-fielddata.html for more info.
//
// options: optional parameters. Supports the following functional options: WithFields, WithBytes, WithFieldsParam, WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithV, see the Option type in this package for more info.
func (c *Cat) Fielddata(options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithFields":        struct{}{},
		"WithBytes":         struct{}{},
		"WithFieldsParam":   struct{}{},
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
