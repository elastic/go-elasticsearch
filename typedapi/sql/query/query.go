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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Get SQL search results.
// Run an SQL request.
package query

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sqlformat"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Query struct {
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

// NewQuery type alias for index.
type NewQuery func() *Query

// NewQueryFunc returns a new instance of Query with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewQueryFunc(tp elastictransport.Interface) NewQuery {
	return func() *Query {
		n := New(tp)

		return n
	}
}

// Get SQL search results.
// Run an SQL request.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-sql-query
func New(tp elastictransport.Interface) *Query {
	r := &Query{
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
func (r *Query) Raw(raw io.Reader) *Query {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Query) Request(req *Request) *Query {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Query) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Query: %w", err)
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
		path.WriteString("_sql")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Query) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "sql.query")
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
		instrument.BeforeRequest(req, "sql.query")
		if reader := instrument.RecordRequestBody(ctx, "sql.query", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "sql.query")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Query query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a query.Response
func (r Query) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "sql.query")
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

// Header set a key, value pair in the Query headers map.
func (r *Query) Header(key, value string) *Query {
	r.headers.Set(key, value)

	return r
}

// Format The format for the response.
// You can also specify a format using the `Accept` HTTP header.
// If you specify both this parameter and the `Accept` HTTP header, this
// parameter takes precedence.
// API name: format
func (r *Query) Format(format sqlformat.SqlFormat) *Query {
	r.values.Set("format", format.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Query) ErrorTrace(errortrace bool) *Query {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Query) FilterPath(filterpaths ...string) *Query {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *Query) Human(human bool) *Query {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Query) Pretty(pretty bool) *Query {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// If `true`, the response has partial results when there are shard request
// timeouts or shard failures.
// If `false`, the API returns an error with no partial results.
// API name: allow_partial_search_results
func (r *Query) AllowPartialSearchResults(allowpartialsearchresults bool) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AllowPartialSearchResults = &allowpartialsearchresults

	return r
}

// The default catalog (cluster) for queries.
// If unspecified, the queries execute on the data in the local cluster only.
// API name: catalog
func (r *Query) Catalog(catalog string) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Catalog = &catalog

	return r
}

// If `true`, the results are in a columnar fashion: one row represents all the
// values of a certain column from the current page of results.
// The API supports this parameter only for CBOR, JSON, SMILE, and YAML
// responses.
// API name: columnar
func (r *Query) Columnar(columnar bool) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Columnar = &columnar

	return r
}

// The cursor used to retrieve a set of paginated results.
// If you specify a cursor, the API only uses the `columnar` and `time_zone`
// request body parameters.
// It ignores other request body parameters.
// API name: cursor
func (r *Query) Cursor(cursor string) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Cursor = &cursor

	return r
}

// The maximum number of rows (or entries) to return in one response.
// API name: fetch_size
func (r *Query) FetchSize(fetchsize int) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.FetchSize = &fetchsize

	return r
}

// If `false`, the API returns an exception when encountering multiple values
// for a field.
// If `true`, the API is lenient and returns the first value from the array with
// no guarantee of consistent results.
// API name: field_multi_value_leniency
func (r *Query) FieldMultiValueLeniency(fieldmultivalueleniency bool) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.FieldMultiValueLeniency = &fieldmultivalueleniency

	return r
}

// The Elasticsearch query DSL for additional filtering.
// API name: filter
func (r *Query) Filter(filter types.QueryVariant) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Filter = filter.QueryCaster()

	return r
}

// If `true`, the search can run on frozen indices.
// API name: index_using_frozen
func (r *Query) IndexUsingFrozen(indexusingfrozen bool) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexUsingFrozen = &indexusingfrozen

	return r
}

// The retention period for an async or saved synchronous search.
// API name: keep_alive
func (r *Query) KeepAlive(duration types.DurationVariant) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.KeepAlive = *duration.DurationCaster()

	return r
}

// If `true`, Elasticsearch stores synchronous searches if you also specify the
// `wait_for_completion_timeout` parameter.
// If `false`, Elasticsearch only stores async searches that don't finish before
// the `wait_for_completion_timeout`.
// API name: keep_on_completion
func (r *Query) KeepOnCompletion(keeponcompletion bool) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.KeepOnCompletion = &keeponcompletion

	return r
}

// The minimum retention period for the scroll cursor.
// After this time period, a pagination request might fail because the scroll
// cursor is no longer available.
// Subsequent scroll requests prolong the lifetime of the scroll cursor by the
// duration of `page_timeout` in the scroll request.
// API name: page_timeout
func (r *Query) PageTimeout(duration types.DurationVariant) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PageTimeout = *duration.DurationCaster()

	return r
}

// The values for parameters in the query.
// API name: params
func (r *Query) Params(params ...json.RawMessage) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range params {

		r.req.Params = append(r.req.Params, v)

	}
	return r
}

// The SQL query to run.
// API name: query
func (r *Query) Query(query string) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = &query

	return r
}

// The timeout before the request fails.
// API name: request_timeout
func (r *Query) RequestTimeout(duration types.DurationVariant) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RequestTimeout = *duration.DurationCaster()

	return r
}

// One or more runtime fields for the search request.
// These fields take precedence over mapped fields with the same name.
// API name: runtime_mappings
func (r *Query) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return r
}

// The ISO-8601 time zone ID for the search.
// API name: time_zone
func (r *Query) TimeZone(timezone string) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TimeZone = &timezone

	return r
}

// The period to wait for complete results.
// It defaults to no timeout, meaning the request waits for complete search
// results.
// If the search doesn't finish within this period, the search becomes async.
//
// To save a synchronous search, you must specify this parameter and the
// `keep_on_completion` parameter.
// API name: wait_for_completion_timeout
func (r *Query) WaitForCompletionTimeout(duration types.DurationVariant) *Query {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.WaitForCompletionTimeout = *duration.DurationCaster()

	return r
}
