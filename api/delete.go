// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Delete - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-delete.html for more info.
//
// index: the name of the index.
//
// documentType: the type of the document.
//
// id: the document ID.
//
// options: optional parameters. Supports the following functional options: WithParent, WithRefresh, WithRouting, WithTimeout, WithVersion, WithVersionType, WithWaitForActiveShards, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) Delete(index string, documentType string, id string, options ...*Option) (*DeleteResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "DELETE",
	}
	methodOptions := supportedOptions["Delete"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &DeleteResponse{resp}, err
}

// DeleteResponse is the response for Delete
type DeleteResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *DeleteResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
