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

// Creates a new follower index configured to follow the referenced leader
// index.
package follow

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
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Follow struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
}

// NewFollow type alias for index.
type NewFollow func(index string) *Follow

// NewFollowFunc returns a new instance of Follow with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFollowFunc(tp elastictransport.Interface) NewFollow {
	return func(index string) *Follow {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Creates a new follower index configured to follow the referenced leader
// index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/ccr-put-follow.html
func New(tp elastictransport.Interface) *Follow {
	r := &Follow{
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
func (r *Follow) Raw(raw io.Reader) *Follow {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Follow) Request(req *Request) *Follow {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Follow) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Follow: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_ccr")
		path.WriteString("/")
		path.WriteString("follow")

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
func (r Follow) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Follow query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a follow.Response
func (r Follow) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the Follow headers map.
func (r *Follow) Header(key, value string) *Follow {
	r.headers.Set(key, value)

	return r
}

// Index The name of the follower index
// API Name: index
func (r *Follow) _index(index string) *Follow {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// WaitForActiveShards Sets the number of shard copies that must be active before returning.
// Defaults to 0. Set to `all` for all shard copies, otherwise set to any
// non-negative value less than or equal to the total number of copies for the
// shard (number of replicas + 1)
// API name: wait_for_active_shards
func (r *Follow) WaitForActiveShards(waitforactiveshards string) *Follow {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// API name: leader_index
func (r *Follow) LeaderIndex(indexname string) *Follow {
	r.req.LeaderIndex = &indexname

	return r
}

// API name: max_outstanding_read_requests
func (r *Follow) MaxOutstandingReadRequests(maxoutstandingreadrequests int64) *Follow {

	r.req.MaxOutstandingReadRequests = &maxoutstandingreadrequests

	return r
}

// API name: max_outstanding_write_requests
func (r *Follow) MaxOutstandingWriteRequests(maxoutstandingwriterequests int64) *Follow {

	r.req.MaxOutstandingWriteRequests = &maxoutstandingwriterequests

	return r
}

// API name: max_read_request_operation_count
func (r *Follow) MaxReadRequestOperationCount(maxreadrequestoperationcount int64) *Follow {

	r.req.MaxReadRequestOperationCount = &maxreadrequestoperationcount

	return r
}

// API name: max_read_request_size
func (r *Follow) MaxReadRequestSize(maxreadrequestsize string) *Follow {

	r.req.MaxReadRequestSize = &maxreadrequestsize

	return r
}

// API name: max_retry_delay
func (r *Follow) MaxRetryDelay(duration types.Duration) *Follow {
	r.req.MaxRetryDelay = duration

	return r
}

// API name: max_write_buffer_count
func (r *Follow) MaxWriteBufferCount(maxwritebuffercount int64) *Follow {

	r.req.MaxWriteBufferCount = &maxwritebuffercount

	return r
}

// API name: max_write_buffer_size
func (r *Follow) MaxWriteBufferSize(maxwritebuffersize string) *Follow {

	r.req.MaxWriteBufferSize = &maxwritebuffersize

	return r
}

// API name: max_write_request_operation_count
func (r *Follow) MaxWriteRequestOperationCount(maxwriterequestoperationcount int64) *Follow {

	r.req.MaxWriteRequestOperationCount = &maxwriterequestoperationcount

	return r
}

// API name: max_write_request_size
func (r *Follow) MaxWriteRequestSize(maxwriterequestsize string) *Follow {

	r.req.MaxWriteRequestSize = &maxwriterequestsize

	return r
}

// API name: read_poll_timeout
func (r *Follow) ReadPollTimeout(duration types.Duration) *Follow {
	r.req.ReadPollTimeout = duration

	return r
}

// API name: remote_cluster
func (r *Follow) RemoteCluster(remotecluster string) *Follow {

	r.req.RemoteCluster = &remotecluster

	return r
}
