// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Search allows you to execute a search query and get back search hits that match the query. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-search.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithType, WithSource, WithSourceExclude, WithSourceInclude, WithAllowNoIndices, WithAnalyzeWildcard, WithAnalyzer, WithBatchedReduceSize, WithDefaultOperator, WithDf, WithDocvalueFields, WithExpandWildcards, WithExplain, WithFielddataFields, WithFrom, WithIgnoreUnavailable, WithLenient, WithPreference, WithQ, WithRequestCache, WithRouting, WithScroll, WithSearchType, WithSize, WithSort, WithStats, WithStoredFields, WithSuggestField, WithSuggestMode, WithSuggestSize, WithSuggestText, WithTerminateAfter, WithTimeout, WithTrackScores, WithTypedKeys, WithVersion, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Search(options ...*Option) (*SearchResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Search"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &SearchResponse{resp}, err
}

// SearchResponse is the response for Search
type SearchResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *SearchResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
