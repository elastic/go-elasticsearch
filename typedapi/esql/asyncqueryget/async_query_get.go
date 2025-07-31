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

// Get async ES|QL query results.
// Get the current status and available results or stored results for an ES|QL
// asynchronous query.
// If the Elasticsearch security features are enabled, only the user who first
// submitted the ES|QL query can retrieve the results using this API.
package asyncqueryget

import (
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

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AsyncQueryGet struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewAsyncQueryGet type alias for index.
type NewAsyncQueryGet func(id string) *AsyncQueryGet

// NewAsyncQueryGetFunc returns a new instance of AsyncQueryGet with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewAsyncQueryGetFunc(tp elastictransport.Interface) NewAsyncQueryGet {
	return func(id string) *AsyncQueryGet {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Get async ES|QL query results.
// Get the current status and available results or stored results for an ES|QL
// asynchronous query.
// If the Elasticsearch security features are enabled, only the user who first
// submitted the ES|QL query can retrieve the results using this API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-async-query-get-api.html
func New(tp elastictransport.Interface) *AsyncQueryGet {
	r := &AsyncQueryGet{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *AsyncQueryGet) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_query")
		path.WriteString("/")
		path.WriteString("async")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.id)
		}
		path.WriteString(r.id)

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r AsyncQueryGet) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "esql.async_query_get")
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
		instrument.BeforeRequest(req, "esql.async_query_get")
		if reader := instrument.RecordRequestBody(ctx, "esql.async_query_get", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "esql.async_query_get")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the AsyncQueryGet query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a asyncqueryget.Response
func (r AsyncQueryGet) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.async_query_get")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r AsyncQueryGet) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.async_query_get")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the AsyncQueryGet query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the AsyncQueryGet headers map.
func (r *AsyncQueryGet) Header(key, value string) *AsyncQueryGet {
	r.headers.Set(key, value)

	return r
}

// Id The unique identifier of the query.
// A query ID is provided in the ES|QL async query API response for a query that
// does not complete in the designated time.
// A query ID is also provided when the request was submitted with the
// `keep_on_completion` parameter set to `true`.
// API Name: id
func (r *AsyncQueryGet) _id(id string) *AsyncQueryGet {
	r.paramSet |= idMask
	r.id = id

	return r
}

// DropNullColumns Indicates whether columns that are entirely `null` will be removed from the
// `columns` and `values` portion of the results.
// If `true`, the response will include an extra section under the name
// `all_columns` which has the name of all the columns.
// API name: drop_null_columns
func (r *AsyncQueryGet) DropNullColumns(dropnullcolumns bool) *AsyncQueryGet {
	r.values.Set("drop_null_columns", strconv.FormatBool(dropnullcolumns))

	return r
}

// Format A short version of the Accept header, for example `json` or `yaml`.
// API name: format
func (r *AsyncQueryGet) Format(format esqlformat.EsqlFormat) *AsyncQueryGet {
	r.values.Set("format", format.String())

	return r
}

// KeepAlive The period for which the query and its results are stored in the cluster.
// When this period expires, the query and its results are deleted, even if the
// query is still ongoing.
// API name: keep_alive
func (r *AsyncQueryGet) KeepAlive(duration string) *AsyncQueryGet {
	r.values.Set("keep_alive", duration)

	return r
}

// WaitForCompletionTimeout The period to wait for the request to finish.
// By default, the request waits for complete query results.
// If the request completes during the period specified in this parameter,
// complete query results are returned.
// Otherwise, the response returns an `is_running` value of `true` and no
// results.
// API name: wait_for_completion_timeout
func (r *AsyncQueryGet) WaitForCompletionTimeout(duration string) *AsyncQueryGet {
	r.values.Set("wait_for_completion_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *AsyncQueryGet) ErrorTrace(errortrace bool) *AsyncQueryGet {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *AsyncQueryGet) FilterPath(filterpaths ...string) *AsyncQueryGet {
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
func (r *AsyncQueryGet) Human(human bool) *AsyncQueryGet {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *AsyncQueryGet) Pretty(pretty bool) *AsyncQueryGet {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
