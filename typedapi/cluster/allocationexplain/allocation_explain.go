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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

// Provides explanations for shard allocations in the cluster.
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

// Provides explanations for shard allocations in the cluster.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-allocation-explain.html
func New(tp elastictransport.Interface) *AllocationExplain {
	r := &AllocationExplain{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),

		req: NewRequest(),
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

// CurrentNode Specifies the node ID or the name of the node to only explain a shard that is
// currently located on the specified node.
// API name: current_node
func (r *AllocationExplain) CurrentNode(currentnode string) *AllocationExplain {

	r.req.CurrentNode = &currentnode

	return r
}

// Index Specifies the name of the index that you would like an explanation for.
// API name: index
func (r *AllocationExplain) Index(indexname string) *AllocationExplain {
	r.req.Index = &indexname

	return r
}

// Primary If true, returns explanation for the primary shard for the given shard ID.
// API name: primary
func (r *AllocationExplain) Primary(primary bool) *AllocationExplain {
	r.req.Primary = &primary

	return r
}

// Shard Specifies the ID of the shard that you would like an explanation for.
// API name: shard
func (r *AllocationExplain) Shard(shard int) *AllocationExplain {
	r.req.Shard = &shard

	return r
}
