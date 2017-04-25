// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package ingest

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Simulate - the ingest plugins extend Elasticsearch by providing additional ingest node capabilities. See https://www.elastic.co/guide/en/elasticsearch/plugins/master/ingest.html for more info.
//
// body: the simulate definition.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithID, WithPretty, WithSourceParam, WithVerbose, see the Option type in this package for more info.
func (i *Ingest) Simulate(body map[string]interface{}, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithErrorTrace":  struct{}{},
		"WithFilterPath":  struct{}{},
		"WithHuman":       struct{}{},
		"WithID":          struct{}{},
		"WithPretty":      struct{}{},
		"WithSourceParam": struct{}{},
		"WithVerbose":     struct{}{},
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
