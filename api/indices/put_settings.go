// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// PutSettings - change specific index level settings in real time. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-update-settings.html for more info.
//
// body: the index settings to be updated.
//
// options: optional parameters. Supports the following functional options: WithIndexList, WithAllowNoIndices, WithExpandWildcards, WithFlatSettings, WithIgnoreUnavailable, WithMasterTimeout, WithPreserveExisting, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) PutSettings(body map[string]interface{}, options ...Option) (*PutSettingsResponse, error) {
	req := i.transport.NewRequest("PUT")
	methodOptions := supportedOptions["PutSettings"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &PutSettingsResponse{resp}, err
}

// PutSettingsResponse is the response for PutSettings.
type PutSettingsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *PutSettingsResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
