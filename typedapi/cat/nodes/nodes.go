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

// Get node information.
//
// Get information about the nodes in a cluster.
// IMPORTANT: cat APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications. For
// application consumption, use the nodes info API.
package nodes

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/catnodecolumn"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Nodes struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewNodes type alias for index.
type NewNodes func() *Nodes

// NewNodesFunc returns a new instance of Nodes with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewNodesFunc(tp elastictransport.Interface) NewNodes {
	return func() *Nodes {
		n := New(tp)

		return n
	}
}

// Get node information.
//
// Get information about the nodes in a cluster.
// IMPORTANT: cat APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications. For
// application consumption, use the nodes info API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodes.html
func New(tp elastictransport.Interface) *Nodes {
	r := &Nodes{
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
func (r *Nodes) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("nodes")

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
func (r Nodes) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cat.nodes")
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
		instrument.BeforeRequest(req, "cat.nodes")
		if reader := instrument.RecordRequestBody(ctx, "cat.nodes", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cat.nodes")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Nodes query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a nodes.Response
func (r Nodes) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.nodes")
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
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r Nodes) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.nodes")
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
		err := fmt.Errorf("an error happened during the Nodes query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Nodes headers map.
func (r *Nodes) Header(key, value string) *Nodes {
	r.headers.Set(key, value)

	return r
}

// Bytes The unit used to display byte values.
// API name: bytes
func (r *Nodes) Bytes(bytes bytes.Bytes) *Nodes {
	r.values.Set("bytes", bytes.String())

	return r
}

// FullId If `true`, return the full node ID. If `false`, return the shortened node ID.
// API name: full_id
func (r *Nodes) FullId(fullid string) *Nodes {
	r.values.Set("full_id", fullid)

	return r
}

// IncludeUnloadedSegments If true, the response includes information from segments that are not loaded
// into memory.
// API name: include_unloaded_segments
func (r *Nodes) IncludeUnloadedSegments(includeunloadedsegments bool) *Nodes {
	r.values.Set("include_unloaded_segments", strconv.FormatBool(includeunloadedsegments))

	return r
}

// H A comma-separated list of columns names to display.
// It supports simple wildcards.
// API name: h
func (r *Nodes) H(catnodecolumns ...catnodecolumn.CatNodeColumn) *Nodes {
	tmp := []string{}
	for _, item := range catnodecolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// S A comma-separated list of column names or aliases that determines the sort
// order.
// Sorting defaults to ascending and can be changed by setting `:asc`
// or `:desc` as a suffix to the column name.
// API name: s
func (r *Nodes) S(names ...string) *Nodes {
	r.values.Set("s", strings.Join(names, ","))

	return r
}

// MasterTimeout The period to wait for a connection to the master node.
// API name: master_timeout
func (r *Nodes) MasterTimeout(duration string) *Nodes {
	r.values.Set("master_timeout", duration)

	return r
}

// Time The unit used to display time values.
// API name: time
func (r *Nodes) Time(time timeunit.TimeUnit) *Nodes {
	r.values.Set("time", time.String())

	return r
}

// Format Specifies the format to return the columnar data in, can be set to
// `text`, `json`, `cbor`, `yaml`, or `smile`.
// API name: format
func (r *Nodes) Format(format string) *Nodes {
	r.values.Set("format", format)

	return r
}

// Help When set to `true` will output available columns. This option
// can't be combined with any other query string option.
// API name: help
func (r *Nodes) Help(help bool) *Nodes {
	r.values.Set("help", strconv.FormatBool(help))

	return r
}

// V When set to `true` will enable verbose output.
// API name: v
func (r *Nodes) V(v bool) *Nodes {
	r.values.Set("v", strconv.FormatBool(v))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Nodes) ErrorTrace(errortrace bool) *Nodes {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Nodes) FilterPath(filterpaths ...string) *Nodes {
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
func (r *Nodes) Human(human bool) *Nodes {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Nodes) Pretty(pretty bool) *Nodes {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
