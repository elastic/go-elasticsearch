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

// Find roles with a query.
//
// Get roles in a paginated manner.
// The role management APIs are generally the preferred way to manage roles,
// rather than using file-based role management.
// The query roles API does not retrieve roles that are defined in roles files,
// nor built-in ones.
// You can optionally filter the results with a query.
// Also, the results can be paginated and sorted.
package queryrole

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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type QueryRole struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewQueryRole type alias for index.
type NewQueryRole func() *QueryRole

// NewQueryRoleFunc returns a new instance of QueryRole with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewQueryRoleFunc(tp elastictransport.Interface) NewQueryRole {
	return func() *QueryRole {
		n := New(tp)

		return n
	}
}

// Find roles with a query.
//
// Get roles in a paginated manner.
// The role management APIs are generally the preferred way to manage roles,
// rather than using file-based role management.
// The query roles API does not retrieve roles that are defined in roles files,
// nor built-in ones.
// You can optionally filter the results with a query.
// Also, the results can be paginated and sorted.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-role.html
func New(tp elastictransport.Interface) *QueryRole {
	r := &QueryRole{
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
func (r *QueryRole) Raw(raw io.Reader) *QueryRole {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *QueryRole) Request(req *Request) *QueryRole {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *QueryRole) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for QueryRole: %w", err)
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
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("_query")
		path.WriteString("/")
		path.WriteString("role")

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
func (r QueryRole) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.query_role")
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
		instrument.BeforeRequest(req, "security.query_role")
		if reader := instrument.RecordRequestBody(ctx, "security.query_role", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.query_role")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the QueryRole query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a queryrole.Response
func (r QueryRole) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.query_role")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

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

// Header set a key, value pair in the QueryRole headers map.
func (r *QueryRole) Header(key, value string) *QueryRole {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *QueryRole) ErrorTrace(errortrace bool) *QueryRole {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *QueryRole) FilterPath(filterpaths ...string) *QueryRole {
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
func (r *QueryRole) Human(human bool) *QueryRole {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *QueryRole) Pretty(pretty bool) *QueryRole {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// From The starting document offset.
// It must not be negative.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` parameter.
// API name: from
func (r *QueryRole) From(from int) *QueryRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.From = &from

	return r
}

// Query A query to filter which roles to return.
// If the query parameter is missing, it is equivalent to a `match_all` query.
// The query supports a subset of query types, including `match_all`, `bool`,
// `term`, `terms`, `match`,
// `ids`, `prefix`, `wildcard`, `exists`, `range`, and `simple_query_string`.
// You can query the following information associated with roles: `name`,
// `description`, `metadata`,
// `applications.application`, `applications.privileges`, and
// `applications.resources`.
// API name: query
func (r *QueryRole) Query(query *types.RoleQueryContainer) *QueryRole {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}

// SearchAfter The search after definition.
// API name: search_after
func (r *QueryRole) SearchAfter(sortresults ...types.FieldValue) *QueryRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.SearchAfter = sortresults

	return r
}

// Size The number of hits to return.
// It must not be negative.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` parameter.
// API name: size
func (r *QueryRole) Size(size int) *QueryRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Size = &size

	return r
}

// Sort The sort definition.
// You can sort on `username`, `roles`, or `enabled`.
// In addition, sort can also be applied to the `_doc` field to sort by index
// order.
// API name: sort
func (r *QueryRole) Sort(sorts ...types.SortCombinations) *QueryRole {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Sort = sorts

	return r
}
