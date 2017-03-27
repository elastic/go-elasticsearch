// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Count allows to easily execute a query and get the number of matches for that query. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-count.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithAllowNoIndices, WithAnalyzeWildcard, WithAnalyzer, WithDefaultOperator, WithDf, WithExpandWildcards, WithIgnoreUnavailable, WithLenient, WithMinScore, WithPreference, WithQ, WithRouting, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Count(options ...*Option) (*CountResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "POST",
	}
	methodOptions := supportedOptions["Count"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &CountResponse{resp}, err
}

// CountResponse is the response for Count
type CountResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *CountResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
