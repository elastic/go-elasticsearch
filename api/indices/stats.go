// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Stats - indices level stats provide statistics on different operations happening on an index. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-stats.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithMetric, WithCompletionFields, WithFielddataFields, WithFields, WithGroups, WithIncludeSegmentFileSizes, WithLevel, WithTypes, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Stats(options ...*Option) (*StatsResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Stats"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &StatsResponse{resp}, err
}

// StatsResponse is the response for Stats
type StatsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *StatsResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
