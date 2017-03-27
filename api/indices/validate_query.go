// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// ValidateQuery - the validate API allows a user to validate a potentially expensive query without executing it. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-validate.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithAllShards, WithAllowNoIndices, WithAnalyzeWildcard, WithAnalyzer, WithDefaultOperator, WithDf, WithExpandWildcards, WithExplain, WithIgnoreUnavailable, WithLenient, WithOperationThreading, WithQ, WithRewrite, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) ValidateQuery(options ...*Option) (*ValidateQueryResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["ValidateQuery"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &ValidateQueryResponse{resp}, err
}

// ValidateQueryResponse is the response for ValidateQuery
type ValidateQueryResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *ValidateQueryResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
