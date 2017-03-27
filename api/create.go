// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Create - the index API adds or updates a typed JSON document in a specific index, making it searchable. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-index_.html for more info.
//
// index: the name of the index.
//
// documentType: the type of the document.
//
// id: document ID.
//
// body: the document.
//
// options: optional parameters. Supports the following functional options: WithParent, WithPipeline, WithRefresh, WithRouting, WithTimeout, WithTimestamp, WithTTL, WithVersion, WithVersionType, WithWaitForActiveShards, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Create(index string, documentType string, id string, body map[string]interface{}, options ...*Option) (*CreateResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "PUT",
	}
	methodOptions := supportedOptions["Create"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &CreateResponse{resp}, err
}

// CreateResponse is the response for Create
type CreateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
