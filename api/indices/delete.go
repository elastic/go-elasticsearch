// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"
)

// Delete - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-delete-index.html for more info.
//
// index: a comma-separated list of indices to delete; use "_all" or "*" string to delete all indices.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Delete(index []string, options ...*Option) (*DeleteResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
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
	resp, err := i.transport.Do(req)
	return &DeleteResponse{resp}, err
}

// DeleteResponse is the response for Delete
type DeleteResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
