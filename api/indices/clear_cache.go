// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// ClearCache - the clear cache API allows to clear either all caches or specific cached associated with one or more indices. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-clearcache.html for more info.
//
// options: optional parameters. Supports the following functional options: WithIndex, WithAllowNoIndices, WithExpandWildcards, WithFieldData, WithFielddata, WithFields, WithIgnoreUnavailable, WithIndexParam, WithQuery, WithRecycler, WithRequest, WithRequestCache, WithErrorTrace, WithFilterPath, WithHuman, WithIgnore, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (i *Indices) ClearCache(options ...Option) (*ClearCacheResponse, error) {
	req := i.transport.NewRequest("POST")
	methodOptions := supportedOptions["ClearCache"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := i.transport.Do(req)
	return &ClearCacheResponse{resp}, err
}

// ClearCacheResponse is the response for ClearCache.
type ClearCacheResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *ClearCacheResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
