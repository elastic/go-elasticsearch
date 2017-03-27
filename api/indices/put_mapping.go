// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// PutMapping - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-put-mapping.html for more info.
//
// index: a comma-separated list of index names the mapping should be added to (supports wildcards); use "_all" or omit to add the mapping on all indices.
//
// documentType: the name of the document type.
//
// body: the mapping definition.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithMasterTimeout, WithTimeout, WithUpdateAllTypes, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) PutMapping(index []string, documentType string, body map[string]interface{}, options ...*Option) (*PutMappingResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: i.transport.URL.Scheme,
			Host:   i.transport.URL.Host,
		},
		Method: "PUT",
	}
	methodOptions := supportedOptions["PutMapping"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &PutMappingResponse{resp}, err
}

// PutMappingResponse is the response for PutMapping
type PutMappingResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

func (r *PutMappingResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
