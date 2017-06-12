// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// Segments - provide low level segments information that a Lucene index (shard level) is built with. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-segments.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndexList, WithAllowNoIndices, WithExpandWildcards, WithIgnoreUnavailable, WithVerbose, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) Segments(options ...Option) (*SegmentsResponse, error) {
	req := i.transport.NewRequest("GET")
	methodOptions := supportedOptions["Segments"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &SegmentsResponse{resp}, err
}

// SegmentsResponse is the response for Segments.
type SegmentsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *SegmentsResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
