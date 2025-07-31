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

// Explain the shard allocations.
// Get explanations for shard allocations in the cluster.
// For unassigned shards, it provides an explanation for why the shard is
// unassigned.
// For assigned shards, it provides an explanation for why the shard is
// remaining on its current node and has not moved or rebalanced to another
// node.
// This API can be very useful when attempting to diagnose why a shard is
// unassigned or why a shard continues to remain on its current node when you
// might expect otherwise.
package allocationexplain

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

type AllocationExplain struct {
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

// NewAllocationExplain type alias for index.
type NewAllocationExplain func() *AllocationExplain

// NewAllocationExplainFunc returns a new instance of AllocationExplain with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewAllocationExplainFunc(tp elastictransport.Interface) NewAllocationExplain {
	return func() *AllocationExplain {
		n := New(tp)

		return n
	}
}

// Explain the shard allocations.
// Get explanations for shard allocations in the cluster.
// For unassigned shards, it provides an explanation for why the shard is
// unassigned.
// For assigned shards, it provides an explanation for why the shard is
// remaining on its current node and has not moved or rebalanced to another
// node.
// This API can be very useful when attempting to diagnose why a shard is
// unassigned or why a shard continues to remain on its current node when you
// might expect otherwise.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-allocation-explain.html
func New(tp elastictransport.Interface) *AllocationExplain {
	r := &AllocationExplain{
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
func (r *AllocationExplain) Raw(raw io.Reader) *AllocationExplain {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *AllocationExplain) Request(req *Request) *AllocationExplain {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *AllocationExplain) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for AllocationExplain: %w", err)
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
		path.WriteString("_cluster")
		path.WriteString("/")
		path.WriteString("allocation")
		path.WriteString("/")
		path.WriteString("explain")

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
func (r AllocationExplain) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.allocation_explain")
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
		instrument.BeforeRequest(req, "cluster.allocation_explain")
		if reader := instrument.RecordRequestBody(ctx, "cluster.allocation_explain", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.allocation_explain")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the AllocationExplain query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a allocationexplain.Response
func (r AllocationExplain) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.allocation_explain")
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

// Header set a key, value pair in the AllocationExplain headers map.
func (r *AllocationExplain) Header(key, value string) *AllocationExplain {
	r.headers.Set(key, value)

	return r
}

// IncludeDiskInfo If true, returns information about disk usage and shard sizes.
// API name: include_disk_info
func (r *AllocationExplain) IncludeDiskInfo(includediskinfo bool) *AllocationExplain {
	r.values.Set("include_disk_info", strconv.FormatBool(includediskinfo))

	return r
}

// IncludeYesDecisions If true, returns YES decisions in explanation.
// API name: include_yes_decisions
func (r *AllocationExplain) IncludeYesDecisions(includeyesdecisions bool) *AllocationExplain {
	r.values.Set("include_yes_decisions", strconv.FormatBool(includeyesdecisions))

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// API name: master_timeout
func (r *AllocationExplain) MasterTimeout(duration string) *AllocationExplain {
	r.values.Set("master_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *AllocationExplain) ErrorTrace(errortrace bool) *AllocationExplain {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *AllocationExplain) FilterPath(filterpaths ...string) *AllocationExplain {
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
func (r *AllocationExplain) Human(human bool) *AllocationExplain {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *AllocationExplain) Pretty(pretty bool) *AllocationExplain {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// CurrentNode Specifies the node ID or the name of the node to only explain a shard that is
// currently located on the specified node.
// API name: current_node
func (r *AllocationExplain) CurrentNode(currentnode string) *AllocationExplain {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.CurrentNode = &currentnode

	return r
}

// Index Specifies the name of the index that you would like an explanation for.
// API name: index
func (r *AllocationExplain) Index(indexname string) *AllocationExplain {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Index = &indexname

	return r
}

// Primary If true, returns explanation for the primary shard for the given shard ID.
// API name: primary
func (r *AllocationExplain) Primary(primary bool) *AllocationExplain {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Primary = &primary

	return r
}

// Shard Specifies the ID of the shard that you would like an explanation for.
// API name: shard
func (r *AllocationExplain) Shard(shard int) *AllocationExplain {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Shard = &shard

	return r
}
