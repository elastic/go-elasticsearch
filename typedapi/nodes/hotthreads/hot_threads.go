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

// Get the hot threads for nodes.
// Get a breakdown of the hot threads on each selected node in the cluster.
// The output is plain text with a breakdown of the top hot threads for each
// node.
package hotthreads

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/threadtype"
)

const (
	nodeidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HotThreads struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewHotThreads type alias for index.
type NewHotThreads func() *HotThreads

// NewHotThreadsFunc returns a new instance of HotThreads with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewHotThreadsFunc(tp elastictransport.Interface) NewHotThreads {
	return func() *HotThreads {
		n := New(tp)

		return n
	}
}

// Get the hot threads for nodes.
// Get a breakdown of the hot threads on each selected node in the cluster.
// The output is plain text with a breakdown of the top hot threads for each
// node.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-nodes-hot-threads.html
func New(tp elastictransport.Interface) *HotThreads {
	r := &HotThreads{
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
func (r *HotThreads) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")
		path.WriteString("hot_threads")

		method = http.MethodGet
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "nodeid", r.nodeid)
		}
		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("hot_threads")

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
		req.Header.Set("Accept", "text/plain")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r HotThreads) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "nodes.hot_threads")
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
		instrument.BeforeRequest(req, "nodes.hot_threads")
		if reader := instrument.RecordRequestBody(ctx, "nodes.hot_threads", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "nodes.hot_threads")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the HotThreads query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a hotthreads.Response
func (r HotThreads) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "nodes.hot_threads")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r HotThreads) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "nodes.hot_threads")
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
		err := fmt.Errorf("an error happened during the HotThreads query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the HotThreads headers map.
func (r *HotThreads) Header(key, value string) *HotThreads {
	r.headers.Set(key, value)

	return r
}

// NodeId List of node IDs or names used to limit returned information.
// API Name: nodeid
func (r *HotThreads) NodeId(nodeid string) *HotThreads {
	r.paramSet |= nodeidMask
	r.nodeid = nodeid

	return r
}

// IgnoreIdleThreads If true, known idle threads (e.g. waiting in a socket select, or to get
// a task from an empty queue) are filtered out.
// API name: ignore_idle_threads
func (r *HotThreads) IgnoreIdleThreads(ignoreidlethreads bool) *HotThreads {
	r.values.Set("ignore_idle_threads", strconv.FormatBool(ignoreidlethreads))

	return r
}

// Interval The interval to do the second sampling of threads.
// API name: interval
func (r *HotThreads) Interval(duration string) *HotThreads {
	r.values.Set("interval", duration)

	return r
}

// Snapshots Number of samples of thread stacktrace.
// API name: snapshots
func (r *HotThreads) Snapshots(snapshots string) *HotThreads {
	r.values.Set("snapshots", snapshots)

	return r
}

// Threads Specifies the number of hot threads to provide information for.
// API name: threads
func (r *HotThreads) Threads(threads string) *HotThreads {
	r.values.Set("threads", threads)

	return r
}

// Timeout Period to wait for a response. If no response is received
// before the timeout expires, the request fails and returns an error.
// API name: timeout
func (r *HotThreads) Timeout(duration string) *HotThreads {
	r.values.Set("timeout", duration)

	return r
}

// Type The type to sample.
// API name: type
func (r *HotThreads) Type(type_ threadtype.ThreadType) *HotThreads {
	r.values.Set("type", type_.String())

	return r
}

// Sort The sort order for 'cpu' type (default: total)
// API name: sort
func (r *HotThreads) Sort(sort threadtype.ThreadType) *HotThreads {
	r.values.Set("sort", sort.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *HotThreads) ErrorTrace(errortrace bool) *HotThreads {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *HotThreads) FilterPath(filterpaths ...string) *HotThreads {
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
func (r *HotThreads) Human(human bool) *HotThreads {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *HotThreads) Pretty(pretty bool) *HotThreads {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
