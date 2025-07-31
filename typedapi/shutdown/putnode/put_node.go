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

// Prepare a node to be shut down.
//
// NOTE: This feature is designed for indirect use by Elastic Cloud, Elastic
// Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
// supported.
//
// If you specify a node that is offline, it will be prepared for shut down when
// it rejoins the cluster.
//
// If the operator privileges feature is enabled, you must be an operator to use
// this API.
//
// The API migrates ongoing tasks and index shards to other nodes as needed to
// prepare a node to be restarted or shut down and removed from the cluster.
// This ensures that Elasticsearch can be stopped safely with minimal disruption
// to the cluster.
//
// You must specify the type of shutdown: `restart`, `remove`, or `replace`.
// If a node is already being prepared for shutdown, you can use this API to
// change the shutdown type.
//
// IMPORTANT: This API does NOT terminate the Elasticsearch process.
// Monitor the node shutdown status to determine when it is safe to stop
// Elasticsearch.
package putnode

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/type_"
)

const (
	nodeidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutNode struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutNode type alias for index.
type NewPutNode func(nodeid string) *PutNode

// NewPutNodeFunc returns a new instance of PutNode with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutNodeFunc(tp elastictransport.Interface) NewPutNode {
	return func(nodeid string) *PutNode {
		n := New(tp)

		n._nodeid(nodeid)

		return n
	}
}

// Prepare a node to be shut down.
//
// NOTE: This feature is designed for indirect use by Elastic Cloud, Elastic
// Cloud Enterprise, and Elastic Cloud on Kubernetes. Direct use is not
// supported.
//
// If you specify a node that is offline, it will be prepared for shut down when
// it rejoins the cluster.
//
// If the operator privileges feature is enabled, you must be an operator to use
// this API.
//
// The API migrates ongoing tasks and index shards to other nodes as needed to
// prepare a node to be restarted or shut down and removed from the cluster.
// This ensures that Elasticsearch can be stopped safely with minimal disruption
// to the cluster.
//
// You must specify the type of shutdown: `restart`, `remove`, or `replace`.
// If a node is already being prepared for shutdown, you can use this API to
// change the shutdown type.
//
// IMPORTANT: This API does NOT terminate the Elasticsearch process.
// Monitor the node shutdown status to determine when it is safe to stop
// Elasticsearch.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/put-shutdown.html
func New(tp elastictransport.Interface) *PutNode {
	r := &PutNode{
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
func (r *PutNode) Raw(raw io.Reader) *PutNode {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutNode) Request(req *Request) *PutNode {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutNode) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutNode: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nodeidMask:
		path.WriteString("/")
		path.WriteString("_nodes")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "nodeid", r.nodeid)
		}
		path.WriteString(r.nodeid)
		path.WriteString("/")
		path.WriteString("shutdown")

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
func (r PutNode) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "shutdown.put_node")
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
		instrument.BeforeRequest(req, "shutdown.put_node")
		if reader := instrument.RecordRequestBody(ctx, "shutdown.put_node", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "shutdown.put_node")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutNode query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putnode.Response
func (r PutNode) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "shutdown.put_node")
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

// Header set a key, value pair in the PutNode headers map.
func (r *PutNode) Header(key, value string) *PutNode {
	r.headers.Set(key, value)

	return r
}

// NodeId The node identifier.
// This parameter is not validated against the cluster's active nodes.
// This enables you to register a node for shut down while it is offline.
// No error is thrown if you specify an invalid node ID.
// API Name: nodeid
func (r *PutNode) _nodeid(nodeid string) *PutNode {
	r.paramSet |= nodeidMask
	r.nodeid = nodeid

	return r
}

// MasterTimeout The period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutNode) MasterTimeout(mastertimeout timeunit.TimeUnit) *PutNode {
	r.values.Set("master_timeout", mastertimeout.String())

	return r
}

// Timeout The period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutNode) Timeout(timeout timeunit.TimeUnit) *PutNode {
	r.values.Set("timeout", timeout.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutNode) ErrorTrace(errortrace bool) *PutNode {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutNode) FilterPath(filterpaths ...string) *PutNode {
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
func (r *PutNode) Human(human bool) *PutNode {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutNode) Pretty(pretty bool) *PutNode {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AllocationDelay Only valid if type is restart.
// Controls how long Elasticsearch will wait for the node to restart and join
// the cluster before reassigning its shards to other nodes.
// This works the same as delaying allocation with the
// index.unassigned.node_left.delayed_timeout setting.
// If you specify both a restart allocation delay and an index-level allocation
// delay, the longer of the two is used.
// API name: allocation_delay
func (r *PutNode) AllocationDelay(allocationdelay string) *PutNode {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AllocationDelay = &allocationdelay

	return r
}

// Reason A human-readable reason that the node is being shut down.
// This field provides information for other cluster operators; it does not
// affect the shut down process.
// API name: reason
func (r *PutNode) Reason(reason string) *PutNode {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Reason = reason

	return r
}

// TargetNodeName Only valid if type is replace.
// Specifies the name of the node that is replacing the node being shut down.
// Shards from the shut down node are only allowed to be allocated to the target
// node, and no other data will be allocated to the target node.
// During relocation of data certain allocation rules are ignored, such as disk
// watermarks or user attribute filtering rules.
// API name: target_node_name
func (r *PutNode) TargetNodeName(targetnodename string) *PutNode {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TargetNodeName = &targetnodename

	return r
}

// Type Valid values are restart, remove, or replace.
// Use restart when you need to temporarily shut down a node to perform an
// upgrade, make configuration changes, or perform other maintenance.
// Because the node is expected to rejoin the cluster, data is not migrated off
// of the node.
// Use remove when you need to permanently remove a node from the cluster.
// The node is not marked ready for shutdown until data is migrated off of the
// node Use replace to do a 1:1 replacement of a node with another node.
// Certain allocation decisions will be ignored (such as disk watermarks) in the
// interest of true replacement of the source node with the target node.
// During a replace-type shutdown, rollover and index creation may result in
// unassigned shards, and shrink may fail until the replacement is complete.
// API name: type
func (r *PutNode) Type(type_ type_.Type) *PutNode {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Type = type_

	return r
}
