// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// Snapshots - see http://www.elastic.co/guide/en/elasticsearch/reference/master/cat-snapshots.html for more info.
//
// repository: name of repository from which to fetch the snapshot information.
//
// options: optional parameters. Supports the following functional options: WithFormat, WithH, WithHelp, WithIgnoreUnavailable, WithMasterTimeout, WithS, WithV, see the Option type in this package for more info.
func (c *Cat) Snapshots(repository []string, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithFormat":            struct{}{},
		"WithH":                 struct{}{},
		"WithHelp":              struct{}{},
		"WithIgnoreUnavailable": struct{}{},
		"WithMasterTimeout":     struct{}{},
		"WithS":                 struct{}{},
		"WithV":                 struct{}{},
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
