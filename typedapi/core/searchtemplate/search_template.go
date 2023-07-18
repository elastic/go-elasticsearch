// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/33e8a1c9cad22a5946ac735c4fba31af2da2cec2

// Allows to use the Mustache language to pre-render a search definition.
package searchtemplate

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SearchTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
}

// NewSearchTemplate type alias for index.
type NewSearchTemplate func() *SearchTemplate

// NewSearchTemplateFunc returns a new instance of SearchTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSearchTemplateFunc(tp elastictransport.Interface) NewSearchTemplate {
	return func() *SearchTemplate {
		n := New(tp)

		return n
	}
}

// Allows to use the Mustache language to pre-render a search definition.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template.html
func New(tp elastictransport.Interface) *SearchTemplate {
	r := &SearchTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *SearchTemplate) Raw(raw io.Reader) *SearchTemplate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SearchTemplate) Request(req *Request) *SearchTemplate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SearchTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for SearchTemplate: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_search")
		path.WriteString("/")
		path.WriteString("template")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_search")
		path.WriteString("/")
		path.WriteString("template")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r SearchTemplate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SearchTemplate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a searchtemplate.Response
func (r SearchTemplate) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	r.TypedKeys(true)

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the SearchTemplate headers map.
func (r *SearchTemplate) Header(key, value string) *SearchTemplate {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices,
// and aliases to search. Supports wildcards (*).
// API Name: index
func (r *SearchTemplate) Index(index string) *SearchTemplate {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *SearchTemplate) AllowNoIndices(allownoindices bool) *SearchTemplate {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// CcsMinimizeRoundtrips Indicates whether network round-trips should be minimized as part of
// cross-cluster search requests execution
// API name: ccs_minimize_roundtrips
func (r *SearchTemplate) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *SearchTemplate {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *SearchTemplate) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *SearchTemplate {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled Whether specified concrete, expanded or aliased indices should be ignored
// when throttled
// API name: ignore_throttled
func (r *SearchTemplate) IgnoreThrottled(ignorethrottled bool) *SearchTemplate {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *SearchTemplate) IgnoreUnavailable(ignoreunavailable bool) *SearchTemplate {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *SearchTemplate) Preference(preference string) *SearchTemplate {
	r.values.Set("preference", preference)

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *SearchTemplate) Routing(routing string) *SearchTemplate {
	r.values.Set("routing", routing)

	return r
}

// Scroll Specifies how long a consistent view of the index
// should be maintained for scrolled search.
// API name: scroll
func (r *SearchTemplate) Scroll(duration string) *SearchTemplate {
	r.values.Set("scroll", duration)

	return r
}

// SearchType The type of the search operation.
// API name: search_type
func (r *SearchTemplate) SearchType(searchtype searchtype.SearchType) *SearchTemplate {
	r.values.Set("search_type", searchtype.String())

	return r
}

// RestTotalHitsAsInt If true, hits.total are rendered as an integer in the response.
// API name: rest_total_hits_as_int
func (r *SearchTemplate) RestTotalHitsAsInt(resttotalhitsasint bool) *SearchTemplate {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// TypedKeys Specify whether aggregation and suggester names should be prefixed by their
// respective types in the response
// API name: typed_keys
func (r *SearchTemplate) TypedKeys(typedkeys bool) *SearchTemplate {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// API name: explain
func (r *SearchTemplate) Explain(explain bool) *SearchTemplate {
	r.req.Explain = &explain

	return r
}

// Id ID of the search template to use. If no source is specified,
// this parameter is required.
// API name: id
func (r *SearchTemplate) Id(id string) *SearchTemplate {
	r.req.Id = &id

	return r
}

// API name: params
func (r *SearchTemplate) Params(params map[string]json.RawMessage) *SearchTemplate {

	r.req.Params = params

	return r
}

// API name: profile
func (r *SearchTemplate) Profile(profile bool) *SearchTemplate {
	r.req.Profile = &profile

	return r
}

// Source An inline search template. Supports the same parameters as the search API's
// request body. Also supports Mustache variables. If no id is specified, this
// parameter is required.
// API name: source
func (r *SearchTemplate) Source(source string) *SearchTemplate {

	r.req.Source = &source

	return r
}
