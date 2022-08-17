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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package follow

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package follow
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/follow/CreateFollowIndexRequest.ts#L25-L52
type Request struct {
	LeaderIndex *types.IndexName `json:"leader_index,omitempty"`

	MaxOutstandingReadRequests *int64 `json:"max_outstanding_read_requests,omitempty"`

	MaxOutstandingWriteRequests *int64 `json:"max_outstanding_write_requests,omitempty"`

	MaxReadRequestOperationCount *int64 `json:"max_read_request_operation_count,omitempty"`

	MaxReadRequestSize *string `json:"max_read_request_size,omitempty"`

	MaxRetryDelay *types.Duration `json:"max_retry_delay,omitempty"`

	MaxWriteBufferCount *int64 `json:"max_write_buffer_count,omitempty"`

	MaxWriteBufferSize *string `json:"max_write_buffer_size,omitempty"`

	MaxWriteRequestOperationCount *int64 `json:"max_write_request_operation_count,omitempty"`

	MaxWriteRequestSize *string `json:"max_write_request_size,omitempty"`

	ReadPollTimeout *types.Duration `json:"read_poll_timeout,omitempty"`

	RemoteCluster *string `json:"remote_cluster,omitempty"`
}

// RequestBuilder is the builder API for the follow.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Follow request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) LeaderIndex(leaderindex types.IndexName) *RequestBuilder {
	rb.v.LeaderIndex = &leaderindex
	return rb
}

func (rb *RequestBuilder) MaxOutstandingReadRequests(maxoutstandingreadrequests int64) *RequestBuilder {
	rb.v.MaxOutstandingReadRequests = &maxoutstandingreadrequests
	return rb
}

func (rb *RequestBuilder) MaxOutstandingWriteRequests(maxoutstandingwriterequests int64) *RequestBuilder {
	rb.v.MaxOutstandingWriteRequests = &maxoutstandingwriterequests
	return rb
}

func (rb *RequestBuilder) MaxReadRequestOperationCount(maxreadrequestoperationcount int64) *RequestBuilder {
	rb.v.MaxReadRequestOperationCount = &maxreadrequestoperationcount
	return rb
}

func (rb *RequestBuilder) MaxReadRequestSize(maxreadrequestsize string) *RequestBuilder {
	rb.v.MaxReadRequestSize = &maxreadrequestsize
	return rb
}

func (rb *RequestBuilder) MaxRetryDelay(maxretrydelay *types.DurationBuilder) *RequestBuilder {
	v := maxretrydelay.Build()
	rb.v.MaxRetryDelay = &v
	return rb
}

func (rb *RequestBuilder) MaxWriteBufferCount(maxwritebuffercount int64) *RequestBuilder {
	rb.v.MaxWriteBufferCount = &maxwritebuffercount
	return rb
}

func (rb *RequestBuilder) MaxWriteBufferSize(maxwritebuffersize string) *RequestBuilder {
	rb.v.MaxWriteBufferSize = &maxwritebuffersize
	return rb
}

func (rb *RequestBuilder) MaxWriteRequestOperationCount(maxwriterequestoperationcount int64) *RequestBuilder {
	rb.v.MaxWriteRequestOperationCount = &maxwriterequestoperationcount
	return rb
}

func (rb *RequestBuilder) MaxWriteRequestSize(maxwriterequestsize string) *RequestBuilder {
	rb.v.MaxWriteRequestSize = &maxwriterequestsize
	return rb
}

func (rb *RequestBuilder) ReadPollTimeout(readpolltimeout *types.DurationBuilder) *RequestBuilder {
	v := readpolltimeout.Build()
	rb.v.ReadPollTimeout = &v
	return rb
}

func (rb *RequestBuilder) RemoteCluster(remotecluster string) *RequestBuilder {
	rb.v.RemoteCluster = &remotecluster
	return rb
}
