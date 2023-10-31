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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Creates a new named collection of auto-follow patterns against a specified
// remote cluster. Newly created indices on the remote cluster matching any of
// the specified patterns will be automatically configured as follower indices.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	name string
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

// Creates a new named collection of auto-follow patterns against a specified
// remote cluster. Newly created indices on the remote cluster matching any of
// the specified patterns will be automatically configured as follower indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/ccr-put-auto-follow-pattern.html
func New(tp elastictransport.Interface) *PutAutoFollowPattern {
	r := &PutAutoFollowPattern{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutAutoFollowPattern: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_ccr")
		path.WriteString("/")
		path.WriteString("auto_follow")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
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
func (r PutAutoFollowPattern) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutAutoFollowPattern query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putautofollowpattern.Response
func (r PutAutoFollowPattern) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
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

// FollowIndexPattern The name of follower index. The template {{leader_index}} can be used to
// derive the name of the follower index from the name of the leader index. When
// following a data stream, use {{leader_index}}; CCR does not support changes
// to the names of a follower data stream’s backing indices.
// API name: follow_index_pattern
func (r *PutAutoFollowPattern) FollowIndexPattern(indexpattern string) *PutAutoFollowPattern {
	r.req.FollowIndexPattern = &indexpattern

	return r
}

// LeaderIndexExclusionPatterns An array of simple index patterns that can be used to exclude indices from
// being auto-followed. Indices in the remote cluster whose names are matching
// one or more leader_index_patterns and one or more
// leader_index_exclusion_patterns won’t be followed.
// API name: leader_index_exclusion_patterns
func (r *PutAutoFollowPattern) LeaderIndexExclusionPatterns(indexpatterns ...string) *PutAutoFollowPattern {
	r.req.LeaderIndexExclusionPatterns = indexpatterns

	return r
}

// LeaderIndexPatterns An array of simple index patterns to match against indices in the remote
// cluster specified by the remote_cluster field.
// API name: leader_index_patterns
func (r *PutAutoFollowPattern) LeaderIndexPatterns(indexpatterns ...string) *PutAutoFollowPattern {
	r.req.LeaderIndexPatterns = indexpatterns

	return r
}

// MaxOutstandingReadRequests The maximum number of outstanding reads requests from the remote cluster.
// API name: max_outstanding_read_requests
func (r *PutAutoFollowPattern) MaxOutstandingReadRequests(maxoutstandingreadrequests int) *PutAutoFollowPattern {
	r.req.MaxOutstandingReadRequests = &maxoutstandingreadrequests

	return r
}

// MaxOutstandingWriteRequests The maximum number of outstanding reads requests from the remote cluster.
// API name: max_outstanding_write_requests
func (r *PutAutoFollowPattern) MaxOutstandingWriteRequests(maxoutstandingwriterequests int) *PutAutoFollowPattern {
	r.req.MaxOutstandingWriteRequests = &maxoutstandingwriterequests

	return r
}

// MaxReadRequestOperationCount The maximum number of operations to pull per read from the remote cluster.
// API name: max_read_request_operation_count
func (r *PutAutoFollowPattern) MaxReadRequestOperationCount(maxreadrequestoperationcount int) *PutAutoFollowPattern {
	r.req.MaxReadRequestOperationCount = &maxreadrequestoperationcount

	return r
}

// MaxReadRequestSize The maximum size in bytes of per read of a batch of operations pulled from
// the remote cluster.
// API name: max_read_request_size
func (r *PutAutoFollowPattern) MaxReadRequestSize(bytesize types.ByteSize) *PutAutoFollowPattern {
	r.req.MaxReadRequestSize = bytesize

	return r
}

// MaxRetryDelay The maximum time to wait before retrying an operation that failed
// exceptionally. An exponential backoff strategy is employed when retrying.
// API name: max_retry_delay
func (r *PutAutoFollowPattern) MaxRetryDelay(duration types.Duration) *PutAutoFollowPattern {
	r.req.MaxRetryDelay = duration

	return r
}

// MaxWriteBufferCount The maximum number of operations that can be queued for writing. When this
// limit is reached, reads from the remote cluster will be deferred until the
// number of queued operations goes below the limit.
// API name: max_write_buffer_count
func (r *PutAutoFollowPattern) MaxWriteBufferCount(maxwritebuffercount int) *PutAutoFollowPattern {
	r.req.MaxWriteBufferCount = &maxwritebuffercount

	return r
}

// MaxWriteBufferSize The maximum total bytes of operations that can be queued for writing. When
// this limit is reached, reads from the remote cluster will be deferred until
// the total bytes of queued operations goes below the limit.
// API name: max_write_buffer_size
func (r *PutAutoFollowPattern) MaxWriteBufferSize(bytesize types.ByteSize) *PutAutoFollowPattern {
	r.req.MaxWriteBufferSize = bytesize

	return r
}

// MaxWriteRequestOperationCount The maximum number of operations per bulk write request executed on the
// follower.
// API name: max_write_request_operation_count
func (r *PutAutoFollowPattern) MaxWriteRequestOperationCount(maxwriterequestoperationcount int) *PutAutoFollowPattern {
	r.req.MaxWriteRequestOperationCount = &maxwriterequestoperationcount

	return r
}

// MaxWriteRequestSize The maximum total bytes of operations per bulk write request executed on the
// follower.
// API name: max_write_request_size
func (r *PutAutoFollowPattern) MaxWriteRequestSize(bytesize types.ByteSize) *PutAutoFollowPattern {
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
	r.req.ReadPollTimeout = duration

	return r
}

// RemoteCluster The remote cluster containing the leader indices to match against.
// API name: remote_cluster
func (r *PutAutoFollowPattern) RemoteCluster(remotecluster string) *PutAutoFollowPattern {

	r.req.RemoteCluster = remotecluster

	return r
}

// Settings Settings to override from the leader index. Note that certain settings can
// not be overrode (e.g., index.number_of_shards).
// API name: settings
func (r *PutAutoFollowPattern) Settings(settings map[string]json.RawMessage) *PutAutoFollowPattern {

	r.req.Settings = settings

	return r
}
