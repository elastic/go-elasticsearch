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

// Run an async ES|QL query.
// Asynchronously run an ES|QL (Elasticsearch query language) query, monitor its
// progress, and retrieve results when they become available.
//
// The API accepts the same parameters and request body as the synchronous query
// API, along with additional async related properties.
package asyncquery

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/esqlformat"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AsyncQuery struct {
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

// NewAsyncQuery type alias for index.
type NewAsyncQuery func() *AsyncQuery

// NewAsyncQueryFunc returns a new instance of AsyncQuery with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewAsyncQueryFunc(tp elastictransport.Interface) NewAsyncQuery {
	return func() *AsyncQuery {
		n := New(tp)

		return n
	}
}

// Run an async ES|QL query.
// Asynchronously run an ES|QL (Elasticsearch query language) query, monitor its
// progress, and retrieve results when they become available.
//
// The API accepts the same parameters and request body as the synchronous query
// API, along with additional async related properties.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-api.html
func New(tp elastictransport.Interface) *AsyncQuery {
	r := &AsyncQuery{
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
func (r *AsyncQuery) Raw(raw io.Reader) *AsyncQuery {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *AsyncQuery) Request(req *Request) *AsyncQuery {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *AsyncQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for AsyncQuery: %w", err)
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
		path.WriteString("_query")
		path.WriteString("/")
		path.WriteString("async")

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
func (r AsyncQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "esql.async_query")
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
		instrument.BeforeRequest(req, "esql.async_query")
		if reader := instrument.RecordRequestBody(ctx, "esql.async_query", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "esql.async_query")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the AsyncQuery query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a asyncquery.Response
func (r AsyncQuery) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.async_query")
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
		response, err = io.ReadAll(res.Body)
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

// Header set a key, value pair in the AsyncQuery headers map.
func (r *AsyncQuery) Header(key, value string) *AsyncQuery {
	r.headers.Set(key, value)

	return r
}

// AllowPartialResults If `true`, partial results will be returned if there are shard failures, but
// the query can continue to execute on other clusters and shards.
// If `false`, the query will fail if there are any failures.
//
// To override the default behavior, you can set the
// `esql.query.allow_partial_results` cluster setting to `false`.
// API name: allow_partial_results
func (r *AsyncQuery) AllowPartialResults(allowpartialresults bool) *AsyncQuery {
	r.values.Set("allow_partial_results", strconv.FormatBool(allowpartialresults))

	return r
}

// Delimiter The character to use between values within a CSV row.
// It is valid only for the CSV format.
// API name: delimiter
func (r *AsyncQuery) Delimiter(delimiter string) *AsyncQuery {
	r.values.Set("delimiter", delimiter)

	return r
}

// DropNullColumns Indicates whether columns that are entirely `null` will be removed from the
// `columns` and `values` portion of the results.
// If `true`, the response will include an extra section under the name
// `all_columns` which has the name of all the columns.
// API name: drop_null_columns
func (r *AsyncQuery) DropNullColumns(dropnullcolumns bool) *AsyncQuery {
	r.values.Set("drop_null_columns", strconv.FormatBool(dropnullcolumns))

	return r
}

// Format A short version of the Accept header, e.g. json, yaml.
//
// `csv`, `tsv`, and `txt` formats will return results in a tabular format,
// excluding other metadata fields from the response.
//
// For async requests, nothing will be returned if the async query doesn't
// finish within the timeout.
// The query ID and running status are available in the
// `X-Elasticsearch-Async-Id` and `X-Elasticsearch-Async-Is-Running` HTTP
// headers of the response, respectively.
// API name: format
func (r *AsyncQuery) Format(format esqlformat.EsqlFormat) *AsyncQuery {
	r.values.Set("format", format.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *AsyncQuery) ErrorTrace(errortrace bool) *AsyncQuery {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *AsyncQuery) FilterPath(filterpaths ...string) *AsyncQuery {
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
func (r *AsyncQuery) Human(human bool) *AsyncQuery {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *AsyncQuery) Pretty(pretty bool) *AsyncQuery {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Columnar By default, ES|QL returns results as rows. For example, FROM returns each
// individual document as one row. For the JSON, YAML, CBOR and smile formats,
// ES|QL can return the results in a columnar fashion where one row represents
// all the values of a certain column in the results.
// API name: columnar
func (r *AsyncQuery) Columnar(columnar bool) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Columnar = &columnar

	return r
}

// Filter Specify a Query DSL query in the filter parameter to filter the set of
// documents that an ES|QL query runs on.
// API name: filter
func (r *AsyncQuery) Filter(filter *types.Query) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Filter = filter

	return r
}

// IncludeCcsMetadata When set to `true` and performing a cross-cluster query, the response will
// include an extra `_clusters`
// object with information about the clusters that participated in the search
// along with info such as shards
// count.
// API name: include_ccs_metadata
func (r *AsyncQuery) IncludeCcsMetadata(includeccsmetadata bool) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IncludeCcsMetadata = &includeccsmetadata

	return r
}

// KeepAlive The period for which the query and its results are stored in the cluster.
// The default period is five days.
// When this period expires, the query and its results are deleted, even if the
// query is still ongoing.
// If the `keep_on_completion` parameter is false, Elasticsearch only stores
// async queries that do not complete within the period set by the
// `wait_for_completion_timeout` parameter, regardless of this value.
// API name: keep_alive
func (r *AsyncQuery) KeepAlive(duration types.Duration) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.KeepAlive = duration

	return r
}

// KeepOnCompletion Indicates whether the query and its results are stored in the cluster.
// If false, the query and its results are stored in the cluster only if the
// request does not complete during the period set by the
// `wait_for_completion_timeout` parameter.
// API name: keep_on_completion
func (r *AsyncQuery) KeepOnCompletion(keeponcompletion bool) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.KeepOnCompletion = &keeponcompletion

	return r
}

// API name: locale
func (r *AsyncQuery) Locale(locale string) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Locale = &locale

	return r
}

// Params To avoid any attempts of hacking or code injection, extract the values in a
// separate list of parameters. Use question mark placeholders (?) in the query
// string for each of the parameters.
// API name: params
func (r *AsyncQuery) Params(params ...types.FieldValue) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Params = params

	return r
}

// Profile If provided and `true` the response will include an extra `profile` object
// with information on how the query was executed. This information is for human
// debugging
// and its format can change at any time but it can give some insight into the
// performance
// of each part of the query.
// API name: profile
func (r *AsyncQuery) Profile(profile bool) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Profile = &profile

	return r
}

// Query The ES|QL query API accepts an ES|QL query string in the query parameter,
// runs it, and returns the results.
// API name: query
func (r *AsyncQuery) Query(query string) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}

// Tables Tables to use with the LOOKUP operation. The top level key is the table
// name and the next level key is the column name.
// API name: tables
func (r *AsyncQuery) Tables(tables map[string]map[string]types.TableValuesContainer) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Tables = tables

	return r
}

// WaitForCompletionTimeout The period to wait for the request to finish.
// By default, the request waits for 1 second for the query results.
// If the query completes during this period, results are returned
// Otherwise, a query ID is returned that can later be used to retrieve the
// results.
// API name: wait_for_completion_timeout
func (r *AsyncQuery) WaitForCompletionTimeout(duration types.Duration) *AsyncQuery {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.WaitForCompletionTimeout = duration

	return r
}
