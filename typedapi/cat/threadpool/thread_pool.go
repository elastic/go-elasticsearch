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

// Get thread pool statistics.
//
// Get thread pool statistics for each node in a cluster.
// Returned information includes all built-in thread pools and custom thread
// pools.
// IMPORTANT: cat APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications. For
// application consumption, use the nodes info API.
package threadpool

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/catthreadpoolcolumn"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
)

const (
	threadpoolpatternsMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ThreadPool struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	threadpoolpatterns string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewThreadPool type alias for index.
type NewThreadPool func() *ThreadPool

// NewThreadPoolFunc returns a new instance of ThreadPool with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewThreadPoolFunc(tp elastictransport.Interface) NewThreadPool {
	return func() *ThreadPool {
		n := New(tp)

		return n
	}
}

// Get thread pool statistics.
//
// Get thread pool statistics for each node in a cluster.
// Returned information includes all built-in thread pools and custom thread
// pools.
// IMPORTANT: cat APIs are only intended for human consumption using the command
// line or Kibana console. They are not intended for use by applications. For
// application consumption, use the nodes info API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-thread-pool.html
func New(tp elastictransport.Interface) *ThreadPool {
	r := &ThreadPool{
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
func (r *ThreadPool) HttpRequest(ctx context.Context) (*http.Request, error) {
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
		path.WriteString("thread_pool")

		method = http.MethodGet
	case r.paramSet == threadpoolpatternsMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("thread_pool")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "threadpoolpatterns", r.threadpoolpatterns)
		}
		path.WriteString(r.threadpoolpatterns)

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
func (r ThreadPool) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cat.thread_pool")
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
		instrument.BeforeRequest(req, "cat.thread_pool")
		if reader := instrument.RecordRequestBody(ctx, "cat.thread_pool", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cat.thread_pool")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ThreadPool query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a threadpool.Response
func (r ThreadPool) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.thread_pool")
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
func (r ThreadPool) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cat.thread_pool")
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
		err := fmt.Errorf("an error happened during the ThreadPool query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the ThreadPool headers map.
func (r *ThreadPool) Header(key, value string) *ThreadPool {
	r.headers.Set(key, value)

	return r
}

// ThreadPoolPatterns A comma-separated list of thread pool names used to limit the request.
// Accepts wildcard expressions.
// API Name: threadpoolpatterns
func (r *ThreadPool) ThreadPoolPatterns(threadpoolpatterns string) *ThreadPool {
	r.paramSet |= threadpoolpatternsMask
	r.threadpoolpatterns = threadpoolpatterns

	return r
}

// H List of columns to appear in the response. Supports simple wildcards.
// API name: h
func (r *ThreadPool) H(catthreadpoolcolumns ...catthreadpoolcolumn.CatThreadPoolColumn) *ThreadPool {
	tmp := []string{}
	for _, item := range catthreadpoolcolumns {
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
func (r *ThreadPool) S(names ...string) *ThreadPool {
	r.values.Set("s", strings.Join(names, ","))

	return r
}

// Time The unit used to display time values.
// API name: time
func (r *ThreadPool) Time(time timeunit.TimeUnit) *ThreadPool {
	r.values.Set("time", time.String())

	return r
}

// Local If `true`, the request computes the list of selected nodes from the
// local cluster state. If `false` the list of selected nodes are computed
// from the cluster state of the master node. In both cases the coordinating
// node will send requests for further information to each selected node.
// API name: local
func (r *ThreadPool) Local(local bool) *ThreadPool {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// MasterTimeout The period to wait for a connection to the master node.
// API name: master_timeout
func (r *ThreadPool) MasterTimeout(duration string) *ThreadPool {
	r.values.Set("master_timeout", duration)

	return r
}

// Format Specifies the format to return the columnar data in, can be set to
// `text`, `json`, `cbor`, `yaml`, or `smile`.
// API name: format
func (r *ThreadPool) Format(format string) *ThreadPool {
	r.values.Set("format", format)

	return r
}

// Help When set to `true` will output available columns. This option
// can't be combined with any other query string option.
// API name: help
func (r *ThreadPool) Help(help bool) *ThreadPool {
	r.values.Set("help", strconv.FormatBool(help))

	return r
}

// V When set to `true` will enable verbose output.
// API name: v
func (r *ThreadPool) V(v bool) *ThreadPool {
	r.values.Set("v", strconv.FormatBool(v))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ThreadPool) ErrorTrace(errortrace bool) *ThreadPool {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ThreadPool) FilterPath(filterpaths ...string) *ThreadPool {
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
func (r *ThreadPool) Human(human bool) *ThreadPool {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ThreadPool) Pretty(pretty bool) *ThreadPool {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
