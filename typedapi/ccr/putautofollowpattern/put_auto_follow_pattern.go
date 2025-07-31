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

// Create or update auto-follow patterns.
// Create a collection of cross-cluster replication auto-follow patterns for a
// remote cluster.
// Newly created indices on the remote cluster that match any of the patterns
// are automatically configured as follower indices.
// Indices on the remote cluster that were created before the auto-follow
// pattern was created will not be auto-followed even if they match the pattern.
//
// This API can also be used to update auto-follow patterns.
// NOTE: Follower indices that were configured automatically before updating an
// auto-follow pattern will remain unchanged even if they do not match against
// the new patterns.
package putautofollowpattern

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

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAutoFollowPattern struct {
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

// NewPutAutoFollowPattern type alias for index.
type NewPutAutoFollowPattern func(name string) *PutAutoFollowPattern

// NewPutAutoFollowPatternFunc returns a new instance of PutAutoFollowPattern with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutAutoFollowPatternFunc(tp elastictransport.Interface) NewPutAutoFollowPattern {
	return func(name string) *PutAutoFollowPattern {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Create or update auto-follow patterns.
// Create a collection of cross-cluster replication auto-follow patterns for a
// remote cluster.
// Newly created indices on the remote cluster that match any of the patterns
// are automatically configured as follower indices.
// Indices on the remote cluster that were created before the auto-follow
// pattern was created will not be auto-followed even if they match the pattern.
//
// This API can also be used to update auto-follow patterns.
// NOTE: Follower indices that were configured automatically before updating an
// auto-follow pattern will remain unchanged even if they do not match against
// the new patterns.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ccr-put-auto-follow-pattern.html
func New(tp elastictransport.Interface) *PutAutoFollowPattern {
	r := &PutAutoFollowPattern{
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
func (r *PutAutoFollowPattern) Raw(raw io.Reader) *PutAutoFollowPattern {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutAutoFollowPattern) Request(req *Request) *PutAutoFollowPattern {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutAutoFollowPattern) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutAutoFollowPattern: %w", err)
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
		path.WriteString("_ccr")
		path.WriteString("/")
		path.WriteString("auto_follow")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

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
func (r PutAutoFollowPattern) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ccr.put_auto_follow_pattern")
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
		instrument.BeforeRequest(req, "ccr.put_auto_follow_pattern")
		if reader := instrument.RecordRequestBody(ctx, "ccr.put_auto_follow_pattern", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ccr.put_auto_follow_pattern")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutAutoFollowPattern query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putautofollowpattern.Response
func (r PutAutoFollowPattern) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ccr.put_auto_follow_pattern")
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

// Header set a key, value pair in the PutAutoFollowPattern headers map.
func (r *PutAutoFollowPattern) Header(key, value string) *PutAutoFollowPattern {
	r.headers.Set(key, value)

	return r
}

// Name The name of the collection of auto-follow patterns.
// API Name: name
func (r *PutAutoFollowPattern) _name(name string) *PutAutoFollowPattern {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// API name: master_timeout
func (r *PutAutoFollowPattern) MasterTimeout(duration string) *PutAutoFollowPattern {
	r.values.Set("master_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutAutoFollowPattern) ErrorTrace(errortrace bool) *PutAutoFollowPattern {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutAutoFollowPattern) FilterPath(filterpaths ...string) *PutAutoFollowPattern {
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
func (r *PutAutoFollowPattern) Human(human bool) *PutAutoFollowPattern {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutAutoFollowPattern) Pretty(pretty bool) *PutAutoFollowPattern {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// FollowIndexPattern The name of follower index. The template {{leader_index}} can be used to
// derive the name of the follower index from the name of the leader index. When
// following a data stream, use {{leader_index}}; CCR does not support changes
// to the names of a follower data stream’s backing indices.
// API name: follow_index_pattern
func (r *PutAutoFollowPattern) FollowIndexPattern(indexpattern string) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FollowIndexPattern = &indexpattern

	return r
}

// LeaderIndexExclusionPatterns An array of simple index patterns that can be used to exclude indices from
// being auto-followed. Indices in the remote cluster whose names are matching
// one or more leader_index_patterns and one or more
// leader_index_exclusion_patterns won’t be followed.
// API name: leader_index_exclusion_patterns
func (r *PutAutoFollowPattern) LeaderIndexExclusionPatterns(indexpatterns ...string) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LeaderIndexExclusionPatterns = indexpatterns

	return r
}

// LeaderIndexPatterns An array of simple index patterns to match against indices in the remote
// cluster specified by the remote_cluster field.
// API name: leader_index_patterns
func (r *PutAutoFollowPattern) LeaderIndexPatterns(indexpatterns ...string) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.LeaderIndexPatterns = indexpatterns

	return r
}

// MaxOutstandingReadRequests The maximum number of outstanding reads requests from the remote cluster.
// API name: max_outstanding_read_requests
func (r *PutAutoFollowPattern) MaxOutstandingReadRequests(maxoutstandingreadrequests int) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxOutstandingReadRequests = &maxoutstandingreadrequests

	return r
}

// MaxOutstandingWriteRequests The maximum number of outstanding reads requests from the remote cluster.
// API name: max_outstanding_write_requests
func (r *PutAutoFollowPattern) MaxOutstandingWriteRequests(maxoutstandingwriterequests int) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxOutstandingWriteRequests = &maxoutstandingwriterequests

	return r
}

// MaxReadRequestOperationCount The maximum number of operations to pull per read from the remote cluster.
// API name: max_read_request_operation_count
func (r *PutAutoFollowPattern) MaxReadRequestOperationCount(maxreadrequestoperationcount int) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxReadRequestOperationCount = &maxreadrequestoperationcount

	return r
}

// MaxReadRequestSize The maximum size in bytes of per read of a batch of operations pulled from
// the remote cluster.
// API name: max_read_request_size
func (r *PutAutoFollowPattern) MaxReadRequestSize(bytesize types.ByteSize) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxReadRequestSize = bytesize

	return r
}

// MaxRetryDelay The maximum time to wait before retrying an operation that failed
// exceptionally. An exponential backoff strategy is employed when retrying.
// API name: max_retry_delay
func (r *PutAutoFollowPattern) MaxRetryDelay(duration types.Duration) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxRetryDelay = duration

	return r
}

// MaxWriteBufferCount The maximum number of operations that can be queued for writing. When this
// limit is reached, reads from the remote cluster will be deferred until the
// number of queued operations goes below the limit.
// API name: max_write_buffer_count
func (r *PutAutoFollowPattern) MaxWriteBufferCount(maxwritebuffercount int) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxWriteBufferCount = &maxwritebuffercount

	return r
}

