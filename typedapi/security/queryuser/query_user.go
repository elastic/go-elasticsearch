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

// Find users with a query.
//
// Get information for users in a paginated manner.
// You can optionally filter the results with a query.
//
// NOTE: As opposed to the get user API, built-in users are excluded from the
// result.
// This API is only for native users.
package queryuser

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

type QueryUser struct {
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

// NewQueryUser type alias for index.
type NewQueryUser func() *QueryUser

// NewQueryUserFunc returns a new instance of QueryUser with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewQueryUserFunc(tp elastictransport.Interface) NewQueryUser {
	return func() *QueryUser {
		n := New(tp)

		return n
	}
}

// Find users with a query.
//
// Get information for users in a paginated manner.
// You can optionally filter the results with a query.
//
// NOTE: As opposed to the get user API, built-in users are excluded from the
// result.
// This API is only for native users.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-user.html
func New(tp elastictransport.Interface) *QueryUser {
	r := &QueryUser{
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
func (r *QueryUser) Raw(raw io.Reader) *QueryUser {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *QueryUser) Request(req *Request) *QueryUser {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *QueryUser) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for QueryUser: %w", err)
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
		path.WriteString("user")

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
func (r QueryUser) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.query_user")
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
		instrument.BeforeRequest(req, "security.query_user")
		if reader := instrument.RecordRequestBody(ctx, "security.query_user", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.query_user")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the QueryUser query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a queryuser.Response
func (r QueryUser) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.query_user")
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

// Header set a key, value pair in the QueryUser headers map.
func (r *QueryUser) Header(key, value string) *QueryUser {
	r.headers.Set(key, value)

	return r
}

// WithProfileUid Determines whether to retrieve the user profile UID, if it exists, for the
// users.
// API name: with_profile_uid
func (r *QueryUser) WithProfileUid(withprofileuid bool) *QueryUser {
	r.values.Set("with_profile_uid", strconv.FormatBool(withprofileuid))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *QueryUser) ErrorTrace(errortrace bool) *QueryUser {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *QueryUser) FilterPath(filterpaths ...string) *QueryUser {
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
func (r *QueryUser) Human(human bool) *QueryUser {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *QueryUser) Pretty(pretty bool) *QueryUser {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// From The starting document offset.
// It must not be negative.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` parameter.
// API name: from
func (r *QueryUser) From(from int) *QueryUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.From = &from

	return r
}

// Query A query to filter which users to return.
// If the query parameter is missing, it is equivalent to a `match_all` query.
// The query supports a subset of query types, including `match_all`, `bool`,
// `term`, `terms`, `match`,
// `ids`, `prefix`, `wildcard`, `exists`, `range`, and `simple_query_string`.
// You can query the following information associated with user: `username`,
// `roles`, `enabled`, `full_name`, and `email`.
// API name: query
func (r *QueryUser) Query(query *types.UserQueryContainer) *QueryUser {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}

// SearchAfter The search after definition
// API name: search_after
func (r *QueryUser) SearchAfter(sortresults ...types.FieldValue) *QueryUser {
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
func (r *QueryUser) Size(size int) *QueryUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Size = &size

	return r
}

// Sort The sort definition.
// Fields eligible for sorting are: `username`, `roles`, `enabled`.
// In addition, sort can also be applied to the `_doc` field to sort by index
// order.
// API name: sort
func (r *QueryUser) Sort(sorts ...types.SortCombinations) *QueryUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Sort = sorts

	return r
}
