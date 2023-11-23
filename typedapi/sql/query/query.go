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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Executes a SQL request
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
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Query struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
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

// Executes a SQL request
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/sql-search-api.html
func New(tp elastictransport.Interface) *Query {
	r := &Query{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Query: %w", err)
		}

		r.buf.Write(data)

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
func (r Query) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Query query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a query.Response
func (r Query) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

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

// Header set a key, value pair in the Query headers map.
func (r *Query) Header(key, value string) *Query {
	r.headers.Set(key, value)

	return r
}

// Format Format for the response.
// API name: format
func (r *Query) Format(format string) *Query {
	r.values.Set("format", format)

	return r
}

// Catalog Default catalog (cluster) for queries. If unspecified, the queries execute on
// the data in the local cluster only.
// API name: catalog
func (r *Query) Catalog(catalog string) *Query {

	r.req.Catalog = &catalog

	return r
}

// Columnar If true, the results in a columnar fashion: one row represents all the values
// of a certain column from the current page of results.
// API name: columnar
func (r *Query) Columnar(columnar bool) *Query {
	r.req.Columnar = &columnar

	return r
}

// Cursor Cursor used to retrieve a set of paginated results.
// If you specify a cursor, the API only uses the `columnar` and `time_zone`
// request body parameters.
// It ignores other request body parameters.
// API name: cursor
func (r *Query) Cursor(cursor string) *Query {

	r.req.Cursor = &cursor

	return r
}

// FetchSize The maximum number of rows (or entries) to return in one response
// API name: fetch_size
func (r *Query) FetchSize(fetchsize int) *Query {
	r.req.FetchSize = &fetchsize

	return r
}

// FieldMultiValueLeniency Throw an exception when encountering multiple values for a field (default) or
// be lenient and return the first value from the list (without any guarantees
// of what that will be - typically the first in natural ascending order).
// API name: field_multi_value_leniency
func (r *Query) FieldMultiValueLeniency(fieldmultivalueleniency bool) *Query {
	r.req.FieldMultiValueLeniency = &fieldmultivalueleniency

	return r
}

// Filter Elasticsearch query DSL for additional filtering.
// API name: filter
func (r *Query) Filter(filter *types.Query) *Query {

	r.req.Filter = filter

	return r
}

// IndexUsingFrozen If true, the search can run on frozen indices. Defaults to false.
// API name: index_using_frozen
func (r *Query) IndexUsingFrozen(indexusingfrozen bool) *Query {
	r.req.IndexUsingFrozen = &indexusingfrozen

	return r
}

// KeepAlive Retention period for an async or saved synchronous search.
// API name: keep_alive
func (r *Query) KeepAlive(duration types.Duration) *Query {
	r.req.KeepAlive = duration

	return r
}

// KeepOnCompletion If true, Elasticsearch stores synchronous searches if you also specify the
// wait_for_completion_timeout parameter. If false, Elasticsearch only stores
// async searches that don’t finish before the wait_for_completion_timeout.
// API name: keep_on_completion
func (r *Query) KeepOnCompletion(keeponcompletion bool) *Query {
	r.req.KeepOnCompletion = &keeponcompletion

	return r
}

// PageTimeout The timeout before a pagination request fails.
// API name: page_timeout
func (r *Query) PageTimeout(duration types.Duration) *Query {
	r.req.PageTimeout = duration

	return r
}

// Params Values for parameters in the query.
// API name: params
func (r *Query) Params(params map[string]json.RawMessage) *Query {

	r.req.Params = params

	return r
}

// Query SQL query to run.
// API name: query
func (r *Query) Query(query string) *Query {

	r.req.Query = &query

	return r
}

// RequestTimeout The timeout before the request fails.
// API name: request_timeout
func (r *Query) RequestTimeout(duration types.Duration) *Query {
	r.req.RequestTimeout = duration

	return r
}

// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
// precedence over mapped fields with the same name.
// API name: runtime_mappings
func (r *Query) RuntimeMappings(runtimefields types.RuntimeFields) *Query {
	r.req.RuntimeMappings = runtimefields

	return r
}

// TimeZone ISO-8601 time zone ID for the search.
// API name: time_zone
func (r *Query) TimeZone(timezone string) *Query {
	r.req.TimeZone = &timezone

	return r
}

// WaitForCompletionTimeout Period to wait for complete results. Defaults to no timeout, meaning the
// request waits for complete search results. If the search doesn’t finish
// within this period, the search becomes async.
// API name: wait_for_completion_timeout
func (r *Query) WaitForCompletionTimeout(duration types.Duration) *Query {
	r.req.WaitForCompletionTimeout = duration

	return r
}