// MaxWriteBufferSize The maximum total bytes of operations that can be queued for writing. When
// this limit is reached, reads from the remote cluster will be deferred until
// the total bytes of queued operations goes below the limit.
// API name: max_write_buffer_size
func (r *PutAutoFollowPattern) MaxWriteBufferSize(bytesize types.ByteSize) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxWriteBufferSize = bytesize

	return r
}

// MaxWriteRequestOperationCount The maximum number of operations per bulk write request executed on the
// follower.
// API name: max_write_request_operation_count
func (r *PutAutoFollowPattern) MaxWriteRequestOperationCount(maxwriterequestoperationcount int) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxWriteRequestOperationCount = &maxwriterequestoperationcount

	return r
}

// MaxWriteRequestSize The maximum total bytes of operations per bulk write request executed on the
// follower.
// API name: max_write_request_size
func (r *PutAutoFollowPattern) MaxWriteRequestSize(bytesize types.ByteSize) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxWriteRequestSize = bytesize

	return r
}

// ReadPollTimeout The maximum time to wait for new operations on the remote cluster when the
// follower index is synchronized with the leader index. When the timeout has
// elapsed, the poll for operations will return to the follower so that it can
// update some statistics. Then the follower will immediately attempt to read
// from the leader again.
// API name: read_poll_timeout
func (r *PutAutoFollowPattern) ReadPollTimeout(duration types.Duration) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ReadPollTimeout = duration

	return r
}

// RemoteCluster The remote cluster containing the leader indices to match against.
// API name: remote_cluster
func (r *PutAutoFollowPattern) RemoteCluster(remotecluster string) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RemoteCluster = remotecluster

	return r
}

// Settings Settings to override from the leader index. Note that certain settings can
// not be overrode (e.g., index.number_of_shards).
// API name: settings
func (r *PutAutoFollowPattern) Settings(settings map[string]json.RawMessage) *PutAutoFollowPattern {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}
