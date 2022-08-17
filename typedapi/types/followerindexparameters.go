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


package types

// FollowerIndexParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ccr/follow_info/types.ts#L38-L49
type FollowerIndexParameters struct {
	MaxOutstandingReadRequests    int      `json:"max_outstanding_read_requests"`
	MaxOutstandingWriteRequests   int      `json:"max_outstanding_write_requests"`
	MaxReadRequestOperationCount  int      `json:"max_read_request_operation_count"`
	MaxReadRequestSize            string   `json:"max_read_request_size"`
	MaxRetryDelay                 Duration `json:"max_retry_delay"`
	MaxWriteBufferCount           int      `json:"max_write_buffer_count"`
	MaxWriteBufferSize            string   `json:"max_write_buffer_size"`
	MaxWriteRequestOperationCount int      `json:"max_write_request_operation_count"`
	MaxWriteRequestSize           string   `json:"max_write_request_size"`
	ReadPollTimeout               Duration `json:"read_poll_timeout"`
}

// FollowerIndexParametersBuilder holds FollowerIndexParameters struct and provides a builder API.
type FollowerIndexParametersBuilder struct {
	v *FollowerIndexParameters
}

// NewFollowerIndexParameters provides a builder for the FollowerIndexParameters struct.
func NewFollowerIndexParametersBuilder() *FollowerIndexParametersBuilder {
	r := FollowerIndexParametersBuilder{
		&FollowerIndexParameters{},
	}

	return &r
}

// Build finalize the chain and returns the FollowerIndexParameters struct
func (rb *FollowerIndexParametersBuilder) Build() FollowerIndexParameters {
	return *rb.v
}

func (rb *FollowerIndexParametersBuilder) MaxOutstandingReadRequests(maxoutstandingreadrequests int) *FollowerIndexParametersBuilder {
	rb.v.MaxOutstandingReadRequests = maxoutstandingreadrequests
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxOutstandingWriteRequests(maxoutstandingwriterequests int) *FollowerIndexParametersBuilder {
	rb.v.MaxOutstandingWriteRequests = maxoutstandingwriterequests
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxReadRequestOperationCount(maxreadrequestoperationcount int) *FollowerIndexParametersBuilder {
	rb.v.MaxReadRequestOperationCount = maxreadrequestoperationcount
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxReadRequestSize(maxreadrequestsize string) *FollowerIndexParametersBuilder {
	rb.v.MaxReadRequestSize = maxreadrequestsize
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxRetryDelay(maxretrydelay *DurationBuilder) *FollowerIndexParametersBuilder {
	v := maxretrydelay.Build()
	rb.v.MaxRetryDelay = v
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxWriteBufferCount(maxwritebuffercount int) *FollowerIndexParametersBuilder {
	rb.v.MaxWriteBufferCount = maxwritebuffercount
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxWriteBufferSize(maxwritebuffersize string) *FollowerIndexParametersBuilder {
	rb.v.MaxWriteBufferSize = maxwritebuffersize
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxWriteRequestOperationCount(maxwriterequestoperationcount int) *FollowerIndexParametersBuilder {
	rb.v.MaxWriteRequestOperationCount = maxwriterequestoperationcount
	return rb
}

func (rb *FollowerIndexParametersBuilder) MaxWriteRequestSize(maxwriterequestsize string) *FollowerIndexParametersBuilder {
	rb.v.MaxWriteRequestSize = maxwriterequestsize
	return rb
}

func (rb *FollowerIndexParametersBuilder) ReadPollTimeout(readpolltimeout *DurationBuilder) *FollowerIndexParametersBuilder {
	v := readpolltimeout.Build()
	rb.v.ReadPollTimeout = v
	return rb
}
