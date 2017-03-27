// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Shards - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-shards.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithFormat, WithH, WithHelp, WithLocal, WithMasterTimeout, WithS, WithV, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (c *Cat) Shards(options ...*Option) (*ShardsResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: c.transport.URL.Scheme,
			Host:   c.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Shards"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := c.transport.Do(req)
	return &ShardsResponse{resp}, err
}

// ShardsResponse is the response for Shards
type ShardsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *ShardsResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
