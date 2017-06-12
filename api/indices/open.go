// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// Open - the open and close index APIs allow to close an index, and later on opening it. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-open-close.html for more info.
//
// index: a comma separated list of indices to open.
//
// options: optional parameters. Supports the following functional options: WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithMasterTimeout, WithTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Open(index []string, options ...Option) (*OpenResponse, error) {
	req := i.transport.NewRequest("POST")
	methodOptions := supportedOptions["Open"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &OpenResponse{resp}, err
}

// OpenResponse is the response for Open.
type OpenResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *OpenResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
