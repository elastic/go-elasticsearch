// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// Msearch - the multi search API allows to execute several search requests within the same API. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-multi-search.html for more info.
//
// body: the request definitions (metadata-search request definition pairs), separated by newlines.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithMaxConcurrentSearches, WithSearchType, WithTypedKeys, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Msearch(body []interface{}, options ...Option) (*MsearchResponse, error) {
	req := a.transport.NewRequest("GET")
	methodOptions := supportedOptions["Msearch"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &MsearchResponse{resp}, err
}

// MsearchResponse is the response for Msearch.
type MsearchResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *MsearchResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
