// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Percolate - for indices created on or after version 5.0.0-alpha1 the percolator automatically indexes the query terms with the percolator queries. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-percolate.html for more info.
//
// index: the index of the document being percolated.
//
// documentType: the type of the document being percolated.
//
// options: optional parameters. Supports the following functional options: WithID, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithPercolateFormat, WithPercolateIndex, WithPercolatePreference, WithPercolateRouting, WithPercolateType, WithPreference, WithRouting, WithVersion, WithVersionType, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Percolate(index string, documentType string, options ...*Option) (*PercolateResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Percolate"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &PercolateResponse{resp}, err
}

// PercolateResponse is the response for Percolate
type PercolateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
