// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Health - see http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-health.html for more info.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithFormat, WithH, WithHelp, WithHuman, WithLocal, WithMasterTimeout, WithPretty, WithS, WithSourceParam, WithTs, WithV, see the Option type in this package for more info.
func (c *Cat) Health(options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":    struct{}{},
		"WithFilterPath":    struct{}{},
		"WithFormat":        struct{}{},
		"WithH":             struct{}{},
		"WithHelp":          struct{}{},
		"WithHuman":         struct{}{},
		"WithLocal":         struct{}{},
		"WithMasterTimeout": struct{}{},
		"WithPretty":        struct{}{},
		"WithS":             struct{}{},
		"WithSourceParam":   struct{}{},
		"WithTs":            struct{}{},
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
