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

package follow

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package follow
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ccr/follow/CreateFollowIndexRequest.ts#L25-L51
type Request struct {
	LeaderIndex                   *string        `json:"leader_index,omitempty"`
	MaxOutstandingReadRequests    *int64         `json:"max_outstanding_read_requests,omitempty"`
	MaxOutstandingWriteRequests   *int64         `json:"max_outstanding_write_requests,omitempty"`
	MaxReadRequestOperationCount  *int64         `json:"max_read_request_operation_count,omitempty"`
	MaxReadRequestSize            *string        `json:"max_read_request_size,omitempty"`
	MaxRetryDelay                 types.Duration `json:"max_retry_delay,omitempty"`
	MaxWriteBufferCount           *int64         `json:"max_write_buffer_count,omitempty"`
	MaxWriteBufferSize            *string        `json:"max_write_buffer_size,omitempty"`
	MaxWriteRequestOperationCount *int64         `json:"max_write_request_operation_count,omitempty"`
	MaxWriteRequestSize           *string        `json:"max_write_request_size,omitempty"`
	ReadPollTimeout               types.Duration `json:"read_poll_timeout,omitempty"`
	RemoteCluster                 *string        `json:"remote_cluster,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Follow request: %w", err)
	}

	return &req, nil
}
