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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Run a search with a search template.
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

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Run a search with a search template.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-template-api.html
func New(tp elastictransport.Interface) *SearchTemplate {
	r := &SearchTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for SearchTemplate: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
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

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r SearchTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "search_template")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "search_template")
		if reader := instrument.RecordRequestBody(ctx, "search_template", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "search_template")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SearchTemplate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a searchtemplate.Response
func (r SearchTemplate) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "search_template")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	r.TypedKeys(true)

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the SearchTemplate headers map.
func (r *SearchTemplate) Header(key, value string) *SearchTemplate {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of data streams, indices, and aliases to search.
// It supports wildcards (`*`).
// API Name: index
func (r *SearchTemplate) Index(index string) *SearchTemplate {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// For example, a request targeting `foo*,bar*` returns an error if an index
// starts with `foo` but no index starts with `bar`.
// API name: allow_no_indices
func (r *SearchTemplate) AllowNoIndices(allownoindices bool) *SearchTemplate {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// CcsMinimizeRoundtrips If `true`, network round-trips are minimized for cross-cluster search
// requests.
// API name: ccs_minimize_roundtrips
func (r *SearchTemplate) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *SearchTemplate {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// ExpandWildcards The type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *SearchTemplate) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *SearchTemplate {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If `true`, specified concrete, expanded, or aliased indices are not included
// in the response when throttled.
// API name: ignore_throttled
func (r *SearchTemplate) IgnoreThrottled(ignorethrottled bool) *SearchTemplate {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *SearchTemplate) IgnoreUnavailable(ignoreunavailable bool) *SearchTemplate {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Preference The node or shard the operation should be performed on.
// It is random by default.
// API name: preference
func (r *SearchTemplate) Preference(preference string) *SearchTemplate {
	r.values.Set("preference", preference)

	return r
}

// Routing A custom value used to route operations to a specific shard.
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

// RestTotalHitsAsInt If `true`, `hits.total` is rendered as an integer in the response.
// If `false`, it is rendered as an object.
// API name: rest_total_hits_as_int
func (r *SearchTemplate) RestTotalHitsAsInt(resttotalhitsasint bool) *SearchTemplate {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// TypedKeys If `true`, the response prefixes aggregation and suggester names with their
// respective types.
// API name: typed_keys
func (r *SearchTemplate) TypedKeys(typedkeys bool) *SearchTemplate {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SearchTemplate) ErrorTrace(errortrace bool) *SearchTemplate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SearchTemplate) FilterPath(filterpaths ...string) *SearchTemplate {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *SearchTemplate) Human(human bool) *SearchTemplate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *SearchTemplate) Pretty(pretty bool) *SearchTemplate {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Explain If `true`, returns detailed information about score calculation as part of
// each hit.
// If you specify both this and the `explain` query parameter, the API uses only
// the query parameter.
// API name: explain
func (r *SearchTemplate) Explain(explain bool) *SearchTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Explain = &explain

	return r
}

// Id The ID of the search template to use. If no `source` is specified,
// this parameter is required.
// API name: id
func (r *SearchTemplate) Id(id string) *SearchTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Id = &id

	return r
}

// Params Key-value pairs used to replace Mustache variables in the template.
// The key is the variable name.
// The value is the variable value.
// API name: params
func (r *SearchTemplate) Params(params map[string]json.RawMessage) *SearchTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Params = params

	return r
}

// Profile If `true`, the query execution is profiled.
// API name: profile
func (r *SearchTemplate) Profile(profile bool) *SearchTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Profile = &profile

	return r
}

// Source An inline search template. Supports the same parameters as the search API's
// request body. It also supports Mustache variables. If no `id` is specified,
// this
// parameter is required.
// API name: source
func (r *SearchTemplate) Source(source string) *SearchTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source = &source

	return r
}
