// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// GetTemplate - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-template.html for more info.
//
// id: template ID.
//
// options: optional parameters. Supports the following functional options: WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (a *API) GetTemplate(id string, options ...Option) (*GetTemplateResponse, error) {
	req := a.transport.NewRequest("GET")
	methodOptions := supportedOptions["GetTemplate"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &GetTemplateResponse{resp}, err
}

// GetTemplateResponse is the response for GetTemplate.
type GetTemplateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetTemplateResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
