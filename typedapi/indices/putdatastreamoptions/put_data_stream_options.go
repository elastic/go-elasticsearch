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

// Update data stream options.
// Update the data stream options of the specified data streams.
package putdatastreamoptions

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataStreamOptions struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutDataStreamOptions type alias for index.
type NewPutDataStreamOptions func(name string) *PutDataStreamOptions

// NewPutDataStreamOptionsFunc returns a new instance of PutDataStreamOptions with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDataStreamOptionsFunc(tp elastictransport.Interface) NewPutDataStreamOptions {
	return func(name string) *PutDataStreamOptions {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Update data stream options.
// Update the data stream options of the specified data streams.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html
func New(tp elastictransport.Interface) *PutDataStreamOptions {
	r := &PutDataStreamOptions{
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
func (r *PutDataStreamOptions) Raw(raw io.Reader) *PutDataStreamOptions {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDataStreamOptions) Request(req *Request) *PutDataStreamOptions {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDataStreamOptions) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDataStreamOptions: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)
		path.WriteString("/")
		path.WriteString("_options")

		method = http.MethodPut
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
func (r PutDataStreamOptions) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_data_stream_options")
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
		instrument.BeforeRequest(req, "indices.put_data_stream_options")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_data_stream_options", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_data_stream_options")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutDataStreamOptions query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdatastreamoptions.Response
func (r PutDataStreamOptions) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_data_stream_options")
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

// Header set a key, value pair in the PutDataStreamOptions headers map.
func (r *PutDataStreamOptions) Header(key, value string) *PutDataStreamOptions {
	r.headers.Set(key, value)

	return r
}

// Name Comma-separated list of data streams used to limit the request.
// Supports wildcards (`*`).
// To target all data streams use `*` or `_all`.
// API Name: name
func (r *PutDataStreamOptions) _name(name string) *PutDataStreamOptions {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// ExpandWildcards Type of data stream that wildcard patterns can match.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *PutDataStreamOptions) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutDataStreamOptions {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an
// error.
// API name: master_timeout
func (r *PutDataStreamOptions) MasterTimeout(duration string) *PutDataStreamOptions {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutDataStreamOptions) Timeout(duration string) *PutDataStreamOptions {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutDataStreamOptions) ErrorTrace(errortrace bool) *PutDataStreamOptions {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutDataStreamOptions) FilterPath(filterpaths ...string) *PutDataStreamOptions {
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
func (r *PutDataStreamOptions) Human(human bool) *PutDataStreamOptions {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutDataStreamOptions) Pretty(pretty bool) *PutDataStreamOptions {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// If defined, it will update the failure store configuration of every data
// stream resolved by the name expression.
// API name: failure_store
func (r *PutDataStreamOptions) FailureStore(failurestore types.DataStreamFailureStoreVariant) *PutDataStreamOptions {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.FailureStore = failurestore.DataStreamFailureStoreCaster()

	return r
}
