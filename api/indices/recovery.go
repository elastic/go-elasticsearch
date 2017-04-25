// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// Recovery - the indices recovery API provides insight into on-going index shard recoveries. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-recovery.html for more info.
//
// options: optional parameters. Supports the following functional options: WithActiveOnly, WithDetailed, WithErrorTrace, WithFilterPath, WithHuman, WithIndex, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Recovery(options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithActiveOnly":  struct{}{},
		"WithDetailed":    struct{}{},
		"WithErrorTrace":  struct{}{},
		"WithFilterPath":  struct{}{},
		"WithHuman":       struct{}{},
		"WithIndex":       struct{}{},
		"WithPretty":      struct{}{},
		"WithSourceParam": struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.Scheme,
			Host:   i.transport.Host,
		},
		Method: "GET",
	}
	for _, option := range options {
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return i.transport.Do(req)
}
