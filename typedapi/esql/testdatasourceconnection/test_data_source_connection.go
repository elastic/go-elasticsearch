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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Test an ES|QL data source connection.
//
// Tests whether the supplied data source configuration can establish a live
// connection. The data source does not need to exist in cluster state: this
// endpoint is intended for validating a new configuration before saving it. The
// request body accepts the same `type` and `settings` fields as the create or
// update data source API.
package testdatasourceconnection

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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type TestDataSourceConnection struct {
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

// NewTestDataSourceConnection type alias for index.
type NewTestDataSourceConnection func() *TestDataSourceConnection

// NewTestDataSourceConnectionFunc returns a new instance of TestDataSourceConnection with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewTestDataSourceConnectionFunc(tp elastictransport.Interface) NewTestDataSourceConnection {
	return func() *TestDataSourceConnection {
		n := New(tp)

		return n
	}
}

// Test an ES|QL data source connection.
//
// Tests whether the supplied data source configuration can establish a live
// connection. The data source does not need to exist in cluster state: this
// endpoint is intended for validating a new configuration before saving it. The
// request body accepts the same `type` and `settings` fields as the create or
// update data source API.
//
// [Elasticsearch] https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-esql-data-source-test-connection
//
// [Serverless] https://www.elastic.co/docs/api/doc/elasticsearch-serverless/operation/operation-esql-data-source-test-connection
func New(tp elastictransport.Interface) *TestDataSourceConnection {
	r := &TestDataSourceConnection{
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
func (r *TestDataSourceConnection) Raw(raw io.Reader) *TestDataSourceConnection {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *TestDataSourceConnection) Request(req *Request) *TestDataSourceConnection {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *TestDataSourceConnection) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for TestDataSourceConnection: %w", err)
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
		path.WriteString("data_source")
		path.WriteString("/")
		path.WriteString("_test")

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
func (r TestDataSourceConnection) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx = instrument.Start(providedCtx, "esql.test_data_source_connection")
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
		instrument.BeforeRequest(req, "esql.test_data_source_connection")
		if reader := instrument.RecordRequestBody(ctx, "esql.test_data_source_connection", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "esql.test_data_source_connection")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the TestDataSourceConnection query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a testdatasourceconnection.Response
func (r TestDataSourceConnection) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.test_data_source_connection")
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

// Header set a key, value pair in the TestDataSourceConnection headers map.
func (r *TestDataSourceConnection) Header(key, value string) *TestDataSourceConnection {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *TestDataSourceConnection) ErrorTrace(errortrace bool) *TestDataSourceConnection {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *TestDataSourceConnection) FilterPath(filterpaths ...string) *TestDataSourceConnection {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and `"exists_time_in_millis":
// 3600000` for computers. When disabled the human readable values will be
// omitted. This makes sense for responses being consumed only by machines.
// API name: human
func (r *TestDataSourceConnection) Human(human bool) *TestDataSourceConnection {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *TestDataSourceConnection) Pretty(pretty bool) *TestDataSourceConnection {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Type-specific connection and authentication settings to test. Uses the same
// structure as the `settings` field in the create or update data source API.
// API name: settings
func (r *TestDataSourceConnection) Settings(settings map[string]json.RawMessage) *TestDataSourceConnection {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Settings = settings
	return r
}

func (r *TestDataSourceConnection) AddSetting(key string, value json.RawMessage) *TestDataSourceConnection {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]json.RawMessage
	if r.req.Settings == nil {
		r.req.Settings = make(map[string]json.RawMessage)
	} else {
		tmp = r.req.Settings
	}

	tmp[key] = value

	r.req.Settings = tmp
	return r
}

// The data source type to test. Must be a known, registered type such as `s3`,
// `gcs`, or `azure`. Unknown types return a `400` error.
// API name: type
func (r *TestDataSourceConnection) Type(type_ string) *TestDataSourceConnection {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Type = type_

	return r
}
