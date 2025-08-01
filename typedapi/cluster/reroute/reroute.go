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

// Reroute the cluster.
// Manually change the allocation of individual shards in the cluster.
// For example, a shard can be moved from one node to another explicitly, an
// allocation can be canceled, and an unassigned shard can be explicitly
// allocated to a specific node.
//
// It is important to note that after processing any reroute commands
// Elasticsearch will perform rebalancing as normal (respecting the values of
// settings such as `cluster.routing.rebalance.enable`) in order to remain in a
// balanced state.
// For example, if the requested allocation includes moving a shard from node1
// to node2 then this may cause a shard to be moved from node2 back to node1 to
// even things out.
//
// The cluster can be set to disable allocations using the
// `cluster.routing.allocation.enable` setting.
// If allocations are disabled then the only allocations that will be performed
// are explicit ones given using the reroute command, and consequent allocations
// due to rebalancing.
//
// The cluster will attempt to allocate a shard a maximum of
// `index.allocation.max_retries` times in a row (defaults to `5`), before
// giving up and leaving the shard unallocated.
// This scenario can be caused by structural problems such as having an analyzer
// which refers to a stopwords file which doesn’t exist on all nodes.
//
// Once the problem has been corrected, allocation can be manually retried by
// calling the reroute API with the `?retry_failed` URI query parameter, which
// will attempt a single retry round for these shards.
package reroute

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

type Reroute struct {
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

// NewReroute type alias for index.
type NewReroute func() *Reroute

// NewRerouteFunc returns a new instance of Reroute with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRerouteFunc(tp elastictransport.Interface) NewReroute {
	return func() *Reroute {
		n := New(tp)

		return n
	}
}

// Reroute the cluster.
// Manually change the allocation of individual shards in the cluster.
// For example, a shard can be moved from one node to another explicitly, an
// allocation can be canceled, and an unassigned shard can be explicitly
// allocated to a specific node.
//
// It is important to note that after processing any reroute commands
// Elasticsearch will perform rebalancing as normal (respecting the values of
// settings such as `cluster.routing.rebalance.enable`) in order to remain in a
// balanced state.
// For example, if the requested allocation includes moving a shard from node1
// to node2 then this may cause a shard to be moved from node2 back to node1 to
// even things out.
//
// The cluster can be set to disable allocations using the
// `cluster.routing.allocation.enable` setting.
// If allocations are disabled then the only allocations that will be performed
// are explicit ones given using the reroute command, and consequent allocations
// due to rebalancing.
//
// The cluster will attempt to allocate a shard a maximum of
// `index.allocation.max_retries` times in a row (defaults to `5`), before
// giving up and leaving the shard unallocated.
// This scenario can be caused by structural problems such as having an analyzer
// which refers to a stopwords file which doesn’t exist on all nodes.
//
// Once the problem has been corrected, allocation can be manually retried by
// calling the reroute API with the `?retry_failed` URI query parameter, which
// will attempt a single retry round for these shards.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-cluster-reroute
func New(tp elastictransport.Interface) *Reroute {
	r := &Reroute{
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
func (r *Reroute) Raw(raw io.Reader) *Reroute {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Reroute) Request(req *Request) *Reroute {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Reroute) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Reroute: %w", err)
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
		path.WriteString("reroute")

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
func (r Reroute) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.reroute")
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
		instrument.BeforeRequest(req, "cluster.reroute")
		if reader := instrument.RecordRequestBody(ctx, "cluster.reroute", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.reroute")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Reroute query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a reroute.Response
func (r Reroute) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.reroute")
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

// Header set a key, value pair in the Reroute headers map.
func (r *Reroute) Header(key, value string) *Reroute {
	r.headers.Set(key, value)

	return r
}

// DryRun If true, then the request simulates the operation.
// It will calculate the result of applying the commands to the current cluster
// state and return the resulting cluster state after the commands (and
// rebalancing) have been applied; it will not actually perform the requested
// changes.
// API name: dry_run
func (r *Reroute) DryRun(dryrun bool) *Reroute {
	r.values.Set("dry_run", strconv.FormatBool(dryrun))

	return r
}

// Explain If true, then the response contains an explanation of why the commands can or
// cannot run.
// API name: explain
func (r *Reroute) Explain(explain bool) *Reroute {
	r.values.Set("explain", strconv.FormatBool(explain))

	return r
}

// Metric Limits the information returned to the specified metrics.
// API name: metric
func (r *Reroute) Metric(metrics ...string) *Reroute {
	r.values.Set("metric", strings.Join(metrics, ","))

	return r
}

// RetryFailed If true, then retries allocation of shards that are blocked due to too many
// subsequent allocation failures.
// API name: retry_failed
func (r *Reroute) RetryFailed(retryfailed bool) *Reroute {
	r.values.Set("retry_failed", strconv.FormatBool(retryfailed))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Reroute) MasterTimeout(duration string) *Reroute {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response. If no response is received before the timeout
// expires, the request fails and returns an error.
// API name: timeout
func (r *Reroute) Timeout(duration string) *Reroute {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Reroute) ErrorTrace(errortrace bool) *Reroute {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Reroute) FilterPath(filterpaths ...string) *Reroute {
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
func (r *Reroute) Human(human bool) *Reroute {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Reroute) Pretty(pretty bool) *Reroute {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Defines the commands to perform.
// API name: commands
func (r *Reroute) Commands(commands ...types.CommandVariant) *Reroute {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range commands {

		r.req.Commands = append(r.req.Commands, *v.CommandCaster())

	}
	return r
}
